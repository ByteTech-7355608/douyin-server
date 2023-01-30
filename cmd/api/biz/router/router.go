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
		//_douyin.GET("/feed", append(_feedMw(), h.Feed)...)
		//_douyin.GET("/user", append(_usermsgMw(), h.UserMsg)...)
		//{
		//	_comment := _douyin.Group("/comment", _commentMw()...)
		//	_comment.POST("/action", append(_comment_ctionMw(), h.CommentAction)...)
		//	_comment.GET("/list", append(_commentlistMw(), h.CommentList)...)
		//}
		//{
		//	_favorite := _douyin.Group("/favorite", _favoriteMw()...)
		//	_favorite.POST("/action", append(_favorite_ctionMw(), h.FavoriteAction)...)
		//	_favorite.GET("/list", append(_favoritelistMw(), h.FavoriteList)...)
		//}
		//{
		//	_publish := _douyin.Group("/publish", _publishMw()...)
		//	_publish.POST("/action", append(_publish_ctionMw(), h.PublishAction)...)
		//	_publish.GET("/list", append(_publishlistMw(), h.PublishList)...)
		//}
		{
			_user := _douyin.Group("/user", _userMw()...)
			//	_user.POST("/login", append(_userloginMw(), h.UserLogin)...)
			_user.POST("/register", append(_userregisterMw(), h.UserRegister)...)
		}
	}
}
