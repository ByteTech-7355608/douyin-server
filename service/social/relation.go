package social

import (
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/social"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"context"
)

func (s *Service) FollowerList(ctx context.Context, req *social.DouyinFollowerListRequest) (resp *social.DouyinFollowerListResponse, err error) {
	resp = social.NewDouyinFollowerListResponse()
	// 根据 uid 从 Relation 表中查找用户粉丝 idlist，然后根据 id 查询 userlist
	followeridlist, err := s.dao.Relation.GetFollowerListByUid(ctx, req.GetUserId())
	if err != nil {
		Log.Errorf("get follower list err: %v", err)
		return
	}

	var followers []*model.User
	for _, followerid := range followeridlist {
		userInstance, err := s.dao.User.FindUserById(ctx, followerid)
		if err != nil {
			Log.Infof("get follower err :%v", err)
			continue
		}

		user := &model.User{
			Id:            userInstance.ID,
			Name:          userInstance.Username,
			FollowCount:   &userInstance.FollowCount,
			FollowerCount: &userInstance.FollowerCount,
			IsFollow:      true,
		}
		followers = append(followers, user)
	}

	resp.SetFollowerList(followers)
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

		var msg string
		var msgType int64
		if msg1.CreatedAt.Before(msg2.CreatedAt) {
			msg = msg2.Content
			msgType = 0
		} else {
			msg = msg1.Content
			msgType = 1
		}

		user := &model.User{
			Id:            userInstance.ID,
			Name:          userInstance.Username,
			FollowCount:   &userInstance.FollowCount,
			FollowerCount: &userInstance.FollowerCount,
			IsFollow:      true,
		}

		friend := &model.FriendUser{
			User:    user,
			Message: &msg,
			MsgType: msgType,
		}

		friends = append(friends, friend)
	}

	resp.SetUserList(friends)
	return resp, nil
}
