package main

import (
	"context"
	feedback "github.com/li1553770945/personal-feedback-service/kitex_gen/feedback"
)

// FeedbackServiceImpl implements the last service interface defined in the IDL.
type FeedbackServiceImpl struct{}

// GetFeedbackCategories implements the FeedbackServiceImpl interface.
func (s *FeedbackServiceImpl) GetFeedbackCategories(ctx context.Context) (resp *feedback.GetFeedbackCategoryResp, err error) {
	// TODO: Your code here...
	return
}

// GetFeedback implements the FeedbackServiceImpl interface.
func (s *FeedbackServiceImpl) GetFeedback(ctx context.Context, req *feedback.GetFeedbackReq) (resp *feedback.GetFeedbackResp, err error) {
	// TODO: Your code here...
	return
}

// AddFeedback implements the FeedbackServiceImpl interface.
func (s *FeedbackServiceImpl) AddFeedback(ctx context.Context, req *feedback.AddFeedBackReq) (resp *feedback.AddFeedbackResp, err error) {
	// TODO: Your code here...
	return
}

// AddReply implements the FeedbackServiceImpl interface.
func (s *FeedbackServiceImpl) AddReply(ctx context.Context, req *feedback.AddReplyReq) (resp *feedback.AddReplyResp, err error) {
	// TODO: Your code here...
	return
}

// GetReply implements the FeedbackServiceImpl interface.
func (s *FeedbackServiceImpl) GetReply(ctx context.Context, req *feedback.GetReplyReq) (resp *feedback.GetReplyResp, err error) {
	// TODO: Your code here...
	return
}
