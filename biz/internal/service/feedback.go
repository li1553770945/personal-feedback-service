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
)

func (s *FeedbackServiceImpl) GetFeedbackCategories(ctx context.Context) (resp *feedback.GetFeedbackCategoryResp, err error) {
	categories, err := s.Repo.FindAllMessageCategory()
	if err != nil {
		return &feedback.GetFeedbackCategoryResp{
			BaseResp: &base.BaseResp{Code: constant.SystemError, Message: fmt.Sprintf("获取消息类别数据库失败:%v", err.Error())},
		}, nil
	}

	resp = converter.FeedbackCategoryEntityToDTO(categories)
	return resp, nil
}

// 获取反馈详情
func (s *FeedbackServiceImpl) GetFeedback(ctx context.Context, req *feedback.GetFeedbackReq) (resp *feedback.GetFeedbackResp, err error) {
	feedbackData, err := s.Repo.FindFeedbackByUUID(req.Uuid)
	if err != nil {
		return &feedback.GetFeedbackResp{
			BaseResp: &base.BaseResp{Code: constant.SystemError, Message: err.Error()},
		}, nil
	}
	if feedbackData.ID == 0 {
		return &feedback.GetFeedbackResp{
			BaseResp: &base.BaseResp{Code: constant.NotFound, Message: "未找到相关建议"},
		}, nil
	}

	resp = converter.FeedbackEntityToDto(feedbackData)
	return resp, nil
}

// 保存反馈信息
func (s *FeedbackServiceImpl) AddFeedback(ctx context.Context, req *feedback.AddFeedBackReq) (resp *feedback.AddFeedbackResp, err error) {
	entity := domain.FeedbackEntity{
		CategoryID: int(req.CategoryId),
		Title:      req.Title,
		Content:    req.Content,
		Name:       req.Name,
		Contact:    *req.Contact,
		HaveRead:   false,
		UUID:       uuid.New().String(),
	}

	err = s.Repo.SaveFeedback(&entity)
	if err != nil {
		return &feedback.AddFeedbackResp{
			BaseResp: &base.BaseResp{Code: constant.SystemError, Message: err.Error()},
		}, nil
	}

	klog.CtxInfof(ctx, "新增一个反馈，反馈id：%d  \n类别：%s  \n署名：%s  \n联系方式：%s", entity.ID, entity.Category.Name, entity.Name, entity.Contact)

	return &feedback.AddFeedbackResp{
		BaseResp: &base.BaseResp{Code: 0},
		Uuid:     &entity.UUID,
	}, nil
}

// GetReply 获取回复内容
func (s *FeedbackServiceImpl) GetReply(ctx context.Context, req *feedback.GetReplyReq) (resp *feedback.GetReplyResp, err error) {
	msg, err := s.Repo.FindFeedbackByID(uint(req.FeedbackId))
	if err != nil {
		return &feedback.GetReplyResp{
			BaseResp: &base.BaseResp{Code: constant.SystemError, Message: err.Error()},
		}, nil
	}
	if msg.ID == 0 {
		return &feedback.GetReplyResp{
			BaseResp: &base.BaseResp{Code: constant.NotFound, Message: "未找到相关建议"},
		}, nil
	}

	reply, err := s.Repo.FindReplyByFeedbackID(uint(msg.ID))
	if err != nil {
		return &feedback.GetReplyResp{
			BaseResp: &base.BaseResp{Code: constant.SystemError, Message: err.Error()},
		}, nil
	}
	if reply.ID == 0 {
		return &feedback.GetReplyResp{
			BaseResp: &base.BaseResp{Code: constant.NotFound, Message: "该建议暂未回复"},
		}, nil
	}

	reply.HaveRead = true
	err = s.Repo.SaveReply(reply)
	if err != nil {
		return &feedback.GetReplyResp{
			BaseResp: &base.BaseResp{Code: constant.SystemError, Message: "保存失败：" + err.Error()},
		}, nil
	}

	resp = &feedback.GetReplyResp{
		BaseResp: &base.BaseResp{Code: 0},
		Content:  &reply.Content,
	}
	return resp, nil
}

// 添加回复
func (s *FeedbackServiceImpl) AddReply(ctx context.Context, req *feedback.AddReplyReq) (resp *feedback.AddReplyResp, err error) {
	reply := domain.ReplyEntity{
		Content:   req.Content,
		MessageID: uint(req.MessageId),
	}

	// 获取关联消息并设置已读
	msg, err := s.Repo.FindFeedbackByID(uint(req.MessageId))
	if err != nil {
		return &feedback.AddReplyResp{
			BaseResp: &base.BaseResp{Code: constant.SystemError, Message: err.Error()},
		}, nil
	}
	if msg.ID == 0 {
		return &feedback.AddReplyResp{
			BaseResp: &base.BaseResp{Code: 4004, Message: "未找到相关建议"},
		}, nil
	}

	msg.HaveRead = true
	err = s.Repo.SaveFeedback(msg)
	if err != nil {
		return &feedback.AddReplyResp{
			BaseResp: &base.BaseResp{Code: constant.SystemError, Message: "保存失败：" + err.Error()},
		}, nil
	}

	err = s.Repo.SaveReply(&reply)
	if err != nil {
		return &feedback.AddReplyResp{
			BaseResp: &base.BaseResp{Code: constant.SystemError, Message: "操作失败：" + err.Error()},
		}, nil
	}

	return &feedback.AddReplyResp{
		BaseResp: &base.BaseResp{Code: 0},
	}, nil
}

func (s *FeedbackServiceImpl) GetUnreadFeedback(ctx context.Context) (resp *feedback.GetUnreadFeedbackResp, err error) {
	unreadMessages, err := s.Repo.GetUnreadFeedback()
	if err != nil {
		return &feedback.GetUnreadFeedbackResp{
			BaseResp: &base.BaseResp{
				Code:    constant.SystemError,
				Message: "获取未读反馈失败：" + err.Error(),
			},
		}, err
	}

	resp = converter.UnreadFeedbackEntitiesToDTO(unreadMessages)
	return resp, nil
}
