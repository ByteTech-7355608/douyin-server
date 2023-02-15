package social_test

import (
	"ByteTech-7355608/douyin-server/dal/dao"
	social2 "ByteTech-7355608/douyin-server/kitex_gen/douyin/social"
	"ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/rpc"
	"ByteTech-7355608/douyin-server/service/social"
	"context"
	"errors"
	"sync"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Relation Test", func() {
	var once sync.Once
	var svc *social.Service
	var mock sqlmock.Sqlmock
	var ctx context.Context
	var userColumns []string
	var relationColumns []string
	var relationColumns2 []string
	var messageColumns []string

	BeforeEach(func() {
		once.Do(func() {
			configs.InitLogger()
			mock = dao.InitMockDB()
			mockRpc := rpc.NewMockRPC(gomock.NewController(GinkgoT()))
			svc = social.NewService(mockRpc)
		})

		ctx = context.Background()
		userColumns = []string{"id", "username", "password", "follow_count", "follower_count"}
		relationColumns = []string{"concerner_id"}
		relationColumns2 = []string{"id", "action"}
		messageColumns = []string{"id", "uid", "to_uid", "content"}
	})

	Context("Test FollowerList", func() {
		It("Test FollowerList success", func() {
			// 根据uid查找关注uid的idlist
			mock.ExpectQuery("SELECT `concerner_id` FROM `relation`").
				WithArgs(1, 1, 0).
				WillReturnRows(sqlmock.NewRows(relationColumns).
					AddRow(2).
					AddRow(3))

			// 根据idlist查找userlist
			mock.ExpectQuery("SELECT (.*) FROM `user` WHERE id = ").
				WithArgs(2, 0).
				WillReturnRows(sqlmock.NewRows(userColumns).
					AddRow(2, "syx", "xxx", 0, 0))

			mock.ExpectQuery("SELECT (.*) FROM `user` WHERE id = ").
				WithArgs(3, 0).
				WillReturnRows(sqlmock.NewRows(userColumns).
					AddRow(3, "czh", "xxx", 0, 0))

			// 测试服务
			req := social2.NewDouyinFollowerListRequest()
			req.UserId = 1
			resp, err := svc.FollowerList(ctx, req)
			Expect(err).To(BeNil())
			Expect(resp).NotTo(BeNil())
		})

		It("Test FollowerList failed", func() {
			mock.ExpectQuery("SELECT `concerner_id` FROM `relation`").
				WithArgs(1, 1, 0).
				WillReturnError(errors.New("some err "))

			req := social2.NewDouyinFollowerListRequest()
			req.UserId = 1
			resp, err := svc.FollowerList(ctx, req)
			Expect(err).NotTo(BeNil())
			Expect(resp).To(BeNil())
		})
	})

	Context("Test FriendList", func() {
		It("Test FriendList success", func() {
			// 根据uid查找关注uid的idlist
			mock.ExpectQuery("SELECT `concerner_id` FROM `relation`").
				WithArgs(1, 1, 0).
				WillReturnRows(sqlmock.NewRows(relationColumns).
					AddRow(2).
					AddRow(3))

			// 1 follow 2
			mock.ExpectQuery("SELECT (.*) FROM `user` WHERE id = ").
				WithArgs(2, 0).
				WillReturnRows(sqlmock.NewRows(userColumns).
					AddRow(2, "syx", "xxx", 0, 0))

			mock.ExpectQuery("SELECT `action` FROM `relation`").
				WithArgs(1, 2, 0).
				WillReturnRows(sqlmock.NewRows(relationColumns2).
					AddRow(1, 1))

			// msg 1->2
			mock.ExpectQuery("SELECT (.*) FROM `message`").
				WithArgs(1, 2, 0).
				WillReturnRows(sqlmock.NewRows(messageColumns).
					AddRow(1, 1, 2, "xxxx"))

			// msg 2->1
			mock.ExpectQuery("SELECT (.*) FROM `message`").
				WithArgs(2, 1, 0).
				WillReturnRows(sqlmock.NewRows(messageColumns).
					AddRow(1, 2, 1, "????"))

			// 1 follow 3
			mock.ExpectQuery("SELECT (.*) FROM `user` WHERE id = ").
				WithArgs(3, 0).
				WillReturnRows(sqlmock.NewRows(userColumns).
					AddRow(3, "czh", "xxx", 0, 0))

			mock.ExpectQuery("SELECT `action` FROM `relation`").
				WithArgs(1, 3, 0).
				WillReturnRows(sqlmock.NewRows(relationColumns2).
					AddRow(2, 0))

			req := social2.NewDouyinRelationFriendListRequest()
			req.UserId = 1
			resp, err := svc.FriendList(ctx, req)
			Expect(err).To(BeNil())
			Expect(resp).NotTo(BeNil())
		})

	})
})
