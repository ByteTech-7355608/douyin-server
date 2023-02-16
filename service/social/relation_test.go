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
	"errors"
	"regexp"
	"sync"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("Relation test", func() {
	var once sync.Once
	var svc *social.Service
	var mock sqlmock.Sqlmock
	var ctx context.Context
	var relationColumns []string

	BeforeEach(func() {
		once.Do(func() {
			configs.InitLogger()
			mock = dao.InitMockDB()
			mockRpc := rpc.NewMockRPC(gomock.NewController(GinkgoT()))
			svc = social.NewService(mockRpc)
		})
		ctx = context.Background()
		relationColumns = []string{"id", "concerner_id", "concerned_id", "action"}
	})

	Context("Test Relation", func() {

		It("test new relation ok", func() {
			// rs := mock.NewRows(relationColumns).AddRow(1, 1, 2, 1)
			relation_ex := model.Relation{
				ConcernerID: 1,
				ConcernedID: 2,
				Action:      true,
			}
			mock.ExpectQuery("SELECT `relation`.`id`,`relation`.`concerner_id`").
				WithArgs(relation_ex.ConcernerID, relation_ex.ConcernedID, 0).
				WillReturnError(gorm.ErrRecordNotFound)
			mock.ExpectBegin()
			//插入一对关系
			mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `relation`")).
				WillReturnResult(sqlmock.NewResult(1, 1))
				//更新关注
			mock.ExpectExec(regexp.QuoteMeta("UPDATE `user` SET ")).
				WillReturnResult(sqlmock.NewResult(1, 1))
				//更新粉丝
			mock.ExpectExec(regexp.QuoteMeta("UPDATE `user` SET ")).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()
			//initialize
			id := int64(1)
			name := "cbn"
			followuser := &model1.BaseReq{
				UserId:   &id,
				Username: &name,
			}
			req := rpcSocial.DouyinFollowActionRequest{
				Token:      "123",
				ToUserId:   2,
				ActionType: 1,
				BaseReq:    followuser,
			}
			//do and check the insert
			resp, err := svc.FollowAction(ctx, &req)
			Expect(err).To(BeNil())
			Expect(resp).NotTo(BeNil())
		})

		It("test check relation failed", func() {
			// rs := mock.NewRows(relationColumns).AddRow(1, 1, 2, 1)
			relation_ex := model.Relation{
				ConcernerID: 1,
				ConcernedID: 2,
				Action:      true,
			}
			mock.ExpectQuery("SELECT `relation`.`id`,`relation`.`concerner_id`").
				WithArgs(relation_ex.ConcernerID, relation_ex.ConcernedID, 0).
				WillReturnError(errors.New("check relation err"))
			id := int64(1)
			name := "cbn"
			followuser := &model1.BaseReq{
				UserId:   &id,
				Username: &name,
			}
			req := rpcSocial.DouyinFollowActionRequest{
				Token:      "123",
				ToUserId:   2,
				ActionType: 1,
				BaseReq:    followuser,
			}
			//do and check the insert
			resp, err := svc.FollowAction(ctx, &req)
			Expect(err).NotTo(BeNil())
			Expect(resp).To(BeNil())
		})
		It("test add new relation failed", func() {
			// rs := mock.NewRows(relationColumns).AddRow(1, 1, 2, 1)
			relation_ex := model.Relation{
				ConcernerID: 1,
				ConcernedID: 2,
				Action:      true,
			}
			mock.ExpectQuery("SELECT `relation`.`id`,`relation`.`concerner_id`").
				WithArgs(relation_ex.ConcernerID, relation_ex.ConcernedID, 0).
				WillReturnError(gorm.ErrRecordNotFound)
			mock.ExpectBegin()
			//插入一对关系
			mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `relation`")).
				WillReturnError(errors.New("add relation failed"))
			mock.ExpectRollback()
			//initialize
			id := int64(1)
			name := "cbn"
			followuser := &model1.BaseReq{
				UserId:   &id,
				Username: &name,
			}
			req := rpcSocial.DouyinFollowActionRequest{
				Token:      "123",
				ToUserId:   2,
				ActionType: 1,
				BaseReq:    followuser,
			}
			//do and check the insert
			resp, err := svc.FollowAction(ctx, &req)
			Expect(err).NotTo(BeNil())
			Expect(resp).To(BeNil())
		})

		It("test update relation to 1 ok", func() {
			// rs := mock.NewRows(relationColumns).AddRow(1, 1, 2, 1)
			relation_ex := model.Relation{
				ConcernerID: 1,
				ConcernedID: 2,
				Action:      false,
			}
			mock.ExpectQuery("SELECT `relation`.`id`,`relation`.`concerner_id`").
				WithArgs(relation_ex.ConcernerID, relation_ex.ConcernedID, 0).
				WillReturnRows(sqlmock.NewRows(relationColumns).AddRow(1, 1, 2, 0))
			mock.ExpectBegin()
			//插入一对关系
			mock.ExpectExec(regexp.QuoteMeta("UPDATE `relation`")).
				WillReturnResult(sqlmock.NewResult(1, 1))
				//更新关注
			mock.ExpectExec(regexp.QuoteMeta("UPDATE `user` SET ")).
				WillReturnResult(sqlmock.NewResult(1, 1))
				//更新粉丝
			mock.ExpectExec(regexp.QuoteMeta("UPDATE `user` SET ")).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()
			//initialize
			id := int64(1)
			name := "cbn"
			followuser := &model1.BaseReq{
				UserId:   &id,
				Username: &name,
			}
			req := rpcSocial.DouyinFollowActionRequest{
				Token:      "123",
				ToUserId:   2,
				ActionType: 1,
				BaseReq:    followuser,
			}
			//do and check the insert
			resp, err := svc.FollowAction(ctx, &req)
			Expect(err).To(BeNil())
			Expect(resp).NotTo(BeNil())
		})
		It("test update relation to 1 failed", func() {
			// rs := mock.NewRows(relationColumns).AddRow(1, 1, 2, 1)
			relation_ex := model.Relation{
				ConcernerID: 1,
				ConcernedID: 2,
				Action:      false,
			}
			mock.ExpectQuery("SELECT `relation`.`id`,`relation`.`concerner_id`").
				WithArgs(relation_ex.ConcernerID, relation_ex.ConcernedID, 0).
				WillReturnRows(sqlmock.NewRows(relationColumns).AddRow(1, 1, 2, 0))
			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta("UPDATE `relation`")).
				WillReturnError(errors.New("update relation to 1 err"))
			mock.ExpectRollback()
			//initialize
			id := int64(1)
			name := "cbn"
			followuser := &model1.BaseReq{
				UserId:   &id,
				Username: &name,
			}
			req := rpcSocial.DouyinFollowActionRequest{
				Token:      "123",
				ToUserId:   2,
				ActionType: 1,
				BaseReq:    followuser,
			}
			//do and check the insert
			resp, err := svc.FollowAction(ctx, &req)
			Expect(err).NotTo(BeNil())
			Expect(resp).To(BeNil())
		})
		It("test update relation to 0 ok", func() {
			// rs := mock.NewRows(relationColumns).AddRow(1, 1, 2, 1)
			relation_ex := model.Relation{
				ConcernerID: 1,
				ConcernedID: 2,
				Action:      false,
			}

			mock.ExpectQuery("SELECT `relation`.`id`,`relation`.`concerner_id`").
				WithArgs(relation_ex.ConcernerID, relation_ex.ConcernedID, 0).
				WillReturnRows(sqlmock.NewRows(relationColumns).AddRow(1, 1, 2, 1))
			mock.ExpectBegin()
			//插入一对关系
			mock.ExpectExec(regexp.QuoteMeta("UPDATE `relation`")).
				WillReturnResult(sqlmock.NewResult(1, 1))
				//更新关注
			mock.ExpectExec(regexp.QuoteMeta("UPDATE `user` SET ")).
				WillReturnResult(sqlmock.NewResult(1, 1))
				//更新粉丝
			mock.ExpectExec(regexp.QuoteMeta("UPDATE `user` SET ")).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()
			//initialize
			id := int64(1)
			name := "cbn"
			followuser := &model1.BaseReq{
				UserId:   &id,
				Username: &name,
			}
			req := rpcSocial.DouyinFollowActionRequest{
				Token:      "123",
				ToUserId:   2,
				ActionType: 2,
				BaseReq:    followuser,
			}
			//do and check the insert
			resp, err := svc.FollowAction(ctx, &req)
			Expect(err).To(BeNil())
			Expect(resp).NotTo(BeNil())
		})
		It("test update relation to 0 failed", func() {
			// rs := mock.NewRows(relationColumns).AddRow(1, 1, 2, 1)
			relation_ex := model.Relation{
				ConcernerID: 1,
				ConcernedID: 2,
				Action:      false,
			}

			mock.ExpectQuery("SELECT `relation`.`id`,`relation`.`concerner_id`").
				WithArgs(relation_ex.ConcernerID, relation_ex.ConcernedID, 0).
				WillReturnRows(sqlmock.NewRows(relationColumns).AddRow(1, 1, 2, 1))
			mock.ExpectBegin()
			//插入一对关系
			mock.ExpectExec(regexp.QuoteMeta("UPDATE `relation`")).
				WillReturnError(errors.New("update relation to 0 err"))
			mock.ExpectRollback()
			//initialize
			id := int64(1)
			name := "cbn"
			followuser := &model1.BaseReq{
				UserId:   &id,
				Username: &name,
			}
			req := rpcSocial.DouyinFollowActionRequest{
				Token:      "123",
				ToUserId:   2,
				ActionType: 2,
				BaseReq:    followuser,
			}
			//do and check the insert
			resp, err := svc.FollowAction(ctx, &req)
			Expect(err).NotTo(BeNil())
			Expect(resp).To(BeNil())
		})

	})
})
