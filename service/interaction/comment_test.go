package interaction_test

import (
	"ByteTech-7355608/douyin-server/dal/dao"
	"ByteTech-7355608/douyin-server/dal/dao/model"
	model1 "ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	"errors"
	"gorm.io/gorm"
	"regexp"
	"sync"
	"time"

	rpc_interaction "ByteTech-7355608/douyin-server/kitex_gen/douyin/interaction"
	"ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/rpc"
	"ByteTech-7355608/douyin-server/service/interaction"
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Comment test", func() {
	var once sync.Once
	var svc *interaction.Service
	var mock sqlmock.Sqlmock
	var ctx context.Context
	var arg_user *model.User
	var video *model.Video
	var commentColumns []string
	var userColumns []string
	var sTime time.Time

	BeforeEach(func() {
		once.Do(func() {
			configs.InitLogger()
			mock = dao.InitMockDB()
			mockRpc := rpc.NewMockRPC(gomock.NewController(GinkgoT()))
			svc = interaction.NewService(mockRpc)
		})
		ctx = context.Background()
		arg_user = &model.User{
			ID:       1,
			Username: "a",
			Password: "b",
		}
		commentColumns = []string{"id", "created_at", "vid", "uid", "content"}
		userColumns = []string{"id", "username", "password", "follow_count", "follower_count"}
		video = &model.Video{
			ID: 1,
		}
	})

	Context("Test Comment", func() {
		var sqlInsert = "INSERT INTO `comment`"
		It("test new comment ok", func() {
			rs := mock.NewRows([]string{"ID", "Username", "Password"}).AddRow(1, arg_user.Username, arg_user.Password)
			mock.ExpectQuery(regexp.QuoteMeta("SELECT `username`,`follow_count`,`follower_count` FROM `user`")).
				WithArgs(1, 0).
				WillReturnRows(rs)
			mock.ExpectBegin()
			mock.ExpectExec(sqlInsert).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()
			//initialize
			commentTest1 := "test1"
			req := rpc_interaction.DouyinCommentActionRequest{
				VideoId:     1,
				ActionType:  1,
				CommentText: &commentTest1,
				BaseReq: &model1.BaseReq{
					UserId:   &arg_user.ID,
					Username: &arg_user.Username,
				},
			}
			//do and check the insert
			resp, err := svc.CommentAction(ctx, &req)
			Expect(err).To(BeNil())
			comment := resp.Comment
			Expect(comment.Content).To(Equal(commentTest1))
			Expect(comment.User.Id).To(Equal(arg_user.ID))
		})

		It("test new comment, user not exist", func() {
			mock.ExpectQuery(regexp.QuoteMeta("SELECT `username`,`follow_count`,`follower_count` FROM `user`")).
				WithArgs(1, 0).
				WillReturnError(errors.New("some err"))
			//initialize
			commentTest1 := "test1"
			req := rpc_interaction.DouyinCommentActionRequest{
				VideoId:     1,
				ActionType:  1,
				CommentText: &commentTest1,
				BaseReq: &model1.BaseReq{
					UserId:   &arg_user.ID,
					Username: &arg_user.Username,
				},
			}
			//do and check the insert
			_, err := svc.CommentAction(ctx, &req)
			Expect(err).NotTo(BeNil())
		})

		It("test new comment ok, insert err", func() {
			rs := mock.NewRows([]string{"ID", "Username", "Password"}).AddRow(1, arg_user.Username, arg_user.Password)
			mock.ExpectQuery(regexp.QuoteMeta("SELECT `username`,`follow_count`,`follower_count`")).
				WithArgs(1, 0).
				WillReturnRows(rs)
			mock.ExpectBegin()
			mock.ExpectExec(sqlInsert).
				WillReturnError(errors.New("some err"))
			mock.ExpectRollback()
			//initialize
			commentTest1 := "test1"
			req := rpc_interaction.DouyinCommentActionRequest{
				VideoId:     1,
				ActionType:  1,
				CommentText: &commentTest1,
				BaseReq: &model1.BaseReq{
					UserId:   &arg_user.ID,
					Username: &arg_user.Username,
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
			req1 := rpc_interaction.DouyinCommentActionRequest{
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
			mock.ExpectRollback()
			// delete the record
			commentID := int64(1)
			req1 := rpc_interaction.DouyinCommentActionRequest{
				ActionType: 2,
				CommentId:  &commentID,
			}
			_, err := svc.CommentAction(ctx, &req1)
			Expect(err).NotTo(BeNil())
		})

		It("test Query Comment List Success", func() {
			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `comment`")).
				WithArgs(video.ID, 0).
				WillReturnRows(sqlmock.NewRows(commentColumns).
					AddRow(1, sTime, 1, 1, "contents1").
					AddRow(2, sTime, 1, 2, "contents2").
					AddRow(3, sTime, 1, 3, "contents3").
					AddRow(4, sTime, 1, 4, "contents4"))

			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user` WHERE id = ")).
				WithArgs(1, 0).
				WillReturnRows(sqlmock.NewRows(userColumns).
					AddRow(1, "111", "password11", 3, 4))
			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user` WHERE id = ")).
				WithArgs(2, 0).
				WillReturnRows(sqlmock.NewRows(userColumns).
					AddRow(2, "222", "password22", 3, 4))
			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user` WHERE id = ")).
				WithArgs(3, 0).
				WillReturnError(gorm.ErrRecordNotFound)
			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user` WHERE id = ")).
				WithArgs(1, 0).
				WillReturnRows(sqlmock.NewRows(userColumns).
					AddRow(4, "444", "password44", 3, 4))

			// 测试服务
			req := rpc_interaction.NewDouyinCommentListRequest()
			req.VideoId = 1
			resp, err := svc.CommentList(ctx, req)
			configs.Log.Infof("%v", resp)
			Expect(err).To(BeNil())
			Expect(resp).NotTo(BeNil())
		})

		/*It("test Query Comment not exist", func() {
			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `comment`")).
				WithArgs(video.ID, 0).
				WillReturnError(gorm.ErrRecordNotFound)
			// 测试服务
			req := rpc_interaction.NewDouyinCommentListRequest()
			req.VideoId = 2
			resp, err := svc.CommentList(ctx, req)
			configs.Log.Infof("%v", resp)
			Expect(err).To(BeNil())
		})*/
	})
})
