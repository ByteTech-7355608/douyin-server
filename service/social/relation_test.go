package social_test

import (
	"ByteTech-7355608/douyin-server/dal/dao"
	social2 "ByteTech-7355608/douyin-server/kitex_gen/douyin/social"
	"ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/rpc"
	"ByteTech-7355608/douyin-server/service/social"
	"context"
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

	BeforeEach(func() {
		once.Do(func() {
			configs.InitLogger()
			mock = dao.InitMockDB()
			mockRpc := rpc.NewMockRPC(gomock.NewController(GinkgoT()))
			svc = social.NewService(mockRpc)
		})

		ctx = context.Background()
		userColumns = []string{"id", "username", "password", "follow_count", "follower_count"}
		relationColumns = []string{"id", "concerner_id", "concerned_id", "action"}
		//messageColumns = []string{"id", "uid", "to_uid", "content"}
	})

	Context("Test FollowerList", func() {
		It("Test FollowerList success", func() {
			// 根据uid查找关注uid的idlist
			mock.ExpectQuery("SELECT `concerner_id` FROM `relation`").
				WithArgs(1, 1, 0).
				WillReturnRows(sqlmock.NewRows(relationColumns).
					AddRow(1, 2, 1, 1).
					AddRow(2, 3, 1, 1))

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
	})

	// Context("Test FriendList", func() {

	// })
})
