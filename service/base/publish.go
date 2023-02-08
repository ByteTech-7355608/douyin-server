package base

import (
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/base"
	"context"
)

func (s *Service) PublishList(ctx context.Context, req *base.DouyinPublishListRequest) (resp *base.DouyinPublishListResponse, err error) {
	resp = base.NewDouyinPublishListResponse()

	videlist, err := s.dao.User.PublishList(ctx, req.GetUserId())
	return
}
