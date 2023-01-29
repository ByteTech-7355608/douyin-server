namespace go douyin.extra.second

struct User {
    1:required i64 id              //用户id
    2:required string name         //用户名称
    3:optional i64 follow_count    //关注总数
    4:optional i64 follower_count  //粉丝总数
    5:required bool is_follow      // true-已关注，false-未关注
}

struct douyin_relation_friend_list_request {
    1:required i64 user_id   // 用户id
    2:required string token  // 用户鉴权token
}

struct douyin_relation_friend_list_response {
    1:required i32 status_code     //状态码，0- 成功，其他值-失败
    2:optional string status_msg   //返回状态描述
    3:required User user_list      //用户列表
}

service NoteService {
    douyin_relation_friend_list_response QueryFriend(1:douyin_relation_friend_list_request req)
}