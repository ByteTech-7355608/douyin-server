include "model.thrift"

namespace go douyin.base

// 视频流接口
struct douyin_feed_request {
    1:optional i64 latest_time
    2:optional string token
    255:optional model.BaseReq base_req
}

struct douyin_feed_response {
    1:required i32 status_code
    2:optional string status_msg
    3:required list<model.Video> video_list
    4:optional i64 next_time
}

// 用户注册接口
struct douyin_user_register_request {
    1:required string username
    2:required string password
}

struct douyin_user_register_response {
    1:required i32 status_code
    2:optional string status_msg
    3:required i64 user_id
    4:required string token
}

// 用户登录接口
struct douyin_user_login_request {
    1:required string username
    2:required string password
}

struct douyin_user_login_response {
    1:required i32 status_code
    2:optional string status_msg
    3:required i64 user_id
    4:required string token
}

// 用户信息
struct douyin_user_request {
    1:required i64 user_id
    2:required string token
    255:optional model.BaseReq base_req
}

struct douyin_user_response {
    1:required i32 status_code
    2:optional string status_msg
    3:required model.User user
}

// 视频投稿
struct douyin_publish_action_request {
    1:required string token
    2:required binary data
    3:required string title
    4:optional string play_url
    5:optional string cover_url
    255:optional model.BaseReq base_req
}

struct douyin_publish_action_response {
    1:required i32 status_code
    2:optional string status_msg
}

// 发布列表
struct douyin_publish_list_request {
    1:required i64 user_id
    2:required string token
    255:optional model.BaseReq base_req
}

struct douyin_publish_list_response {
    1:required i32 status_code
    2:optional string status_msg
    3:required list<model.Video> video_list
}

// rpc service
service BaseService {
    douyin_feed_response Feed(1:douyin_feed_request req)
    douyin_user_register_response UserRegister(1:douyin_user_register_request req)
    douyin_user_login_response UserLogin(1:douyin_user_login_request req)
    douyin_user_response UserMsg(1:douyin_user_request req)
    douyin_publish_action_response PublishAction(1:douyin_publish_action_request req)
    douyin_publish_list_response PublishList(1:douyin_publish_list_request req)
}