package base_test

import (
	"ByteTech-7355608/douyin-server/dal/dao/model"
	base2 "ByteTech-7355608/douyin-server/kitex_gen/douyin/base"
	"ByteTech-7355608/douyin-server/rpc"
	"ByteTech-7355608/douyin-server/service"
	"ByteTech-7355608/douyin-server/service/base"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
	"sync"
)

var _ = Describe("Publish Test", func() {
	var once sync.Once
	var svc *base.Service
	var mock sqlmock.Sqlmock
	var ctx context.Context
	var video *model.Video
	BeforeEach(func() {
		once.Do(func() {
			var db *gorm.DB
			db, mock = service.GetMockDB()
			mockRpc := rpc.NewMockRPC(gomock.NewController(GinkgoT()))
			svc = base.NewService(db, mockRpc)
		})
		ctx = context.Background()

		video = &model.Video{
			PlayURL:  "https://www.bilibili.com/video/BV13v4y1b7tD/?vd_source=745b384339e17e2e9c0971310a560da7",
			CoverURL: "https://i1.hdslb.com/bfs/archive/f2841e4aa39150a17d5323a1041e7bbe7a37b09e.jpg@672w_378h_1c_!web-home-common-cover.webp",
			Title:    "50台和平主义者出击！薇薇公主终于出现！",
			UID:      1,
		}
	})

	Context("Test PublishAction", func() {
		var sqlInsert = "INSERT INTO `video`"

		It("test publish action success", func() {
			mock.ExpectBegin()
			mock.ExpectExec(sqlInsert).
				WithArgs(video.PlayURL, video.CoverURL, video.Title, video.UID, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			req := base2.NewDouyinPublishActionRequest()
			req.Title = "a"
			req.Data = []byte{'a', 'b', 'c'}
			resp, err := svc.PublishAction(ctx, req)
			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
		})

		It("test publish action fail", func() {
			mock.ExpectBegin()
			var sqlInsert = "INSERT INTO `user`"
			mock.ExpectExec(sqlInsert).
				WithArgs(video.PlayURL, video.CoverURL, video.Title, video.UID, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
				WillReturnError(errors.New("some err"))
			mock.ExpectCommit()

			req := base2.NewDouyinPublishActionRequest()
			req.Title = "a"
			req.Data = []byte{'a', 'b', 'c'}
			resp, err := svc.PublishAction(ctx, req)
			Expect(err).NotTo(BeNil())
			Expect(resp).To(BeNil())
		})
	})
})
