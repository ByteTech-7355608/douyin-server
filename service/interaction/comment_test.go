package interactionimport_test

import (
	"ByteTech-7355608/douyin-server/dal/dao/model"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/interaction"
	model1 "ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"ByteTech-7355608/douyin-server/rpc"
	interactionimport "ByteTech-7355608/douyin-server/service/interaction"
	"context"
	"fmt"
	"sync"

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
	var ctx context.Context
	var user *model.User
	var db *gorm.DB
	BeforeEach(func() {
		once.Do(func() {
			db = getDB()
			mockRpc := rpc.NewMockRPC(gomock.NewController(GinkgoT()))
			svc = interactionimport.NewService(db, mockRpc)
		})
		ctx = context.Background()
		user = &model.User{
			Username: "a",
			Password: "b",
		}
	})

	Context("Test CommentAction", func() {
		It("test new comment", func() {
			//initialize
			commentTest1 := "test1"
			fmt.Println("start test: test new comment")
			err := db.WithContext(ctx).Model(model.User{}).Create(user).Error
			Expect(err).To(BeNil())
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
			// find the inserted record
			record := model.Comment{}
			err = db.First(&record, comment.Id).Error
			Expect(err).To(BeNil())
			Expect(record.Content).To(Equal(commentTest1))
			// clean
			err = db.WithContext(ctx).Where("ID = ?", user.ID).Delete(&model.User{}).Error
			Expect(err).To(BeNil())
			err = db.WithContext(ctx).Where("ID = ?", comment.Id).Delete(&model.Comment{}).Error
			Expect(err).To(BeNil())
		})

		It("test delete comment", func() {
			//initialize
			commentTest1 := "test1"
			fmt.Println("start test: test delete comment")
			err := db.WithContext(ctx).Model(model.User{}).Create(user).Error
			Expect(err).To(BeNil())
			req := interaction.DouyinCommentActionRequest{
				VideoId:     1,
				ActionType:  1,
				CommentText: &commentTest1,
				BaseReq: &model1.BaseReq{
					UserId:   &user.ID,
					Username: &user.Username,
				},
			}
			// do and check the insert
			resp, err := svc.CommentAction(ctx, &req)
			Expect(err).To(BeNil())
			comment := resp.Comment
			Expect(comment.Content).To(Equal(commentTest1))
			Expect(comment.User.Id).To(Equal(user.ID))
			// find the inserted record
			record := model.Comment{}
			err = db.First(&record, comment.Id).Error
			Expect(err).To(BeNil())
			Expect(record.Content).To(Equal(commentTest1))
			// delete the record
			req1 := interaction.DouyinCommentActionRequest{
				ActionType: 2,
				CommentId:  &comment.Id,
			}
			resp, err = svc.CommentAction(ctx, &req1)
			Expect(err).To(BeNil())
			// check the record is realy deleted
			err = db.First(&record, comment.Id).Error
			Expect(err).NotTo(BeNil())
			// clean
			err = db.WithContext(ctx).Where("ID = ?", user.ID).Delete(&model.User{}).Error
			Expect(err).To(BeNil())
			err = db.WithContext(ctx).Where("ID = ?", comment.Id).Delete(&model.Comment{}).Error
			Expect(err).To(BeNil())
		})
	})

})
