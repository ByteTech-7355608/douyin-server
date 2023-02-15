package social

import (
	model2 "ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/social"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"context"
)

const (
	KAddType    = 1
	KDeleteType = 2
)

func (s *Service) FollowAction(ctx context.Context, req *social.DouyinFollowActionRequest) (resp *social.DouyinFollowActionResponse, err error) {
	resp = social.NewDouyinFollowActionResponse()
	concerner_id := *req.BaseReq.UserId
	//concerned_id := req.to_user_id
	concerned_id := req.FollowerId
	flag, record, err := s.dao.Relation.CheckRecord(ctx, concerner_id, concerned_id)
	if err != nil {
		Log.Errorf("check relation err: %v", err)
		return nil, err
	}
	switch req.ActionType {
	case KAddType:
		if !flag {
			//没有此记录
			err = s.dao.Relation.AddRelation(ctx, concerner_id, concerned_id)
			if err != nil {
				Log.Errorf("add relation err: %v", err)
				return nil, err
			}
		} else {
			//存在该记录
			err = s.dao.Relation.UpdatedRelation(ctx, record, 1)
			if err != nil {
				Log.Errorf("update relation err: %v", err)
				return nil, err
			}

		}

	case KDeleteType:
		if flag {
			err = s.dao.Relation.UpdatedRelation(ctx, record, 0)
			if err != nil {
				Log.Errorf("update relation err: %v", err)
				return nil, err
			}
		}
	}
	return
}

func (s *Service) FollowList(ctx context.Context, req *social.DouyinFollowingListRequest) (resp *social.DouyinFollowingListResponse, err error) {
	resp = social.NewDouyinFollowingListResponse()
	user_id := req.UserId
	list, err := s.dao.Relation.FollowList(ctx, user_id)
	if err != nil {
		Log.Errorf("get follow list err:%v", err)
		return nil, err
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
