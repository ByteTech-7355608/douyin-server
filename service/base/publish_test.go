package base_test

import (
	"ByteTech-7355608/douyin-server/dal/dao"
	"ByteTech-7355608/douyin-server/dal/dao/model"
	base2 "ByteTech-7355608/douyin-server/kitex_gen/douyin/base"
	"ByteTech-7355608/douyin-server/pkg/configs"

	"ByteTech-7355608/douyin-server/rpc"
	"ByteTech-7355608/douyin-server/service/base"
	"context"
	"sync"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Publish Test", func() {
	var once sync.Once
	var svc *base.Service
	var mock sqlmock.Sqlmock
	var ctx context.Context
	var video *model.Video
	BeforeEach(func() {
		once.Do(func() {
			configs.InitLogger()
			mock = dao.InitMockDB()
			mockRpc := rpc.NewMockRPC(gomock.NewController(GinkgoT()))
			svc = base.NewService(mockRpc)

			video = &model.Video{
				Title: "title_test",
				UID:   1,
			}
		})
		ctx = context.Background()
	})

	Context("Test PublishAction", func() {
		var sqlInsert = "INSERT INTO `video`"

		It("test publish action success", func() {
			mock.ExpectBegin()
			mock.ExpectExec(sqlInsert).
				WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), video.Title, video.UID).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			req := base2.NewDouyinPublishActionRequest()
			req.Title = "title_test"
			req.Data = []byte{'a', 'b', 'c'}
			req.Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6Inp5ajMiLCJpc3MiOiJkb3V5aW4tc2VydmljZSIsImV4cCI6MTY3NTkzMTc2NH0.N933u_g8ilx82CUogDhy7cK0CZ8iIIa3hNGDEIRyiQ4"
			resp, err := svc.PublishAction(ctx, req)
			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(int32(200)))
		})

	})
})
