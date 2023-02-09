package base_test

import (
	"ByteTech-7355608/douyin-server/dal/dao"
	"ByteTech-7355608/douyin-server/dal/dao/model"
	base2 "ByteTech-7355608/douyin-server/kitex_gen/douyin/base"
	model2 "ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	"ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"ByteTech-7355608/douyin-server/util"
	"errors"
	"time"

	"ByteTech-7355608/douyin-server/rpc"
	"ByteTech-7355608/douyin-server/service/base"
	"context"
	"regexp"
	"sync"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("User Test", func() {
	var once sync.Once
	var svc *base.Service
	var mock sqlmock.Sqlmock
	var ctx context.Context
	var user *model.User
	var userColumns []string
	var sTime time.Time
	var relationCol []string

	BeforeEach(func() {
		once.Do(func() {
			configs.InitLogger()
			mock = dao.InitMockDB()
			mockRpc := rpc.NewMockRPC(gomock.NewController(GinkgoT()))
			svc = base.NewService(mockRpc)
		})
		ctx = context.Background()

		userColumns = []string{"id", "created_at", "updated_at", "deleted_at", "username", "password",
			"follow_count", "follower_count"}

		user = &model.User{
			Username: "aaa",
			Password: "bbb",
		}
		relationCol = []string{"count(*)"}
	})

	Context("Test UserRegister", func() {

		It("test register user success", func() {

			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user`")).
				WithArgs(user.Username, 0).
				WillReturnError(gorm.ErrRecordNotFound)

			mock.ExpectBegin()
			mock.ExpectExec("INSERT INTO `user`").
				WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), user.Username, util.EncryptPassword(user.Password), sqlmock.AnyArg(), sqlmock.AnyArg()).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			req := base2.NewDouyinUserRegisterRequest()
			req.Username = "aaa"
			req.Password = "bbb"
			resp, err := svc.UserRegister(ctx, req)
			Expect(err).To(BeNil())
			Expect(resp.UserId).To(Equal(int64(1)))
		})

		It("test register user exist", func() {

			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user`")).
				WithArgs(user.Username, 0).
				WillReturnRows(sqlmock.NewRows(userColumns).
					AddRow(2, sTime, sTime, 0, "aaa", "xxx", 0, 0))

			req := base2.NewDouyinUserRegisterRequest()
			req.Username = "aaa"
			req.Password = "bbb"
			resp, err := svc.UserRegister(ctx, req)
			Expect(err).To(Equal(constants.ErrUserExist))
			Expect(resp.UserId).To(Equal(int64(0)))
		})

		It("test register user faild1", func() {

			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user`")).
				WithArgs(user.Username, 0).
				WillReturnError(gorm.ErrRecordNotFound)

			mock.ExpectBegin()
			mock.ExpectExec("INSERT INTO `user`").
				WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), user.Username, util.EncryptPassword(user.Password), sqlmock.AnyArg(), sqlmock.AnyArg()).
				WillReturnError(errors.New("some err !!!"))
			mock.ExpectRollback()

			req := base2.NewDouyinUserRegisterRequest()
			req.Username = "aaa"
			req.Password = "bbb"
			resp, err := svc.UserRegister(ctx, req)
			Expect(err).NotTo(BeNil())
			Expect(resp).NotTo(BeNil())
		})

		It("test register user faild2", func() {

			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user`")).
				WithArgs(user.Username, 0).
				WillReturnError(errors.New("some err !!!"))

			req := base2.NewDouyinUserRegisterRequest()
			req.Username = "aaa"
			req.Password = "bbb"
			resp, err := svc.UserRegister(ctx, req)
			Expect(err).NotTo(BeNil())
			Expect(resp.UserId).To(Equal(int64(0)))
		})
	})

	Context("Test UserLogin", func() {
		It("test login user success", func() {

			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user`")).
				WithArgs(user.Username, 0).
				WillReturnRows(sqlmock.NewRows(userColumns).
					AddRow(2, sTime, sTime, 0, "aaa", util.EncryptPassword("bbb"), 0, 0))

			req := base2.NewDouyinUserLoginRequest()
			req.Username = "aaa"
			req.Password = "bbb"
			resp, err := svc.UserLogin(ctx, req)
			Expect(err).To(BeNil())
			Expect(resp.UserId).To(Equal(int64(2)))
		})

		It("test login user not exist", func() {

			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user`")).
				WithArgs(user.Username, 0).
				WillReturnError(gorm.ErrRecordNotFound)

			req := base2.NewDouyinUserLoginRequest()
			req.Username = "aaa"
			req.Password = "bbb"
			resp, err := svc.UserLogin(ctx, req)
			Expect(err).To(Equal(constants.ErrUserNotExist))
			Expect(resp.UserId).To(Equal(int64(0)))
		})

		It("test login user invalid password", func() {

			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user`")).
				WithArgs(user.Username, 0).
				WillReturnError(errors.New("some err !!!"))

			req := base2.NewDouyinUserLoginRequest()
			req.Username = "aaa"
			req.Password = "bbb"
			resp, err := svc.UserLogin(ctx, req)
			Expect(err).NotTo(BeNil())
			Expect(resp.UserId).To(Equal(int64(0)))
		})

		It("test login user failed", func() {

			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user`")).
				WithArgs(user.Username, 0).
				WillReturnRows(sqlmock.NewRows(userColumns).
					AddRow(2, sTime, sTime, 0, "aaa", "ccc", 0, 0))

			req := base2.NewDouyinUserLoginRequest()
			req.Username = "aaa"
			req.Password = "bbb"
			resp, err := svc.UserLogin(ctx, req)
			Expect(err).To(Equal(constants.ErrInvalidPassword))
			Expect(resp.UserId).To(Equal(int64(0)))
		})
	})
	Context("Test UserMsg", func() {
		It("test select user success", func() {
			mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `relation` WHERE concerner_id=?")).WithArgs(1, 0).
				WillReturnRows(sqlmock.NewRows(relationCol).AddRow(1))
			mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `relation` WHERE concerned_id=?")).WithArgs(1, 0).
				WillReturnRows(sqlmock.NewRows(relationCol).AddRow(2))

			req := base2.NewDouyinUserRequest()
			req.BaseReq = new(model2.BaseReq)
			req.BaseReq.UserId = new(int64)
			req.BaseReq.Username = new(string)
			var userID int64 = 1
			var username string = "cbn"
			*req.BaseReq.UserId = userID
			*req.BaseReq.Username = username
			resp, err := svc.UserMsg(ctx, req)
			Expect(err).To(BeNil())
			Expect(resp.User.Id).To(Equal(int64(1)))
			var num1 int64 = 1
			var num2 int64 = 2
			Expect(resp.User.FollowCount).To(Equal(&num1))
			Expect(resp.User.FollowerCount).To(Equal(&num2))

		})

		It("test select user fail", func() {
			mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `relation` WHERE")).WithArgs(1, 0).
				WillReturnError(errors.New("count follow nums err"))
			mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `relation` WHERE")).WithArgs(1, 0).
				WillReturnError(errors.New("count follower nums err"))
			req := base2.NewDouyinUserRequest()
			req.BaseReq = new(model2.BaseReq)
			req.BaseReq.UserId = new(int64)
			req.BaseReq.Username = new(string)
			var userID int64 = 1
			var username string = "cbn"
			*req.BaseReq.UserId = userID
			*req.BaseReq.Username = username
			resp, err := svc.UserMsg(ctx, req)
			Expect(err).NotTo(BeNil())
			Expect(resp).To(BeNil())
		})
	})
})
