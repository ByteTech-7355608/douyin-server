include "model.thrift"

namespace go douyin.social

// 关注操作
struct douyin_follow_action_request {
    1:required string token
    2:required i64 following_id
    3:required i64 follower_id
    4:required i32 action_type
    255: optional model.BaseReq base_req
}

struct douyin_follow_action_response {
    1:required i32 status_code
    2:optional string status_msg
}

// 关注列表
struct douyin_following_list_request {
    1:required i64 user_id
    2:required string token
    255: optional model.BaseReq base_req
}

struct douyin_following_list_response {
    1:required i32 status_code
    2:optional string status_msg
    3:required list<model.User> following_list
}

// 粉丝列表
struct douyin_follower_list_request {
    1:required i64 user_id
    2:required string token
    255: optional model.BaseReq base_req
}

struct douyin_follower_list_response {
    1:required i32 status_code
    2:optional string status_msg
    3:required list<model.User> follower_list
}

// rpc service
service SocialService {
    douyin_follow_action_response FollowAction(1:douyin_follow_action_request req)
    douyin_following_list_response FollowingList(1:douyin_following_list_request req)
    douyin_follower_list_response FollowerList(1:douyin_follower_list_request req)
}