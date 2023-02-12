package social

import (
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/social"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"context"
)

const (
	KAddType    = 1
	KDeleteType = 2
)

func (s *Service) RelationAction(ctx context.Context, req *social.DouyinFollowActionRequest) (resp *social.DouyinFollowActionResponse, err error) {
	resp = social.NewDouyinFollowActionResponse()
	concerner_id := req.FollowingId
	concerned_id := req.FollowerId
	switch req.ActionType {
	case KAddType:
		flag, err := s.dao.Relation.IsUserFollowed(ctx, concerner_id, concerned_id)
		if flag == false && err != nil {
			//没有此记录
			err = s.dao.Relation.AddRelation(ctx, concerner_id, concerned_id)
			if err != nil {
				Log.Errorf("add comment err: %v", err)
				return nil, err
			}
		} else {
			//存在该记录
			err = s.dao.Relation.UpdatedRelation()
			if err != nil {
				Log.Errorf("add comment err: %v", err)
				return nil, err
			}

		}

	case KDeleteType:
		err = s.dao.Comment.DeleteComment(ctx, req)
		if err != nil {
			Log.Errorf("delete comment err: %v", err)
			return nil, err
		}
	}
	return
}
