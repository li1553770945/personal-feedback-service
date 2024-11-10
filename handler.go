package main

import (
	"context"
	"github.com/li1553770945/personal-feedback-service/biz/infra/container"
	feedback "github.com/li1553770945/personal-feedback-service/kitex_gen/feedback"
)

// FeedbackServiceImpl implements the last service interface defined in the IDL.
type FeedbackServiceImpl struct{}

// GetFeedbackCategories implements the FeedbackServiceImpl interface.
func (s *FeedbackServiceImpl) GetFeedbackCategories(ctx context.Context) (resp *feedback.GetFeedbackCategoryResp, err error) {
	APP := container.GetGlobalContainer()
	resp, err = APP.FeedbackService.GetFeedbackCategories(ctx)
	return
}

// GetFeedback implements the FeedbackServiceImpl interface.
func (s *FeedbackServiceImpl) GetFeedback(ctx context.Context, req *feedback.GetFeedbackReq) (resp *feedback.GetFeedbackResp, err error) {
	APP := container.GetGlobalContainer()
	resp, err = APP.FeedbackService.GetFeedback(ctx, req)
	return
}

// AddFeedback implements the FeedbackServiceImpl interface.
func (s *FeedbackServiceImpl) AddFeedback(ctx context.Context, req *feedback.AddFeedBackReq) (resp *feedback.AddFeedbackResp, err error) {
	APP := container.GetGlobalContainer()
	resp, err = APP.FeedbackService.AddFeedback(ctx, req)
	return
}

// AddReply implements the FeedbackServiceImpl interface.
func (s *FeedbackServiceImpl) AddReply(ctx context.Context, req *feedback.AddReplyReq) (resp *feedback.AddReplyResp, err error) {
	APP := container.GetGlobalContainer()
	resp, err = APP.FeedbackService.AddReply(ctx, req)
	return
}

// GetReply implements the FeedbackServiceImpl interface.
func (s *FeedbackServiceImpl) GetReply(ctx context.Context, req *feedback.GetReplyReq) (resp *feedback.GetReplyResp, err error) {
	APP := container.GetGlobalContainer()
	resp, err = APP.FeedbackService.GetReply(ctx, req)
	return
}

// GetUnreadFeedback implements the FeedbackServiceImpl interface.
func (s *FeedbackServiceImpl) GetUnreadFeedback(ctx context.Context) (resp *feedback.GetUnreadFeedbackResp, err error) {
	// TODO: Your code here...
	App := container.GetGlobalContainer()
	resp, err = App.FeedbackService.GetUnreadFeedback(ctx)
	return
}
