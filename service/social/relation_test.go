package social_test

import (
	"ByteTech-7355608/douyin-server/dal/dao"

	"ByteTech-7355608/douyin-server/dal/dao/model"
	model2 "ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
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

	var userColumns []string
	var relationColumns1 []string
	var relationColumns2 []string
	var messageColumns []string
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

		userColumns = []string{"id", "username", "password", "follow_count", "follower_count"}
		relationColumns1 = []string{"concerner_id"}
		relationColumns2 = []string{"id", "action"}
		messageColumns = []string{"id", "uid", "to_uid", "content"}
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
			followuser := &model2.BaseReq{
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
			followuser := &model2.BaseReq{
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
			followuser := &model2.BaseReq{
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
			_, err := svc.FollowAction(ctx, &req)
			Expect(err).NotTo(BeNil())
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
			followuser := &model2.BaseReq{
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
			followuser := &model2.BaseReq{
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
			_, err := svc.FollowAction(ctx, &req)
			Expect(err).NotTo(BeNil())
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
			followuser := &model2.BaseReq{
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
			followuser := &model2.BaseReq{
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
			_, err := svc.FollowAction(ctx, &req)
			Expect(err).NotTo(BeNil())
		})

	})

	Context("Test FollowerList", func() {
		It("Test FollowerList success", func() {
			// 根据uid查找关注uid的idlist
			mock.ExpectQuery("SELECT `concerner_id` FROM `relation`").
				WithArgs(1, 1, 0).
				WillReturnRows(sqlmock.NewRows(relationColumns1).
					AddRow(2).
					AddRow(3))

			// 根据idlist查找userlist
			mock.ExpectQuery("SELECT (.*) FROM `user`").
				WithArgs(2, 0).
				WillReturnRows(sqlmock.NewRows(userColumns).
					AddRow(2, "syx", "xxx", 0, 0))

			mock.ExpectQuery("SELECT `action` FROM `relation`").
				WithArgs(100, 2, 0).
				WillReturnRows(sqlmock.NewRows(relationColumns2).
					AddRow(1, 1))

			mock.ExpectQuery("SELECT (.*) FROM `user`").
				WithArgs(3, 0).
				WillReturnRows(sqlmock.NewRows(userColumns).
					AddRow(3, "czh", "xxx", 0, 0))

			mock.ExpectQuery("SELECT `action` FROM `relation`").
				WithArgs(100, 3, 0).
				WillReturnRows(sqlmock.NewRows(relationColumns2).
					AddRow(2, 1))

			// 测试服务
			req := rpcSocial.NewDouyinFollowerListRequest()
			req.UserId = 1
			userID := int64(100)
			username := "aaa"
			baseReq := &model2.BaseReq{
				UserId:   &userID,
				Username: &username,
			}
			req.BaseReq = baseReq
			resp, err := svc.FollowerList(ctx, req)
			Expect(err).To(BeNil())
			Expect(resp).NotTo(BeNil())
		})

		It("Test FollowerList failed", func() {
			mock.ExpectQuery("SELECT `concerner_id` FROM `relation`").
				WithArgs(1, 1, 0).
				WillReturnError(errors.New("some err "))

			req := rpcSocial.NewDouyinFollowerListRequest()
			req.UserId = 1
			userID := int64(100)
			username := "aaa"
			baseReq := &model2.BaseReq{
				UserId:   &userID,
				Username: &username,
			}
			req.BaseReq = baseReq
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
				WillReturnRows(sqlmock.NewRows(relationColumns1).
					AddRow(2).
					AddRow(3))

			// 1 follow 2
			mock.ExpectQuery("SELECT (.*) FROM `user`").
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

			// 1 follow 3
			mock.ExpectQuery("SELECT (.*) FROM `user` WHERE id = ").
				WithArgs(3, 0).
				WillReturnRows(sqlmock.NewRows(userColumns).
					AddRow(3, "czh", "xxx", 0, 0))

			mock.ExpectQuery("SELECT `action` FROM `relation`").
				WithArgs(1, 3, 0).
				WillReturnRows(sqlmock.NewRows(relationColumns2).
					AddRow(2, 0))

			req := rpcSocial.NewDouyinRelationFriendListRequest()
			req.UserId = 1
			resp, err := svc.FriendList(ctx, req)
			Expect(err).To(BeNil())
			Expect(resp).NotTo(BeNil())
		})

	})

})
