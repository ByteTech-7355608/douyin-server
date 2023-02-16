package base_test

import (
	"ByteTech-7355608/douyin-server/dal/dao"
	"ByteTech-7355608/douyin-server/dal/dao/model"
	model2 "ByteTech-7355608/douyin-server/kitex_gen/douyin/model"

	base2 "ByteTech-7355608/douyin-server/kitex_gen/douyin/base"
	"ByteTech-7355608/douyin-server/pkg/configs"

	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/rpc"
	"ByteTech-7355608/douyin-server/service/base"
	"context"

	"errors"
	"sync"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Publish Test", func() {
	var once sync.Once
	var svc *base.Service
	var mock sqlmock.Sqlmock
	var ctx context.Context
	var baseReq *model2.BaseReq
	var video *model.Video
	var user *model.User
	var userColumns []string
	var videoColumns []string

	BeforeEach(func() {
		once.Do(func() {
			configs.InitLogger()
			mock = dao.InitMockDB()
			mockRpc := rpc.NewMockRPC(gomock.NewController(GinkgoT()))
			svc = base.NewService(mockRpc)
			video = &model.Video{
				Title:         "title_test",
				UID:           1,
				PlayURL:       "http://clips.vorwaerts-gmbh.de/big_buck_bunny.mp4",
				CoverURL:      "https://fuss10.elemecdn.com/3/28/bbf893f792f03a54408b3b7a7ebf0jpeg.jpeg",
				FavoriteCount: 0,
				CommentCount:  0,
			}

			user = &model.User{
				ID:       1,
				Username: "zyj",
				Password: "123",
			}

			userID := int64(100)
			username := "aaa"
			baseReq = &model2.BaseReq{
				UserId:   &userID,
				Username: &username,
			}
		})
		ctx = context.Background()

		userColumns = []string{"id", "username", "password",
			"follow_count", "follower_count"}

		videoColumns = []string{"id", "play_url", "cover_url", "favorite_count", "comment_count", "title", "uid"}

	})

	Context("Test PublishAction", func() {
		var sqlInsert = "INSERT INTO `video`"

		It("test publish action success", func() {
			mock.ExpectBegin()
			mock.ExpectExec(sqlInsert).
				WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), video.PlayURL, video.CoverURL, video.FavoriteCount, video.CommentCount, video.Title, video.UID).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()
			req := base2.NewDouyinPublishActionRequest()
			req.Title = video.Title
			req.PlayUrl = &video.PlayURL
			req.CoverUrl = &video.CoverURL
			req.Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6Inp5ajMiLCJpc3MiOiJkb3V5aW4tc2VydmljZSIsImV4cCI6MTY3NjE5NzYwOH0.0j_FKdiUulkth1JtiNzEVX38kzfxrAlmN104SR_j6gY"
			req.BaseReq = &model2.BaseReq{
				UserId:   &user.ID,
				Username: &user.Username,
			}
			resp, err := svc.PublishAction(ctx, req)
			//Expect(err).To(BeNil())
			Log.Infof("resp:%+v, err:%+v", resp, err)
			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(int32(0)))
		})

	})

	Context("Test Publish lists", func() {

		It("test publish lists success", func() {

			mock.ExpectQuery("SELECT (.*) FROM `user`").
				WithArgs(user.ID, 0).
				WillReturnRows(sqlmock.NewRows(userColumns).
					AddRow(1, "aaa", "xxx", 0, 0))

			mock.ExpectQuery("SELECT (.*) FROM `video`").
				WithArgs(user.ID, 0).
				WillReturnRows(sqlmock.NewRows(videoColumns).
					AddRow(2, "xxx", "xxx", 0, 0, "xxx", 1).
					AddRow(3, "xxx", "xxx", 0, 0, "xxx", 1))

			mock.ExpectQuery("SELECT `action` FROM `relation`").
				WithArgs(100, 2, 0).
				WillReturnRows(sqlmock.NewRows([]string{"action"}).
					AddRow(true))

			mock.ExpectQuery("SELECT `action` FROM `relation`").
				WithArgs(100, 3, 0).
				WillReturnRows(sqlmock.NewRows([]string{"action"}).
					AddRow(true))

			req := base2.NewDouyinPublishListRequest()
			req.UserId = int64(1)
			req.BaseReq = baseReq
			resp, err := svc.PublishList(ctx, req)
			Expect(err).To(BeNil())
			Expect(resp.VideoList).NotTo(BeNil())
		})

		It("test publish lists failed1", func() {

			mock.ExpectQuery("SELECT (.*) FROM `user`").
				WithArgs(user.ID, 0).
				WillReturnRows(sqlmock.NewRows(userColumns).
					AddRow(1, "aaa", "xxx", 0, 0))

			mock.ExpectQuery("SELECT (.*) FROM `video`").
				WithArgs(user.ID, 0).
				WillReturnError(errors.New("some err "))

			req := base2.NewDouyinPublishListRequest()
			req.UserId = int64(1)
			resp, err := svc.PublishList(ctx, req)
			Expect(err).NotTo(BeNil())
			Expect(resp.VideoList).To(BeNil())
		})

		It("test publish lists failed2", func() {

			mock.ExpectQuery("SELECT (.*) FROM `user`").
				WithArgs(user.ID, 0).
				WillReturnError(errors.New("some err "))

			req := base2.NewDouyinPublishListRequest()
			req.UserId = int64(1)
			resp, err := svc.PublishList(ctx, req)
			Expect(err).NotTo(BeNil())
			Expect(resp.VideoList).To(BeNil())
		})

	})

})
