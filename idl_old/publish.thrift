namespace go douyin.core

struct douyin_publish_action_request {
    1:required string token // 用户鉴权
    2:required byte data    // 视频数据
    3:required string title // 视频标题
}

struct douyin_publish_action_response {
    1:required i32 status_code //状态码， 0-成功， 其他数-失败
    2:optional string status_msg
}

service NoteService {
    douyin_publish_action_response UpdateVideo(1:douyin_publish_action_request req)
}