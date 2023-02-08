package interactionimport_test

import (
	"ByteTech-7355608/douyin-server/dal/dao"
	"ByteTech-7355608/douyin-server/dal/dao/model"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/interaction"
	model1 "ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	"ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/pkg/constants"
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
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getDB() (db *gorm.DB) {
	db, err := gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	return
}

var _ = Describe("CommentAction Test", func() {
	var once sync.Once
	var svc *interactionimport.Service
	var mock sqlmock.Sqlmock
	var ctx context.Context
	var user *model.User
	BeforeEach(func() {
		once.Do(func() {
			configs.InitLogger()
			mock = dao.InitMockDB()
			mockRpc := rpc.NewMockRPC(gomock.NewController(GinkgoT()))
			svc = interactionimport.NewService(mockRpc)
		})
		ctx = context.Background()
		user = &model.User{
			ID:       1,
			Username: "a",
			Password: "b",
		}
	})

	Context("Test CommentAction", func() {
		var sqlInsert = "INSERT INTO `comment`"
		It("test new comment ok", func() {
			rs := mock.NewRows([]string{"ID", "Username", "Password"}).AddRow(1, user.Username, user.Password)
			mock.ExpectQuery(regexp.QuoteMeta("SELECT `username`,`follow_count`,`follower_count` FROM `user` WHERE ID = ")).
				WithArgs(1).
				WillReturnRows(rs)
			mock.ExpectBegin()
			mock.ExpectExec(sqlInsert).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()
			//initialize
			commentTest1 := "test1"
			req := interaction.DouyinCommentActionRequest{
				VideoId:     1,
				ActionType:  1,
				CommentText: &commentTest1,
				BaseReq: &model1.BaseReq{
					UserId:   &user.ID,
					Username: &user.Username,
				},
			}
			//do and check the insert
			resp, err := svc.CommentAction(ctx, &req)
			Expect(err).To(BeNil())
			comment := resp.Comment
			Expect(comment.Content).To(Equal(commentTest1))
			Expect(comment.User.Id).To(Equal(user.ID))
		})

		It("test new comment, user not exist", func() {
			mock.ExpectQuery(regexp.QuoteMeta("SELECT `username`,`follow_count`,`follower_count` FROM `user` WHERE ID = ")).
				WithArgs(1).
				WillReturnError(errors.New("some err"))
			//initialize
			commentTest1 := "test1"
			req := interaction.DouyinCommentActionRequest{
				VideoId:     1,
				ActionType:  1,
				CommentText: &commentTest1,
				BaseReq: &model1.BaseReq{
					UserId:   &user.ID,
					Username: &user.Username,
				},
			}
			//do and check the insert
			_, err := svc.CommentAction(ctx, &req)
			Expect(err).NotTo(BeNil())
		})

		It("test new comment ok, insert err", func() {
			rs := mock.NewRows([]string{"ID", "Username", "Password"}).AddRow(1, user.Username, user.Password)
			mock.ExpectQuery(regexp.QuoteMeta("SELECT `username`,`follow_count`,`follower_count` FROM `user` WHERE ID = ")).
				WithArgs(1).
				WillReturnRows(rs)
			mock.ExpectBegin()
			mock.ExpectExec(sqlInsert).
				WillReturnError(errors.New("some err"))
			mock.ExpectRollback()
			//initialize
			commentTest1 := "test1"
			req := interaction.DouyinCommentActionRequest{
				VideoId:     1,
				ActionType:  1,
				CommentText: &commentTest1,
				BaseReq: &model1.BaseReq{
					UserId:   &user.ID,
					Username: &user.Username,
				},
			}
			//do and check the insert
			_, err := svc.CommentAction(ctx, &req)
			Expect(err).NotTo(BeNil())
		})

		It("test delete comment", func() {
			//initialize
			var sqlDelete = "UPDATE `comment`"
			mock.ExpectBegin()
			mock.ExpectExec(sqlDelete).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()
			// delete the record
			commentID := int64(1)
			req1 := interaction.DouyinCommentActionRequest{
				ActionType: 2,
				CommentId:  &commentID,
			}
			_, err := svc.CommentAction(ctx, &req1)
			Expect(err).To(BeNil())
		})

		It("test delete comment err", func() {
			//initialize
			var sqlDelete = "UPDATE `comment`"
			mock.ExpectBegin()
			mock.ExpectExec(sqlDelete).
				WillReturnError(errors.New("some err"))
			mock.ExpectCommit()
			// delete the record
			commentID := int64(1)
			req1 := interaction.DouyinCommentActionRequest{
				ActionType: 2,
				CommentId:  &commentID,
			}
			_, err := svc.CommentAction(ctx, &req1)
			Expect(err).NotTo(BeNil())
		})
	})

})
