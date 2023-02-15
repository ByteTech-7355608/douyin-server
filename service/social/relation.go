package social

import (
	model2 "ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/social"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"context"

	"gorm.io/gorm"
)

const (
	KAddType = 1
)

func (s *Service) FollowAction(ctx context.Context, req *social.DouyinFollowActionRequest) (resp *social.DouyinFollowActionResponse, err error) {
	resp = social.NewDouyinFollowActionResponse()
	concernerId := *req.BaseReq.UserId
	concernedId := req.FollowerId
	record, err := s.dao.Relation.FindRecordByConcernerIDAndConcernedID(ctx, concernerId, concernedId)
	if req.ActionType == KAddType {
		// 点赞
		// 1.数据库不存在两者关系数据
		if err == gorm.ErrRecordNotFound {
			err = s.dao.Relation.AddRelation(ctx, concernerId, concernedId)
			if err != nil {
				Log.Errorf("add relation err: %v", err)
				return resp, err
			}
			return
		}
		// 2.数据库存在两者关系数据--已点赞
		if err == nil && record.Action {
			return
		}
		// 3.数据库存在两者关系数据--未点赞
		if err == nil && !record.Action {
			err = s.dao.Relation.UpdatedRelation(ctx, record, 1)
			if err != nil {
				Log.Errorf("updated relation err: %v", err)
				return resp, err
			}
			return
		}
	} else {
		// 取消点赞
		// 1.数据库不存在两者关系数据
		if err == gorm.ErrRecordNotFound {
			return
		}
		// 2.数据库存在两者关系数据--已点赞
		if err == nil && record.Action {
			err = s.dao.Relation.UpdatedRelation(ctx, record, 0)
			if err != nil {
				Log.Errorf("add comment err: %v", err)
				return nil, err
			}
			return
		}
		// 3.数据库存在两者关系数据--未点赞
		if err == nil && !record.Action {
			return
		}
	}
	return resp, err
}

func (s *Service) FollowList(ctx context.Context, req *social.DouyinFollowingListRequest) (resp *social.DouyinFollowingListResponse, err error) {
	resp = social.NewDouyinFollowingListResponse()
	user_id := req.UserId
	list, err := s.dao.Relation.FollowList(ctx, user_id)
	if err != nil && len(list) == 0 {
		Log.Errorf("get follow list err:%v", err)
		return resp, err
	}
	user_list := []*model2.User{}
	for _, v := range list {
		user := &model2.User{
			Id:            v.ID,
			Name:          v.Username,
			FollowCount:   &v.FollowerCount,
			FollowerCount: &v.FollowerCount,
			IsFollow:      true,
		}
		user_list = append(user_list, user)
	}
	resp.FollowingList = user_list
	return
}
