package service

import (
	"context"
	"github.com/li1553770945/personal-feedback-service/kitex_gen/feedback"
)

func (f FeedbackServiceImpl) GetFeedbackCategories(ctx context.Context) (resp *feedback.GetFeedbackCategoryResp, err error) {
	//TODO implement me
	panic("implement me")
}

func (f FeedbackServiceImpl) GetFeedback(ctx context.Context, req *feedback.GetFeedbackReq) (resp *feedback.GetFeedbackResp, err error) {
	//TODO implement me
	panic("implement me")
}

func (f FeedbackServiceImpl) AddFeedback(ctx context.Context, req *feedback.AddFeedBackReq) (resp *feedback.AddFeedbackResp, err error) {
	//TODO implement me
	panic("implement me")
}

func (f FeedbackServiceImpl) AddReply(ctx context.Context, req *feedback.AddReplyReq) (resp *feedback.AddReplyResp, err error) {
	//TODO implement me
	panic("implement me")
}

func (f FeedbackServiceImpl) GetReply(ctx context.Context, req *feedback.GetReplyReq) (resp *feedback.GetReplyResp, err error) {
	//TODO implement me
	panic("implement me")
}
