namespace go douyin.model

struct User {
    1:required i64 id
    2:required string name
    3:optional i64 follow_count
    4:optional i64 follower_count
    5:required bool is_follow
}

struct FriendUser {
    1: User user
    2: optional string message // 和该好友的最新聊天记录
    3: required i64 msgType    // 0-当前用户接收的消息， 1-当前用户发送的消息
}

struct Video {
    1:required i64 id
    2:required User author
    3:required string play_url
    4:required string cover_url
    5:required i64 favorite_count
    6:required i64 comment_count
    7:required bool is_favorite
    8:required string title
}

struct Comment {
    1:required i64 id
    2:required User user
    3:required string content
    4:required string create_date
}

struct BaseReq {
    1:optional i64 user_id
    2:optional string username
}

// 消息
struct Message {
    1: required i64 id             // 消息id
    2: required i64 to_user_id     // 消息接收者
    3: required i64 from_user_id   // 消息发送者
    4: required string content     // 消息内容
    5: optional string create_time // 消息创建时间
}