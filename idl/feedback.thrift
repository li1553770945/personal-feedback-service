namespace go feedback
include "base.thrift"

struct GetFeedbackCategoryRespData {
    1: required i64 id
    2: required string name      // 分类名称
}

struct GetFeedbackCategoryResp{
    1: required base.BaseResp baseResp
    2: optional list<GetFeedbackCategoryRespData> data;
}

// 消息实体
struct AddFeedBackReq {
  1: required i64 category_id                  // 分类ID
  3: required string title                     // 消息标题
  4: required string content                  // 消息内容
  5: required string name                      // 发信人名称
  6: optional string contact                   // 发信人联系方式
}

struct AddFeedbackResp {
  1: required base.BaseResp baseResp
  2: optional string uuid
}


struct GetFeedbackReq {
    1: required string uuid
}

struct GetFeedbackResp{
    1: required base.BaseResp baseResp
    2: required i64 id
    3: required string title                     // 消息标题
    4: required string content                  // 消息内容
    5: required string name                      // 发信人名称
    6: required string contact                   // 发信人联系方式
    7: required i64 createdAt
    8: required string category
}

struct AddReplyReq{
    1: required i64 feedback_id
    2: required string content
}

struct AddReplyResp{
    1: required base.BaseResp baseResp
}
struct GetReplyReq {
    1: required string feedbackUuid
}

struct GetReplyResp {
    1: required base.BaseResp baseResp
    2: optional string content       // 回复内容
    3: optional i64 createdAt
}

struct UnreadFeedbackData{
    1: required i64 id
    2: required string title
    3: required string name
    4: required i64 createdAt
    5: required string uuid
}
struct GetUnreadFeedbackResp{
    1: required base.BaseResp baseResp
    2: optional list<UnreadFeedbackData> data
}

service FeedbackService {
    GetFeedbackCategoryResp GetFeedbackCategories()
    GetFeedbackResp GetFeedback(GetFeedbackReq req)
    AddFeedbackResp AddFeedback(AddFeedBackReq req)
    AddReplyResp AddReply(AddReplyReq req)
    GetReplyResp GetReply(GetReplyReq req)
    GetUnreadFeedbackResp GetUnreadFeedback()

}
