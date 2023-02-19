package social

import (
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/model"

	model2 "ByteTech-7355608/douyin-server/kitex_gen/douyin/model"

	"ByteTech-7355608/douyin-server/kitex_gen/douyin/social"
	. "ByteTech-7355608/douyin-server/pkg/configs"
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
		Log.Errorf("get follow list err:%v", err)
		return nil, err
	}

	user_list := []*model2.User{}
	for _, v := range list {

		isfollow, err := s.dao.Relation.IsUserFollowed(ctx, userID, v.ID)
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
			IsFollow:      isfollow,
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
	for _, followerid := range followeridlist {
		userInstance, err := s.dao.User.FindUserById(ctx, followerid)
		if err != nil {
			Log.Infof("get follower err :%v", err)
			continue
		}

		isfollow, err := s.dao.Relation.IsUserFollowed(ctx, userID, followerid)
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
			IsFollow:      isfollow,
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
	for _, followerid := range followeridlist {
		userInstance, err := s.dao.User.FindUserById(ctx, followerid)
		if err != nil {
			Log.Infof("get follower err :%v", err)
			continue
		}

		isFriend, err := s.dao.Relation.IsUserFollowed(ctx, req.GetUserId(), followerid)
		if err != nil {
			Log.Infof("check friend err :%v", err)
			continue
		}

		if !isFriend {
			// 如果没有互相关注，则不是好友，跳过
			continue
		}

		// 获取和该好友最新的聊天信息
		msg1, err := s.dao.Message.GetLastMessageByUid(ctx, req.GetUserId(), followerid)
		if err != nil {
			Log.Infof("get message err :%v", err)
			continue
		}

		msg2, err := s.dao.Message.GetLastMessageByUid(ctx, followerid, req.GetUserId())
		if err != nil {
			Log.Infof("get message err :%v", err)
			continue
		}

		// 互相没有发送过信息
		if msg1.ID == 0 && msg2.ID == 0 {
			friend := &model.FriendUser{
				Id:            userInstance.ID,
				Name:          userInstance.Username,
				FollowCount:   &userInstance.FollowCount,
				FollowerCount: &userInstance.FollowerCount,
				Avatar:        &userInstance.Avatar,
				IsFollow:      true,
			}
			friends = append(friends, friend)
			Log.Infof(" %v to %v message is empty", req.GetUserId(), followerid)
			continue
		}

		var msg string
		var msgType int64
		if msg1.CreatedAt.Before(msg2.CreatedAt) {
			msg = msg2.Content
			msgType = 0
		} else {
			msg = msg1.Content
			msgType = 1
		}

		friend := &model.FriendUser{
			Id:            userInstance.ID,
			Name:          userInstance.Username,
			FollowCount:   &userInstance.FollowCount,
			FollowerCount: &userInstance.FollowerCount,
			Avatar:        &userInstance.Avatar,
			IsFollow:      true,
			Message:       &msg,
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
