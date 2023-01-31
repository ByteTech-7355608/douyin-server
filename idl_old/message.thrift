namespace go douyin.extra.second

struct douyin_message_chat_request {
    1: required string token  // 用户鉴权token
    2: required i64 to_use_id // 对方用户id
}

struct douyin_message_chat_response {
    1: required i32 status_code // 状态码 0-成功， 其他值-失败
    2: optional string status_msg // 返回状态描述
    3: required Message message_list // 消息列表
}

struct Message {
    1: required i64 id  // 消息id
    2: required string content  // 消息内容
    3: optional string create_time // 消息创建时间
}

struct douyin_relation_action_request {
    1: required string token    // 用户鉴权token
    2: required i64 to_user_id  // 对方用户id
    3: required i32 action_type // 1-发送消息
    4: required string content  // 消息内容
}

struct douyin_relation_action_response {
    1: required i32 status_code  // 状态码 0-成功， 其他值-失败
    2: optional string status_msg
}

service MessageService {
    douyin_message_chat_response QueryMessage(1:douyin_message_chat_request req)
    douyin_relation_action_response UpdateMessage(1:douyin_relation_action_request req)
}
