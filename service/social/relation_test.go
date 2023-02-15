package social_test

import (
	"ByteTech-7355608/douyin-server/cmd/api/biz/model/douyin/model"
	"ByteTech-7355608/douyin-server/cmd/api/biz/model/douyin/social"
	"context"
	"errors"
	rpcsocial "ByteTech-7355608/douyin-server/kitex_gen/douyin/social"

	"github.com/DATA-DOG/go-sqlmock"
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
		relationColumns = []string{"id", "concerner_id","concerned_id","action"}
		userColumns = []string{"id", "username", "password", "follow_count", "follower_count"}
	})

	Context("Test Comment", func() {
		 
		It("test new relation ok", func() {
			rs := mock.NewRows(relationColumns).AddRow(1,1,2,1)
			mock.ExpectQuery(regexp.QuoteMeta("SELECT * from `relation`")).
				WithArgs(1, 2).
				WillReturnError(gorm.ErrRecordNotFound)
			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `relation`")).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()
			//initialize
			
			req := rpc.
			//do and check the insert
			resp, err := svc.CommentAction(ctx, &req)
			Expect(err).To(BeNil())
			comment := resp.Comment
			Expect(comment.Content).To(Equal(commentTest1))
			Expect(comment.User.Id).To(Equal(arg_user.ID))
		})

	})
})