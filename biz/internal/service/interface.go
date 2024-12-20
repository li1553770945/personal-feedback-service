package service

import (
	"context"
	"github.com/li1553770945/personal-feedback-service/biz/internal/repo"
	"github.com/li1553770945/personal-feedback-service/kitex_gen/feedback"
	"github.com/li1553770945/personal-notify-service/kitex_gen/notify/notifyservice"
)

type FeedbackServiceImpl struct {
	Repo         repo.IRepository
	NotifyClient notifyservice.Client
}

type IFeedbackService interface {
	GetFeedbackCategories(ctx context.Context) (resp *feedback.GetFeedbackCategoryResp, err error)
	GetFeedback(ctx context.Context, req *feedback.GetFeedbackReq) (resp *feedback.GetFeedbackResp, err error)
	AddFeedback(ctx context.Context, req *feedback.AddFeedBackReq) (resp *feedback.AddFeedbackResp, err error)
	AddReply(ctx context.Context, req *feedback.AddReplyReq) (resp *feedback.AddReplyResp, err error)
	GetReply(ctx context.Context, req *feedback.GetReplyReq) (resp *feedback.GetReplyResp, err error)
	GetUnreadFeedback(ctx context.Context) (resp *feedback.GetUnreadFeedbackResp, err error)
}

func NewFeedbackService(repo repo.IRepository, notifyClient notifyservice.Client) IFeedbackService {
	return &FeedbackServiceImpl{
		Repo:         repo,
		NotifyClient: notifyClient,
	}
}
