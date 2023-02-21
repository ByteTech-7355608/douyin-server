package base_test

import (
	"ByteTech-7355608/douyin-server/dal/dao"
	base2 "ByteTech-7355608/douyin-server/kitex_gen/douyin/base"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	"ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/rpc"
	"ByteTech-7355608/douyin-server/service/base"
	"context"
	"regexp"
	"sync"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User Test", func() {
	var once sync.Once
	var svc *base.Service
	var mock sqlmock.Sqlmock
	var ctx context.Context
	var sTime time.Time
	var baseReq *model.BaseReq
	var videoColumns []string
	//var likeColumns []string
	//var relationColumns []string
	var userColumns []string
	BeforeEach(func() {
		once.Do(func() {
			configs.InitLogger()
			mock = dao.InitMockDB()
			mockRpc := rpc.NewMockRPC(gomock.NewController(GinkgoT()))
			svc = base.NewService(mockRpc)

			userID := int64(100)
			username := "aaa"
			baseReq = &model.BaseReq{
				UserId:   &userID,
				Username: &username,
			}
			sTime = time.Now()
			videoColumns = []string{"id", "created_at", "updated_at", "deleted_at", "play_url", "cover_url",
				"favorite_count", "comment_count", "title", "uid"}
			//likeColumns = []string{"id", "created_at", "updated_at", "deleted_at", "uid", "vid",
			//	"action"}
			//relationColumns = []string{"id", "created_at", "updated_at", "deleted_at", "concerner_id", "concerned_id",
			//	"action"}
			userColumns = []string{"id", "created_at", "updated_at", "deleted_at", "username", "password",
				"follow_count", "follower_count"}
		})
		ctx = context.Background()
	})

	Context("Test Feed", func() {
		It("test video list", func() {
			req := base2.NewDouyinFeedRequest()
			req.BaseReq = baseReq

			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `video`")).
				WithArgs(sqlmock.AnyArg(), 0).
				WillReturnRows(sqlmock.NewRows(videoColumns).
					AddRow(1, sTime, sTime, 0, "/video/1", "/cover/1", 5, 3, "test_video", 2))
			mock.ExpectQuery(regexp.QuoteMeta("SELECT `action` FROM `like`")).
				WithArgs(100, 1, 0).
				WillReturnError(gorm.ErrRecordNotFound)
			mock.ExpectQuery("SELECT (.*) FROM `user`").
				WithArgs(sqlmock.AnyArg(), 0).
				WillReturnRows(sqlmock.NewRows(userColumns).
					AddRow(2, sTime, sTime, 0, "aaa", "xxx", 5, 3))
			mock.ExpectQuery(regexp.QuoteMeta("SELECT `action` FROM `relation`")).
				WithArgs(100, 2, 0).
				WillReturnRows(sqlmock.NewRows([]string{"action"}).
					AddRow(true))

			resp, err := svc.Feed(ctx, req)
			Expect(err).To(BeNil())
			Expect(len(resp.VideoList)).To(Equal(1))
			Expect(resp.VideoList[0].IsFavorite).To(Equal(false))
			Expect(resp.VideoList[0].Author.IsFollow).To(Equal(true))
			Expect(*resp.NextTime).To(Equal(sTime.Unix()))
		})
	})
})
