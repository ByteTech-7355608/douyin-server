include "model.thrift"

namespace go douyin.social

// 关注操作
struct douyin_follow_action_request {
    1:required string token
    2:required i64 to_user_id
    3:required i32 action_type
    255:optional model.BaseReq base_req
}

struct douyin_follow_action_response {
    1:required i32 status_code
    2:optional string status_msg
}

// 关注列表
struct douyin_following_list_request {
    1:required i64 user_id
    2:required string token
    255:optional model.BaseReq base_req
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
    255:optional model.BaseReq base_req
}

struct douyin_follower_list_response {
    1:required i32 status_code
    2:optional string status_msg
    3:required list<model.User> follower_list
}

// 好友列表
struct douyin_relation_friend_list_request {
    1:required i64 user_id   // 用户id
    2:required string token  // 用户鉴权token
    255:optional model.BaseReq base_req
}

struct douyin_relation_friend_list_response {
    1:required i32 status_code                      //状态码，0- 成功，其他值-失败
    2:optional string status_msg                    //返回状态描述
    3:required list<model.FriendUser> user_list     //好友用户列表，好友指的是相互关注的用户
}

// 查询消息
struct douyin_message_chat_request {
    1:required string token  // 用户鉴权token
    2:required i64 to_user_id // 对方用户id
    255:optional model.BaseReq base_req
}

struct douyin_message_chat_response {
    1:required i32 status_code                  // 状态码 0-成功， 其他值-失败
    2:optional string status_msg                // 返回状态描述
    3:required list<model.Message> message_list // 消息列表
}

// 发送消息
struct douyin_message_action_request {
    1: required string token    // 用户鉴权token
    2: required i64 to_user_id  // 对方用户id
    3: required i32 action_type // 1-发送消息
    4: required string content  // 消息内容
    255: optional model.BaseReq base_req

}

struct douyin_message_action_response {
    1: required i32 status_code  // 状态码 0-成功， 其他值-失败
    2: optional string status_msg
}

// rpc service
service SocialService {
    douyin_follow_action_response FollowAction(1:douyin_follow_action_request req)
    douyin_following_list_response FollowList(1:douyin_following_list_request req)
    douyin_follower_list_response FollowerList(1:douyin_follower_list_request req)
    douyin_relation_friend_list_response FriendList(1:douyin_relation_friend_list_request req)
    douyin_message_chat_response MessageList(1:douyin_message_chat_request req)
    douyin_message_action_response SendMessage(1:douyin_message_action_request req)
}