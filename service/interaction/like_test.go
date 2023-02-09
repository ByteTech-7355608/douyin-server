package interaction_test

import (
	"ByteTech-7355608/douyin-server/dal/dao"
	interaction2 "ByteTech-7355608/douyin-server/kitex_gen/douyin/interaction"
	"ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/rpc"
	interactionimport "ByteTech-7355608/douyin-server/service/interaction"

	"context"
	"errors"
	"regexp"
	"sync"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("Like test", func() {
	var once sync.Once
	var svc *interactionimport.Service
	var mock sqlmock.Sqlmock
	var ctx context.Context
	var likeColumns []string
	var videoColumns []string
	var userColumns []string
	var relationColumns []string

	BeforeEach(func() {
		once.Do(func() {
			configs.InitLogger()
			mock = dao.InitMockDB()
			mockRpc := rpc.NewMockRPC(gomock.NewController(GinkgoT()))
			svc = interactionimport.NewService(mockRpc)
		})
		ctx = context.Background()
		likeColumns = []string{"id", "vid"}
		videoColumns = []string{"id", "play_url", "cover_url", "favorite_count", "comment_count", "title", "uid"}
		userColumns = []string{"id", "username", "password", "follow_count", "follower_count"}
		relationColumns = []string{"id", "action"}
	})

	Context("Test FavoriteList", func() {

		It("test FavoriteList success", func() {
			// 根据 uid 查询喜欢的 vid list
			mock.ExpectQuery("SELECT `vid` FROM `like`").
				WithArgs(1, 1, 0).
				WillReturnRows(sqlmock.NewRows(likeColumns).
					AddRow(1, 2).
					AddRow(2, 3).
					AddRow(3, 7).
					AddRow(4, 4).
					AddRow(5, 5))

			// 根据 vid 查询 video list
			// success
			mock.ExpectQuery("SELECT (.*) FROM `video` WHERE id = ").
				WithArgs(2, 0).
				WillReturnRows(sqlmock.NewRows(videoColumns).
					AddRow(2, "xxx", "xxx", 3, 4, "xxx", 2))
			// success
			mock.ExpectQuery("SELECT (.*) FROM `video` WHERE id = ").
				WithArgs(3, 0).
				WillReturnRows(sqlmock.NewRows(videoColumns).
					AddRow(3, "xxx", "xxx", 3, 4, "xxx", 3))
			// error page not found
			mock.ExpectQuery("SELECT (.*) FROM `video` WHERE id = ").
				WithArgs(7, 0).
				WillReturnError(gorm.ErrRecordNotFound)
			// success
			mock.ExpectQuery("SELECT (.*) FROM `video` WHERE id = ").
				WithArgs(4, 0).
				WillReturnRows(sqlmock.NewRows(videoColumns).
					AddRow(4, "xxx", "xxx", 3, 4, "xxx", 4))
			// success
			mock.ExpectQuery("SELECT (.*) FROM `video` WHERE id = ").
				WithArgs(5, 0).
				WillReturnRows(sqlmock.NewRows(videoColumns).
					AddRow(5, "xxx", "xxx", 3, 4, "xxx", 5))

			// 1 follow 2
			mock.ExpectQuery("SELECT (.*) FROM `user` WHERE id = ").
				WithArgs(2, 0).
				WillReturnRows(sqlmock.NewRows(userColumns).
					AddRow(2, "bbb", "xxx", 3, 4))
			mock.ExpectQuery(regexp.QuoteMeta("SELECT `action` FROM `relation`")).
				WithArgs(1, 2, 0).
				WillReturnRows(sqlmock.NewRows(relationColumns).
					AddRow(1, 1))

			// 1 not follow 3
			mock.ExpectQuery("SELECT (.*) FROM `user` WHERE id = ").
				WithArgs(3, 0).
				WillReturnRows(sqlmock.NewRows(userColumns).
					AddRow(3, "ccc", "xxx", 3, 4))
			mock.ExpectQuery(regexp.QuoteMeta("SELECT `action` FROM `relation`")).
				WithArgs(1, 3, 0).
				WillReturnRows(sqlmock.NewRows(relationColumns).
					AddRow(1, 0))

			// error page not found
			mock.ExpectQuery("SELECT (.*) FROM `user` WHERE id = ").
				WithArgs(4, 0).
				WillReturnRows(sqlmock.NewRows(userColumns).
					AddRow(4, "ddd", "xxx", 3, 4))
			mock.ExpectQuery(regexp.QuoteMeta("SELECT `action` FROM `relation`")).
				WithArgs(1, 4, 0).
				WillReturnError(gorm.ErrRecordNotFound)

			// user not found
			mock.ExpectQuery("SELECT (.*) FROM `user` WHERE id = ").
				WithArgs(5, 0).
				WillReturnError(gorm.ErrRecordNotFound)

			// 测试服务
			req := interaction2.NewDouyinFavoriteListRequest()
			req.UserId = 1
			resp, err := svc.FavoriteList(ctx, req)
			configs.Log.Infof("%v", resp)
			Expect(err).To(BeNil())
			Expect(resp).NotTo(BeNil())
		})

		It("test FavoriteList Fail by select like err", func() {
			// 根据 uid 查询喜欢的 vid list
			mock.ExpectQuery("SELECT `vid` FROM `like`").
				WithArgs(1, 1, 0).
				WillReturnError(errors.New("some err"))

			// 测试服务
			req := interaction2.NewDouyinFavoriteListRequest()
			req.UserId = 1
			resp, err := svc.FavoriteList(ctx, req)
			configs.Log.Infof("%v", resp)
			Expect(err).NotTo(BeNil())
			Expect(resp.VideoList).To(BeNil())
		})

		It("test FavoriteList Fail by select video err", func() {

			mock.ExpectQuery("SELECT `vid` FROM `like`").
				WithArgs(1, 1, 0).
				WillReturnRows(sqlmock.NewRows(likeColumns).
					AddRow(1, 2))

			mock.ExpectQuery("SELECT (.*) FROM `video` WHERE id = ").
				WithArgs(2, 0).
				WillReturnError(errors.New("some err"))

			// 测试服务
			req := interaction2.NewDouyinFavoriteListRequest()
			req.UserId = 1
			resp, err := svc.FavoriteList(ctx, req)
			configs.Log.Infof("%v", resp)
			Expect(err).NotTo(BeNil())
			Expect(resp.VideoList).To(BeNil())
		})

	})
})
