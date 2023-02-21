package social

import (
	"ByteTech-7355608/douyin-server/dal/cache"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	"strconv"

	model2 "ByteTech-7355608/douyin-server/kitex_gen/douyin/model"

	"ByteTech-7355608/douyin-server/kitex_gen/douyin/social"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"context"
)

func (s *Service) FollowAction(ctx context.Context, req *social.DouyinFollowActionRequest) (resp *social.DouyinFollowActionResponse, err error) {
	resp = social.NewDouyinFollowActionResponse()
	from_id, to_id := req.GetBaseReq().GetUserId(), req.GetToUserId()

	// 1. 判断需要操作的对象在缓存中是否存在
	if s.cache.Relation.FollowIsExists(ctx, from_id) == 0 {
		// 缓存中不存在用户粉丝列表
		userList, err := s.dao.Relation.FollowidList(ctx, from_id)
		if err != nil {
			Log.Errorf("get follow list err: %v, uid: %v", err, from_id)
		}
		if len(userList) > 0 {
			kv := make([]string, 0)
			for _, user := range userList {
				kv = append(kv, strconv.FormatInt(user, 10))
				kv = append(kv, "1")
			}
			if !s.cache.Relation.SetFollowList(ctx, from_id, kv...) {
				Log.Errorf("set follow list to redis err")
				return resp, constants.ErrWriteCache
			}
		}
	}

	if s.cache.Relation.FollowerIsExists(ctx, to_id) == 0 {
		// 缓存中不存在用户粉丝列表
		userList, err := s.dao.Relation.FolloweridList(ctx, to_id)
		if err != nil {
			Log.Errorf("get follower list err: %v, uid: %v", err, to_id)
		}
		if len(userList) > 0 {
			kv := make([]string, 0)
			for _, user := range userList {
				kv = append(kv, strconv.FormatInt(user, 10))
				kv = append(kv, "1")
			}
			if !s.cache.Relation.SetFollowerList(ctx, to_id, kv...) {
				Log.Errorf("set follower list to redis err")
				return resp, constants.ErrWriteCache
			}
		}
	}

	if s.cache.User.IsExists(ctx, from_id) == 0 {
		// 当关注的用户不存在
		user, err := s.dao.User.QueryUser(ctx, from_id)
		if err != nil {
			Log.Errorf("query user %v err: %v", from_id, err)
			return resp, err
		}
		userModel := &cache.UserModel{
			Id:              user.ID,
			Name:            user.Username,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoriteCount,
		}
		if !s.cache.User.SetUserMessage(ctx, userModel) {
			Log.Errorf("set user message to redis err: %v", err)
			return resp, constants.ErrWriteCache
		}
	}

	if s.cache.User.IsExists(ctx, to_id) == 0 {
		// 当被关注的用户不存在
		user, err := s.dao.User.QueryUser(ctx, to_id)
		if err != nil {
			Log.Errorf("query user %v err: %v", to_id, err)
			return resp, err
		}
		userModel := &cache.UserModel{
			Id:              user.ID,
			Name:            user.Username,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoriteCount,
		}
		if !s.cache.User.SetUserMessage(ctx, userModel) {
			Log.Errorf("set user message to redis err: %v", err)
			return resp, constants.ErrWriteCache
		}
	}

	// 2. 判断操作类型
	var action int64
	follow := s.cache.Relation.IsFollow(ctx, from_id, to_id)
	if follow && req.GetActionType() == 2 {
		// 已关注并取消关注
		action = -1
	} else if !follow && req.GetActionType() == 1 {
		// 关注
		action = 1
	} else {
		// 关系和操作不匹配
		return resp, constants.ErrUpdateRecord
	}

	// 3. 原子性的进行操作
	// 更新用户关注列表
	if !s.cache.Relation.FollowAction(ctx, from_id, to_id, action) {
		return resp, constants.ErrWriteCache
	}
	// 更新用户关注数量
	if !s.cache.User.IncrUserField(ctx, from_id, "follow_count", action) {
		return resp, constants.ErrWriteCache
	}
	// 更新用户被关注数量
	if !s.cache.User.IncrUserField(ctx, to_id, "follower_count", action) {
		return resp, constants.ErrWriteCache
	}

	return
}

