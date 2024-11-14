package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"github.com/li1553770945/personal-feedback-service/biz/constant"
	"github.com/li1553770945/personal-feedback-service/biz/internal/converter"
	"github.com/li1553770945/personal-feedback-service/biz/internal/domain"
	"github.com/li1553770945/personal-feedback-service/kitex_gen/base"
	"github.com/li1553770945/personal-feedback-service/kitex_gen/feedback"
	"github.com/li1553770945/personal-notify-service/kitex_gen/notify"
)

func (s *FeedbackServiceImpl) GetFeedbackCategories(ctx context.Context) (resp *feedback.GetFeedbackCategoryResp, err error) {
	klog.CtxInfof(ctx, "开始获取所有反馈类别")
	categories, err := s.Repo.FindAllFeedbackCategories()
	if err != nil {
		klog.CtxErrorf(ctx, "获取反馈类别数据库失败: %v", err)
		return &feedback.GetFeedbackCategoryResp{
			BaseResp: &base.BaseResp{Code: constant.SystemError, Message: fmt.Sprintf("获取消息类别数据库失败: %v", err.Error())},
		}, nil
	}

	resp = converter.FeedbackCategoryEntityToDTO(categories)
	fmt.Println(resp.Data)
	klog.CtxInfof(ctx, "成功获取反馈类别，共 %d 条记录", len(*categories))
	return resp, nil
}

func (s *FeedbackServiceImpl) GetFeedback(ctx context.Context, req *feedback.GetFeedbackReq) (resp *feedback.GetFeedbackResp, err error) {
	klog.CtxInfof(ctx, "开始获取反馈详情，UUID: %s", req.Uuid)
	feedbackData, err := s.Repo.FindFeedbackByUUID(req.Uuid)
	if err != nil {
		klog.CtxErrorf(ctx, "根据 UUID 获取反馈失败，UUID: %s，错误: %v", req.Uuid, err)
		return &feedback.GetFeedbackResp{
			BaseResp: &base.BaseResp{Code: constant.SystemError, Message: err.Error()},
		}, nil
	}
	if feedbackData.ID == 0 {
		klog.CtxInfof(ctx, "未找到相关反馈，UUID: %s", req.Uuid)
		return &feedback.GetFeedbackResp{
			BaseResp: &base.BaseResp{Code: constant.NotFound, Message: "未找到相关建议"},
		}, nil
	}

	resp = converter.FeedbackEntityToDto(feedbackData)
	klog.CtxInfof(ctx, "反馈详情获取成功，UUID: %s", req.Uuid)
	return resp, nil
}

func (s *FeedbackServiceImpl) AddFeedback(ctx context.Context, req *feedback.AddFeedBackReq) (resp *feedback.AddFeedbackResp, err error) {
	klog.CtxInfof(ctx, "开始新增反馈信息，类别ID: %d，标题: %s", req.CategoryId, req.Title)
	entity := converter.AddFeedbackReqToEntity(req, uuid.New().String())

	err = s.Repo.SaveFeedback(&entity)
	if err != nil {
		klog.CtxErrorf(ctx, "保存反馈信息失败: %v", err)
		return &feedback.AddFeedbackResp{
			BaseResp: &base.BaseResp{Code: constant.SystemError, Message: err.Error()},
		}, nil
	}

	klog.CtxInfof(ctx, "新增反馈成功，反馈ID: %d，UUID: %s", entity.ID, entity.UUID)
	message := fmt.Sprintf("消息id：%d  \n类别：%s  \n署名：%s  \n联系方式：%s  \n内容：%s",
		entity.ID,
		entity.Category.Name,
		entity.Name,
		entity.Contact,
		entity.Content)

	notifyReq := notify.SendNotifyReq{
		Platform: "wechat",
		Title:    "收到一条新的反馈",
		Message:  message,
	}
	notifyResp, err := s.NotifyClient.SendNotify(ctx, &notifyReq)
	if err != nil {
		klog.CtxErrorf(ctx, "发送消息失败: %v", err)
	}
	if notifyResp.BaseResp.Code != constant.Success {
		klog.CtxErrorf(ctx, "发送消息失败:%s", notifyResp.BaseResp.Message)
	}
	if err == nil && notifyResp.BaseResp.Code == constant.Success {
		klog.CtxInfof(ctx, "发送消息成功")
	}

	return &feedback.AddFeedbackResp{
		BaseResp: &base.BaseResp{Code: 0},
		Uuid:     &entity.UUID,
	}, nil
}

