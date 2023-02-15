package social_test

import (
	"ByteTech-7355608/douyin-server/dal/dao"
	"ByteTech-7355608/douyin-server/dal/dao/model"
	model1 "ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	rpcSocial "ByteTech-7355608/douyin-server/kitex_gen/douyin/social"
	"ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/rpc"
	"ByteTech-7355608/douyin-server/service/social"
	"context"
	"regexp"
	"sync"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("Comment test", func() {
	var once sync.Once
	var svc *social.Service
	var mock sqlmock.Sqlmock
	var ctx context.Context
	var arg_user *model.User
	var userColumns []string
	var baseReq *model1.BaseReq
	var relationColumns []string
	var sTime time.Time

	BeforeEach(func() {
		once.Do(func() {
			configs.InitLogger()
			mock = dao.InitMockDB()
			mockRpc := rpc.NewMockRPC(gomock.NewController(GinkgoT()))
			svc = social.NewService(mockRpc)
		})
		ctx = context.Background()
		arg_user = &model.User{
			ID:       1,
			Username: "a",
			Password: "b",
		}
		userID := int64(100)
		username := "aaa"
		baseReq = &model1.BaseReq{
			UserId:   &userID,
			Username: &username,
		}
		userColumns = []string{"id", "username", "password", "follow_count", "follower_count"}
		relationColumns = []string{"id", "created_at", "updated_at", "deleted_at", "concerner_id", "concerned_id", "action"}
	})

	Context("Test Relation", func() {
		It("test query FollowList success", func() {
			mock.ExpectQuery("SELECT `concerned_id` FROM `relation`").
				WithArgs(arg_user.ID, 1, 0).
				WillReturnRows(sqlmock.NewRows([]string{"concerned_id"}).
					AddRow(2).
					AddRow(3).
					AddRow(4).
					AddRow(5))
			mock.ExpectQuery("SELECT (.*) FROM `user` WHERE id = ").
				WithArgs(2, 0).
				WillReturnRows(sqlmock.NewRows(userColumns).
					AddRow(2, "222", "password22", 3, 4))
			mock.ExpectQuery("SELECT (.*) FROM `user` WHERE id = ").
				WithArgs(3, 0).
				WillReturnError(gorm.ErrRecordNotFound)
			mock.ExpectQuery("SELECT (.*) FROM `user` WHERE id = ").
				WithArgs(4, 0).
				WillReturnRows(sqlmock.NewRows(userColumns).
					AddRow(4, "444", "password44", 3, 4))
			mock.ExpectQuery("SELECT (.*) FROM `user` WHERE id = ").
				WithArgs(5, 0).
				WillReturnRows(sqlmock.NewRows(userColumns).
					AddRow(5, "555", "password55", 3, 4))

			req := rpcSocial.NewDouyinFollowingListRequest()
			req.UserId = 1
			resp, err := svc.FollowList(ctx, req)
			configs.Log.Infof("%v", resp)
			Expect(err).To(BeNil())
			Expect(resp).NotTo(BeNil())
		})

		It("test new FollowList fail, no data from relation", func() {
			mock.ExpectQuery("SELECT `concerned_id` FROM `relation`").
				WithArgs(arg_user.ID, 1, 0).
				WillReturnError(gorm.ErrRecordNotFound)

			req := rpcSocial.NewDouyinFollowingListRequest()
			req.UserId = 1
			resp, err := svc.FollowList(ctx, req)
			configs.Log.Infof("%v", resp)
			Expect(err).NotTo(BeNil())
		})

		It("test new FollowList fail, no data from user", func() {
			mock.ExpectQuery("SELECT `concerned_id` FROM `relation`").
				WithArgs(arg_user.ID, 1, 0).
				WillReturnRows(sqlmock.NewRows([]string{"concerned_id"}).
					AddRow(2))

			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user` WHERE id = ")).
				WithArgs(2, 0).
				WillReturnError(gorm.ErrRecordNotFound)
			mock.ExpectRollback()
			req := rpcSocial.NewDouyinFollowingListRequest()
			req.UserId = 1
			resp, err := svc.FollowList(ctx, req)
			configs.Log.Infof("%v", resp)
			Expect(err).NotTo(BeNil())
		})

		It("test FollowAction to add follow success1", func() {
			//数据库不存在两者关系数据
			mock.ExpectQuery("SELECT (.*) FROM `relation`").
				WithArgs(1, 2, 0).
				WillReturnError(gorm.ErrRecordNotFound)

			mock.ExpectBegin()
			var sqlInsert = "INSERT INTO `relation`"
			mock.ExpectExec(sqlInsert).
				WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), 1, 2, 1).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			req := rpcSocial.NewDouyinFollowActionRequest()
			req.ActionType = 1
			req.BaseReq = baseReq
			req.FollowerId = 2
			resp, err := svc.FollowAction(ctx, req)
			configs.Log.Infof("%v", resp)
			Expect(err).NotTo(BeNil())
		})

		It("test FollowAction to add follow success2", func() {
			//数据库存在两者关系数据--已点赞
			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `relation`")).
				WithArgs(1, 2, 0).
				WillReturnRows(sqlmock.NewRows(relationColumns).
					AddRow(1, sTime, sTime, sTime, 1, 2, 1))
			req := rpcSocial.NewDouyinFollowActionRequest()
			req.ActionType = 1
			req.BaseReq = baseReq
			req.FollowerId = 2
			resp, err := svc.FollowAction(ctx, req)
			configs.Log.Infof("%v", resp)
			Expect(err).NotTo(BeNil())
		})

		It("test FollowAction to add follow success3", func() {
			//数据库存在两者关系数据--未点赞
			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `relation`")).
				WithArgs(1, 2, 0).
				WillReturnRows(sqlmock.NewRows(relationColumns).
					AddRow(1, sTime, sTime, sTime, 1, 2, 0))

			var sqlUpdate1 = "UPDATE `relation`"
			var sqlUpdate2 = "UPDATE `user`"
			mock.ExpectBegin()
			mock.ExpectExec(sqlUpdate1).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectExec(sqlUpdate2).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectExec(sqlUpdate2).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()
			req := rpcSocial.NewDouyinFollowActionRequest()
			req.ActionType = 1
			req.BaseReq = baseReq
			req.FollowerId = 2
			resp, err := svc.FollowAction(ctx, req)
			configs.Log.Infof("%v", resp)
			Expect(err).NotTo(BeNil())
		})

		It("test FollowAction to delete follow success1", func() {
			//数据库存在两者关系数据--未点赞
			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `relation`")).
				WithArgs(1, 2, 0).
				WillReturnRows(sqlmock.NewRows(relationColumns).
					AddRow(1, sTime, sTime, sTime, 1, 2, 0))

			req := rpcSocial.NewDouyinFollowActionRequest()
			req.ActionType = 2
			req.BaseReq = baseReq
			req.FollowerId = 2
			resp, err := svc.FollowAction(ctx, req)
			configs.Log.Infof("%v", resp)
			Expect(err).NotTo(BeNil())
		})

		It("test FollowAction to add follow success2", func() {
			//数据库存在两者关系数据--已点赞
			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `relation`")).
				WithArgs(1, 2, 0).
				WillReturnRows(sqlmock.NewRows(relationColumns).
					AddRow(1, sTime, sTime, sTime, 1, 2, 1))

			var sqlUpdate1 = "UPDATE `relation`"
			var sqlUpdate2 = "UPDATE `user`"
			mock.ExpectBegin()
			mock.ExpectExec(sqlUpdate1).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectExec(sqlUpdate2).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectExec(sqlUpdate2).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()
			req := rpcSocial.NewDouyinFollowActionRequest()
			req.ActionType = 2
			req.BaseReq = baseReq
			req.FollowerId = 2
			resp, err := svc.FollowAction(ctx, req)
			configs.Log.Infof("%v", resp)
			Expect(err).NotTo(BeNil())
		})

		It("test FollowAction to delete follow success3", func() {
			//数据库不存在两者关系数据
			mock.ExpectQuery("SELECT (.*) FROM `relation`").
				WithArgs(1, 2, 0).
				WillReturnError(gorm.ErrRecordNotFound)

			req := rpcSocial.NewDouyinFollowActionRequest()
			req.ActionType = 2
			req.BaseReq = baseReq
			req.FollowerId = 2
			resp, err := svc.FollowAction(ctx, req)
			configs.Log.Infof("%v", resp)
			Expect(err).NotTo(BeNil())
		})
	})
})