func (s *Service) FollowList(ctx context.Context, req *social.DouyinFollowingListRequest) (resp *social.DouyinFollowingListResponse, err error) {
	resp = social.NewDouyinFollowingListResponse()
	user_id, from_id := req.GetBaseReq().GetUserId(), req.GetUserId()

	var followidList []int64
	if s.cache.Relation.FollowIsExists(ctx, from_id) == 0 {
		// 缓存中不存在查询用户粉丝列表
		followidList, err := s.dao.Relation.FollowidList(ctx, from_id)
		if err != nil {
			Log.Errorf("get follow list err: %v, uid: %v", err, from_id)
		}
		if len(followidList) > 0 {
			kv := make([]string, 0)
			for _, user := range followidList {
				kv = append(kv, strconv.FormatInt(user, 10))
				kv = append(kv, "1")
			}
			if !s.cache.Relation.SetFollowList(ctx, from_id, kv...) {
				Log.Errorf("set follow list to redis err")
				return resp, constants.ErrWriteCache
			}
		}
	} else {
		// 缓存中存在用户粉丝列表
		followidList = s.cache.Relation.GetFollowList(ctx, from_id)
	}

	if s.cache.Relation.FollowIsExists(ctx, user_id) == 0 {
		// 缓存中不存在登录用户粉丝列表
		userList, err := s.dao.Relation.FollowidList(ctx, user_id)
		if err != nil {
			Log.Errorf("get follow list err: %v, uid: %v", err, user_id)
		}
		if len(userList) > 0 {
			kv := make([]string, 0)
			for _, user := range userList {
				kv = append(kv, strconv.FormatInt(user, 10))
				kv = append(kv, "1")
			}
			if !s.cache.Relation.SetFollowList(ctx, user_id, kv...) {
				Log.Errorf("set follow list to redis err")
				return resp, constants.ErrWriteCache
			}
		}

	}

	// 遍历查找查询用户所关注的用户
	followList := []*model2.User{}
	if len(followidList) > 0 {
		for _, followid := range followidList {
			var follow = &model2.User{}
			if s.cache.User.IsExists(ctx, followid) == 0 {
				// 如果要查询的用户不在缓存中
				v, err := s.dao.User.QueryUser(ctx, followid)
				if err != nil {
					Log.Errorf("query user %v err: %v", followid, err)
					return resp, err
				}
				follow = &model2.User{
					Id:            v.ID,
					Name:          v.Username,
					FollowCount:   &v.FollowerCount,
					FollowerCount: &v.FollowerCount,
					Avatar:        &v.Avatar,
				}

			} else {
				// 要查询的用户位于缓存中
				userModel, err := s.cache.User.GetUserMessage(ctx, followid)
				if err != nil {
					return resp, constants.ErrReadCache
				}

				follow = cache.UserModel2User(userModel)
			}

			// 查找缓存看登录用户是否关注当前用户
			isfollow := s.cache.Relation.IsFollow(ctx, user_id, followid)
			follow.SetIsFollow(isfollow)
			followList = append(followList, follow)
		}
	}

	resp.SetUserList(followList)
	return
}

func (s *Service) FollowerList(ctx context.Context, req *social.DouyinFollowerListRequest) (resp *social.DouyinFollowerListResponse, err error) {
	resp = social.NewDouyinFollowerListResponse()
	user_id, to_id := req.GetBaseReq().GetUserId(), req.GetUserId()

	var followeridList []int64
	if s.cache.Relation.FollowerIsExists(ctx, to_id) == 0 {
		followeridList, err := s.dao.Relation.FolloweridList(ctx, to_id)
		if err != nil {
			Log.Errorf("get follow list err: %v, uid: %v", err, to_id)
		}
		if len(followeridList) > 0 {
			kv := make([]string, 0)
			for _, user := range followeridList {
				kv = append(kv, strconv.FormatInt(user, 10))
				kv = append(kv, "1")
			}
			if !s.cache.Relation.SetFollowerList(ctx, to_id, kv...) {
				Log.Errorf("set follow list to redis err")
				return resp, constants.ErrWriteCache
			}
		}
	} else {
		followeridList = s.cache.Relation.GetFollowerList(ctx, to_id)
	}

	if s.cache.Relation.FollowIsExists(ctx, user_id) == 0 {
		// 缓存中不存在登录用户粉丝列表
		userList, err := s.dao.Relation.FollowidList(ctx, user_id)
		if err != nil {
			Log.Errorf("get follow list err: %v, uid: %v", err, user_id)
		}
		if len(userList) > 0 {
			kv := make([]string, 0)
			for _, user := range userList {
				kv = append(kv, strconv.FormatInt(user, 10))
				kv = append(kv, "1")
			}
			if !s.cache.Relation.SetFollowList(ctx, user_id, kv...) {
				Log.Errorf("set follow list to redis err")
				return resp, constants.ErrWriteCache
			}
		}

	}

	// 遍历查找查询用户所关注的用户
	followerList := []*model2.User{}
	if len(followeridList) > 0 {
		for _, followerid := range followeridList {
			var follower = &model2.User{}
			if s.cache.User.IsExists(ctx, followerid) == 0 {
				// 如果要查询的用户不在缓存中
				v, err := s.dao.User.QueryUser(ctx, followerid)
				if err != nil {
					Log.Errorf("query user %v err: %v", followerid, err)
					return resp, err
				}
				follower = &model2.User{
					Id:            v.ID,
					Name:          v.Username,
					FollowCount:   &v.FollowerCount,
					FollowerCount: &v.FollowerCount,
					Avatar:        &v.Avatar,
				}

			} else {
				// 要查询的用户位于缓存中
				userModel, err := s.cache.User.GetUserMessage(ctx, followerid)
				if err != nil {
					return resp, constants.ErrReadCache
				}

				follower = cache.UserModel2User(userModel)
			}

			// 查找缓存看登录用户是否关注当前用户
			isfollow := s.cache.Relation.IsFollow(ctx, user_id, followerid)
			follower.SetIsFollow(isfollow)
			followerList = append(followerList, follower)
		}
	}

	resp.SetUserList(followerList)
	return
}