func (s *FeedbackServiceImpl) GetReply(ctx context.Context, req *feedback.GetReplyReq) (resp *feedback.GetReplyResp, err error) {
	klog.CtxInfof(ctx, "开始获取反馈回复，反馈UIID: %s", req.FeedbackUuid)
	msg, err := s.Repo.FindFeedbackByUUID(req.FeedbackUuid)
	if err != nil {
		klog.CtxErrorf(ctx, "获取反馈内容失败，反馈UUIID: %s，错误: %v", req.FeedbackUuid, err)
		return &feedback.GetReplyResp{
			BaseResp: &base.BaseResp{Code: constant.SystemError, Message: err.Error()},
		}, nil
	}
	if msg.ID == 0 {
		klog.CtxInfof(ctx, "未找到相关反馈，反馈UUIID: %s", req.FeedbackUuid)
		return &feedback.GetReplyResp{
			BaseResp: &base.BaseResp{Code: constant.NotFound, Message: "未找到相关建议"},
		}, nil
	}

	reply, err := s.Repo.FindReplyByFeedbackID(uint(msg.ID))
	if err != nil {
		klog.CtxErrorf(ctx, "获取回复失败，反馈UUID: %s，错误: %v", req.FeedbackUuid, err)
		return &feedback.GetReplyResp{
			BaseResp: &base.BaseResp{Code: constant.SystemError, Message: err.Error()},
		}, nil
	}
	if reply.ID == 0 {
		klog.CtxInfof(ctx, "该反馈暂未回复，反馈UUID: %s", req.FeedbackUuid)
		return &feedback.GetReplyResp{
			BaseResp: &base.BaseResp{Code: constant.NotFound, Message: "该建议暂未回复"},
		}, nil
	}

	reply.HaveRead = true
	err = s.Repo.SaveReply(reply)
	if err != nil {
		klog.CtxErrorf(ctx, "保存回复读取状态失败，反馈ID: %s，错误: %s", req.FeedbackUuid, err)
		return &feedback.GetReplyResp{
			BaseResp: &base.BaseResp{Code: constant.SystemError, Message: "保存失败：" + err.Error()},
		}, nil
	}

	klog.CtxInfof(ctx, "获取反馈回复成功，反馈ID: %v", reply.ID)
	createdAt := reply.CreatedAt.Unix()
	resp = &feedback.GetReplyResp{
		BaseResp:  &base.BaseResp{Code: 0},
		Content:   &reply.Content,
		CreatedAt: &createdAt,
	}
	return resp, nil
}

func (s *FeedbackServiceImpl) AddReply(ctx context.Context, req *feedback.AddReplyReq) (resp *feedback.AddReplyResp, err error) {
	klog.CtxInfof(ctx, "开始添加回复，反馈ID: %d", req.FeedbackId)
	reply := domain.ReplyEntity{
		Content:   req.Content,
		MessageID: uint(req.FeedbackId),
	}

	feedbackByID, err := s.Repo.FindFeedbackByID(uint(req.FeedbackId))
	if err != nil {
		klog.CtxErrorf(ctx, "获取反馈失败，反馈ID: %d，错误: %v", req.FeedbackId, err)
		return &feedback.AddReplyResp{
			BaseResp: &base.BaseResp{Code: constant.SystemError, Message: err.Error()},
		}, nil
	}
	if feedbackByID.ID == 0 {
		klog.CtxInfof(ctx, "未找到相关反馈，反馈ID: %d", req.FeedbackId)
		return &feedback.AddReplyResp{
			BaseResp: &base.BaseResp{Code: constant.NotFound, Message: "未找到相关建议"},
		}, nil
	}
	existReply, err := s.Repo.FindReplyByFeedbackID(uint(req.FeedbackId))
	if err != nil {
		klog.CtxErrorf(ctx, "获取回复失败，反馈ID: %d，错误: %v", req.FeedbackId, err)
		return &feedback.AddReplyResp{
			BaseResp: &base.BaseResp{Code: constant.SystemError, Message: err.Error()},
		}, nil
	}
	if existReply.ID != 0 {
		klog.CtxErrorf(ctx, "反馈已经回复不可重复回复，反馈ID: %d，错误: %v", req.FeedbackId, err)
		return &feedback.AddReplyResp{
			BaseResp: &base.BaseResp{Code: constant.InvalidInput, Message: "反馈已经回复不可重复回复"},
		}, nil
	}
	feedbackByID.HaveRead = true
	err = s.Repo.SaveFeedback(feedbackByID)
	if err != nil {
		klog.CtxErrorf(ctx, "更新反馈已读状态失败，反馈ID: %d，错误: %v", req.FeedbackId, err)
		return &feedback.AddReplyResp{
			BaseResp: &base.BaseResp{Code: constant.SystemError, Message: "保存失败：" + err.Error()},
		}, nil
	}

	err = s.Repo.SaveReply(&reply)
	if err != nil {
		klog.CtxErrorf(ctx, "保存回复失败，反馈ID: %d，错误: %v", req.FeedbackId, err)
		return &feedback.AddReplyResp{
			BaseResp: &base.BaseResp{Code: constant.SystemError, Message: "操作失败：" + err.Error()},
		}, nil
	}

	klog.CtxInfof(ctx, "添加回复成功，反馈ID: %d", req.FeedbackId)
	return &feedback.AddReplyResp{
		BaseResp: &base.BaseResp{Code: 0},
	}, nil
}

func (s *FeedbackServiceImpl) GetUnreadFeedback(ctx context.Context) (resp *feedback.GetUnreadFeedbackResp, err error) {
	klog.CtxInfof(ctx, "开始获取未读反馈列表")
	unreadMessages, err := s.Repo.GetUnreadFeedback()
	if err != nil {
		klog.CtxErrorf(ctx, "获取未读反馈失败: %v", err)
		return &feedback.GetUnreadFeedbackResp{
			BaseResp: &base.BaseResp{
				Code:    constant.SystemError,
				Message: "获取未读反馈失败：" + err.Error(),
			},
		}, err
	}

	resp = converter.UnreadFeedbackEntitiesToDTO(unreadMessages)
	klog.CtxInfof(ctx, "未读反馈列表获取成功，共 %d 条记录", len(*unreadMessages))
	return resp, nil
}
