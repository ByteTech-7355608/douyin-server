package base_test

import (
	"ByteTech-7355608/douyin-server/dal/dao"
	"ByteTech-7355608/douyin-server/dal/dao/model"
	base2 "ByteTech-7355608/douyin-server/kitex_gen/douyin/base"
	"ByteTech-7355608/douyin-server/pkg/configs"
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
	var user *model.User
	var userColumns []string
	var videoColumns []string

	BeforeEach(func() {
		once.Do(func() {
			configs.InitLogger()
			mock = dao.InitMockDB()
			mockRpc := rpc.NewMockRPC(gomock.NewController(GinkgoT()))
			svc = base.NewService(mockRpc)
		})
		ctx = context.Background()

		userColumns = []string{"id", "username", "password", "follow_count", "follower_count"}

		videoColumns = []string{"id", "play_url", "cover_url", "favorite_count", "comment_count", "title", "uid"}

		user = &model.User{
			ID: 1,
		}

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

			req := base2.NewDouyinPublishListRequest()
			req.UserId = int64(1)
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
