package social

import (
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/model"

	model2 "ByteTech-7355608/douyin-server/kitex_gen/douyin/model"

	"ByteTech-7355608/douyin-server/kitex_gen/douyin/social"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"context"
)

const (
	KAddType    = 1
	KDeleteType = 2
)

func (s *Service) FollowList(ctx context.Context, req *social.DouyinFollowingListRequest) (resp *social.DouyinFollowingListResponse, err error) {
	resp = social.NewDouyinFollowingListResponse()
	userID := req.GetBaseReq().GetUserId()

	list, err := s.dao.Relation.FollowList(ctx, req.GetUserId())
	if err != nil {
		err = constants.ErrQueryRecord
		Log.Errorf("get follow list err:%v", err)
		return nil, err
	}

	var user_list []*model2.User
	for _, v := range list {

		isFollow, err := s.dao.Relation.IsFollower(ctx, userID, v.ID)
		if err != nil {
			Log.Infof("check follow err :%v", err)
			continue
		}

		user := &model2.User{
			Id:            v.ID,
			Name:          v.Username,
			FollowCount:   &v.FollowerCount,
			FollowerCount: &v.FollowerCount,
			Avatar:        &v.Avatar,
			IsFollow:      isFollow,
		}
		user_list = append(user_list, user)
	}
	resp.UserList = user_list
	return
}

func (s *Service) FollowerList(ctx context.Context, req *social.DouyinFollowerListRequest) (resp *social.DouyinFollowerListResponse, err error) {
	resp = social.NewDouyinFollowerListResponse()
	userID := req.GetBaseReq().GetUserId()
	// 根据 uid 从 Relation 表中查找用户粉丝 idlist，然后根据 id 查询 userlist
	followeridlist, err := s.dao.Relation.GetFollowerListByUid(ctx, req.GetUserId())
	if err != nil {
		Log.Errorf("get follower list err: %v", err)
		return nil, err
	}

	var followers []*model.User
	for _, followerId := range followeridlist {
		userInstance, err := s.dao.User.FindUserById(ctx, followerId)
		if err != nil {
			Log.Infof("get follower err :%v", err)
			continue
		}

		isFollow, err := s.dao.Relation.IsFollower(ctx, userID, followerId)
		if err != nil {
			Log.Infof("check follow err :%v", err)
			continue
		}

		user := &model.User{
			Id:            userInstance.ID,
			Name:          userInstance.Username,
			FollowCount:   &userInstance.FollowCount,
			FollowerCount: &userInstance.FollowerCount,
			Avatar:        &userInstance.Avatar,
			IsFollow:      isFollow,
		}
		followers = append(followers, user)
	}

	resp.SetUserList(followers)
	return resp, nil
}

func (s *Service) FriendList(ctx context.Context, req *social.DouyinRelationFriendListRequest) (resp *social.DouyinRelationFriendListResponse, err error) {
	resp = social.NewDouyinRelationFriendListResponse()
	// 根据 uid 从 Relation 表中查找用户粉丝 idlist，然后根据 id 查询 userlist 并判断是否互相关注
	followeridlist, err := s.dao.Relation.GetFollowerListByUid(ctx, req.GetUserId())
	if err != nil {
		Log.Errorf("get follower list err: %v", err)
		return
	}

	var friends []*model.FriendUser
	for _, followerId := range followeridlist {
		userInstance, err := s.dao.User.FindUserById(ctx, followerId)
		if err != nil {
			Log.Infof("get follower err :%v", err)
			continue
		}

		isFriend, err := s.dao.Relation.IsFollower(ctx, req.GetUserId(), followerId)
		if err != nil {
			Log.Infof("check friend err :%v", err)
			continue
		}

		if !isFriend {
			// 如果没有互相关注，则不是好友，跳过
			continue
		}

		// 获取和该好友最新的聊天信息
		msg, err := s.dao.Message.GetLastMessageByUid(ctx, req.GetUserId(), followerId)
		if err != nil {
			Log.Infof("get message err :%v", err)
			continue
		}

		// 互相没有发送过信息
		if msg.ID == 0 {
			friend := &model.FriendUser{
				Id:            userInstance.ID,
				Name:          userInstance.Username,
				FollowCount:   &userInstance.FollowCount,
				FollowerCount: &userInstance.FollowerCount,
				Avatar:        &userInstance.Avatar,
				IsFollow:      true,
			}
			friends = append(friends, friend)
			Log.Infof(" %v to %v message is empty", req.GetUserId(), followerId)
			continue
		}

		var msgType int64
		if msg.ToUID == followerId {
			msgType = 1
		} else {
			msgType = 0
		}

		friend := &model.FriendUser{
			Id:            userInstance.ID,
			Name:          userInstance.Username,
			FollowCount:   &userInstance.FollowCount,
			FollowerCount: &userInstance.FollowerCount,
			Avatar:        &userInstance.Avatar,
			IsFollow:      true,
			Message:       &msg.Content,
			MsgType:       msgType,
		}

		friends = append(friends, friend)
	}

	resp.SetUserList(friends)
	return resp, nil
}

func (s *Service) FollowAction(ctx context.Context, req *social.DouyinFollowActionRequest) (resp *social.DouyinFollowActionResponse, err error) {
	resp = social.NewDouyinFollowActionResponse()
	concerner_id := *req.BaseReq.UserId
	concerned_id := req.ToUserId
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
				err = constants.ErrCreateRecord
				Log.Errorf("add relation err: %v", err)
				return resp, err
			}
		} else {
			//存在该记录
			err = s.dao.Relation.UpdatedRelation(ctx, record, 1)
			if err != nil {
				err = constants.ErrUpdateRecord
				Log.Errorf("update relation err: %v", err)
				return resp, err
			}

		}

	case KDeleteType:
		if flag {
			err = s.dao.Relation.UpdatedRelation(ctx, record, 0)
			if err != nil {
				err = constants.ErrUpdateRecord
				Log.Errorf("update relation err: %v", err)
				return resp, err
			}
		}
	}
	return
}
