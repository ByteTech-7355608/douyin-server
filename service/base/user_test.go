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

var _ = Describe("User Test", func() {
	var once sync.Once
	var svc *base.Service
	var mock sqlmock.Sqlmock
	var ctx context.Context
	var user *model.User
	BeforeEach(func() {
		once.Do(func() {
			configs.InitLogger()
			mock = dao.InitMockDB()
			mockRpc := rpc.NewMockRPC(gomock.NewController(GinkgoT()))
			svc = base.NewService(mockRpc)
		})
		ctx = context.Background()

		user = &model.User{
			Username: "a",
			Password: "b",
		}
	})

	Context("Test UserRegister", func() {
		var sqlInsert = "INSERT INTO `user`"

		It("test register user success", func() {
			mock.ExpectBegin()
			mock.ExpectExec(sqlInsert).
				WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), user.Username, user.Password, sqlmock.AnyArg(), sqlmock.AnyArg()).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			req := base2.NewDouyinUserRegisterRequest()
			req.Username = "a"
			req.Password = "b"
			resp, err := svc.UserRegister(ctx, req)
			Expect(err).To(BeNil())
			Expect(resp.UserId).To(Equal(int64(1)))
		})

		It("test register user fail", func() {
			mock.ExpectBegin()
			var sqlInsert = "INSERT INTO `user`"
			mock.ExpectExec(sqlInsert).
				WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), user.Username, user.Password, sqlmock.AnyArg(), sqlmock.AnyArg()).
				WillReturnError(errors.New("some err"))
			mock.ExpectCommit()

			req := base2.NewDouyinUserRegisterRequest()
			req.Username = "a"
			req.Password = "b"
			resp, err := svc.UserRegister(ctx, req)
			Expect(err).NotTo(BeNil())
			Expect(resp).To(BeNil())
		})
	})
})
