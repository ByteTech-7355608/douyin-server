package social_test

import (
	"ByteTech-7355608/douyin-server/dal/dao"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	social2 "ByteTech-7355608/douyin-server/kitex_gen/douyin/social"
	"ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"ByteTech-7355608/douyin-server/rpc"
	"ByteTech-7355608/douyin-server/service/social"
	"context"
	"errors"
	"regexp"
	"sync"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Publish Test", func() {
	var once sync.Once
	var svc *social.Service
	var mock sqlmock.Sqlmock
	var ctx context.Context
	var messageColumns []string
	var baseReq *model.BaseReq
	var sTime time.Time

	BeforeEach(func() {
		once.Do(func() {
			configs.InitLogger()
			mock = dao.InitMockDB()
			mockRpc := rpc.NewMockRPC(gomock.NewController(GinkgoT()))
			svc = social.NewService(mockRpc)
			sTime = time.Now()

			userID := int64(100)
			username := "aaa"
			baseReq = &model.BaseReq{
				UserId:   &userID,
				Username: &username,
			}
		})
		ctx = context.Background()

		messageColumns = []string{"id", "created_at", "updated_at", "deleted_at", "uid", "to_uid", "content"}
	})

	Context("Test Message List", func() {
		It("test message list success", func() {
			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `message`")).
				WithArgs(0, 100, 1, 1, 100, 0).
				WillReturnRows(sqlmock.NewRows(messageColumns).
					AddRow(1, sTime.Add(-1), sTime, 0, 100, 1, "100 to 1").
					AddRow(2, sTime, sTime, 0, 1, 100, "1 to 100"))

			req := social2.NewDouyinMessageChatRequest()
			req.BaseReq = baseReq
			req.ToUserId = 1
			resp, err := svc.MessageList(ctx, req)
			Expect(err).To(BeNil())
			Expect(len(resp.MessageList)).To(Equal(2))
		})
	})

	Context("Test Send Message", func() {
		It("test send message success", func() {
			mock.ExpectBegin()
			mock.ExpectExec("INSERT INTO `message`").
				WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), 100, 1, "test").
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			req := social2.NewDouyinMessageActionRequest()
			req.BaseReq = baseReq
			req.ToUserId = 1
			req.Content = "test"
			req.ActionType = 1
			_, err := svc.SendMessage(ctx, req)
			Expect(err).To(BeNil())
		})

		It("test send message failure", func() {
			mock.ExpectBegin()
			mock.ExpectExec("INSERT INTO `message`").
				WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), 100, 1, "test").
				WillReturnError(errors.New("some err"))
			mock.ExpectCommit()

			req := social2.NewDouyinMessageActionRequest()
			req.BaseReq = baseReq
			req.ToUserId = 1
			req.Content = "test"
			req.ActionType = 1
			_, err := svc.SendMessage(ctx, req)
			Expect(err).To(Equal(constants.ErrCreateRecord))
		})
	})
})
