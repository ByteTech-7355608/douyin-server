package router

import (
	"ByteTech-7355608/douyin-server/cmd/api/biz/handler"

	"github.com/cloudwego/hertz/pkg/app/server"
)

// Register 注册路由
func Register(r *server.Hertz, h *handler.Handler) {

	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		{
			_comment := _douyin.Group("/comment", _commentMw()...)
			{
				_action := _comment.Group("/action", _actionMw()...)
				_action.POST("/", append(_comment_ctionMw(), h.CommentAction)...)
			}
			{
				_list := _comment.Group("/list", _listMw()...)
				_list.GET("/", append(_commentlistMw(), h.CommentList)...)
			}
		}
		{
			_favorite := _douyin.Group("/favorite", _favoriteMw()...)
			{
				_action0 := _favorite.Group("/action", _action0Mw()...)
				_action0.POST("/", append(_favorite_ctionMw(), h.FavoriteAction)...)
			}
			{
				_list0 := _favorite.Group("/list", _list0Mw()...)
				_list0.GET("/", append(_favoritelistMw(), h.FavoriteList)...)
			}
		}
		{
			_feed := _douyin.Group("/feed", _feedMw()...)
			_feed.GET("/", append(_feed0Mw(), h.Feed)...)
		}
		//{
		//	_message := _douyin.Group("/message", _messageMw()...)
		//	{
		//		_action1 := _message.Group("/action", _action1Mw()...)
		//		_action1.GET("/", append(_sendmessageMw(), h.SendMessage)...)
		//	}
		//	{
		//		_list1 := _message.Group("/list", _list1Mw()...)
		//		_list1.GET("/", append(_messagelistMw(), h.MessageList)...)
		//	}
		//}
		//{
		//	_relatioin := _douyin.Group("/relatioin", _relatioinMw()...)
		//	{
		//		_follow := _relatioin.Group("/follow", _followMw()...)
		//		{
		//			_list3 := _follow.Group("/list", _list3Mw()...)
		//			_list3.GET("/", append(_followlistMw(), h.FollowList)...)
		//		}
		//	}
		//}
		{
			_relation := _douyin.Group("/relation", _relationMw()...)
			//	{
			//		_action3 := _relation.Group("/action", _action3Mw()...)
			//		_action3.GET("/", append(_follow_ctionMw(), h.FollowAction)...)
			//	}
			{
				_follower := _relation.Group("/follower", _followerMw()...)
				{
					_list4 := _follower.Group("/list", _list4Mw()...)
					_list4.GET("/", append(_followerlistMw(), h.FollowerList)...)
				}
			}
			{
				_friend := _relation.Group("/friend", _friendMw()...)
				{
					_list5 := _friend.Group("/list", _list5Mw()...)
					_list5.GET("/", append(_friendlistMw(), h.FriendList)...)
				}
			}
		}
		{
			_user := _douyin.Group("/user", _userMw()...)
			_user.GET("/", append(_usermsgMw(), h.UserMsg)...)
			{
				_login := _user.Group("/login", _loginMw()...)
				_login.POST("/", append(_userloginMw(), h.UserLogin)...)
			}
			{
				_register := _user.Group("/register", _registerMw()...)
				_register.POST("/", append(_userregisterMw(), h.UserRegister)...)
			}
		}

		{
			_publish := _douyin.Group("/publish", _publishMw()...)
			{
				_action2 := _publish.Group("/action", _action2Mw()...)
				_action2.POST("/", append(_publish_ctionMw(), h.PublishAction)...)
			}
			{
				_list2 := _publish.Group("/list", _list2Mw()...)
				_list2.GET("/", append(_publishlistMw(), h.PublishList)...)
			}
		}
	}
}