func (s *Service) FriendList(ctx context.Context, req *social.DouyinRelationFriendListRequest) (resp *social.DouyinRelationFriendListResponse, err error) {
	resp = social.NewDouyinRelationFriendListResponse()
	user_id := req.GetUserId()
	var followList []int64
	// Try get the user follower list from cache.
	// if missing, fill cache from dao.
	if s.cache.Relation.FollowIsExists(ctx, user_id) == 0 {
		followList = make([]int64, 0)
		userList, err := s.dao.Relation.FollowidList(ctx, user_id)
		if err != nil {
			Log.Errorf("get follow list err: %v, uid: %v", err, user_id)
		}
		if len(userList) > 0 {
			kv := make([]string, 0)
			for _, user := range userList {
				followList = append(followList, user)
				kv = append(kv, strconv.FormatInt(user, 10))
				kv = append(kv, "1")
			}
			if !s.cache.Relation.SetFollowList(ctx, user_id, kv...) {
				Log.Errorf("set follow list to redis err")
				return resp, constants.ErrWriteCache
			}
		}
	} else {
		followList = s.cache.Relation.GetFollowList(ctx, user_id)
	}

	var friends []*model.FriendUser
	for _, followerId := range followList {
		// Add the follower from cache
		if s.cache.Relation.FollowIsExists(ctx, followerId) == 0 {
			userList, err := s.dao.Relation.FollowidList(ctx, followerId)
			if err != nil {
				Log.Errorf("get follow list err: %v, uid: %v", err, user_id)
			}
			if len(userList) > 0 {
				kv := make([]string, 0)
				for _, user := range userList {
					kv = append(kv, strconv.FormatInt(user, 10))
					kv = append(kv, "1")
				}
				if !s.cache.Relation.SetFollowList(ctx, followerId, kv...) {
					Log.Errorf("set follow list to redis err")
					return resp, constants.ErrWriteCache
				}
			}
		}
		isFriend := s.cache.Relation.IsFollow(ctx, followerId, user_id)
		if isFriend == false {
			continue
		}
		// Try get follower UserMessage from cache
		userModel := &cache.UserModel{}
		if s.cache.User.IsExists(ctx, followerId) == 0 {
			user, err := s.dao.User.QueryUser(ctx, followerId)
			if err != nil {
				Log.Errorf("query user %v err: %v", followerId, err)
				return resp, err
			}
			userModel = &cache.UserModel{
				Id:              user.ID,
				Name:            user.Username,
				FollowCount:     user.FollowCount,
				FollowerCount:   user.FollowerCount,
				Avatar:          user.Avatar,
				BackgroundImage: user.BackgroundImage,
				Signature:       user.Signature,
				TotalFavorited:  user.TotalFavorited,
				WorkCount:       user.WorkCount,
				FavoriteCount:   user.FavoriteCount,
			}
			if !s.cache.User.SetUserMessage(ctx, userModel) {
				Log.Errorf("set user message to redis err: %v", err)
				return resp, constants.ErrWriteCache
			}
		} else {
			userModel, err = s.cache.User.GetUserMessage(ctx, followerId)
			if err != nil {
				return resp, constants.ErrReadCache
			}
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
				Id:            userModel.Id,
				Name:          userModel.Name,
				FollowCount:   &userModel.FollowCount,
				FollowerCount: &userModel.FollowerCount,
				Avatar:        &userModel.Avatar,
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
			Id:            userModel.Id,
			Name:          userModel.Name,
			FollowCount:   &userModel.FollowCount,
			FollowerCount: &userModel.FollowerCount,
			Avatar:        &userModel.Avatar,
			IsFollow:      true,
			Message:       &msg.Content,
			MsgType:       msgType,
		}

		friends = append(friends, friend)
	}
	resp.SetUserList(friends)
	return resp, nil
}
