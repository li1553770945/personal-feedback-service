// Code generated by Kitex v0.7.2. DO NOT EDIT.

package feedbackservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	feedback "github.com/li1553770945/personal-feedback-service/kitex_gen/feedback"
)

func serviceInfo() *kitex.ServiceInfo {
	return feedbackServiceServiceInfo
}

var feedbackServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "FeedbackService"
	handlerType := (*feedback.FeedbackService)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetFeedbackCategories": kitex.NewMethodInfo(getFeedbackCategoriesHandler, newFeedbackServiceGetFeedbackCategoriesArgs, newFeedbackServiceGetFeedbackCategoriesResult, false),
		"GetFeedback":           kitex.NewMethodInfo(getFeedbackHandler, newFeedbackServiceGetFeedbackArgs, newFeedbackServiceGetFeedbackResult, false),
		"AddFeedback":           kitex.NewMethodInfo(addFeedbackHandler, newFeedbackServiceAddFeedbackArgs, newFeedbackServiceAddFeedbackResult, false),
		"AddReply":              kitex.NewMethodInfo(addReplyHandler, newFeedbackServiceAddReplyArgs, newFeedbackServiceAddReplyResult, false),
		"GetReply":              kitex.NewMethodInfo(getReplyHandler, newFeedbackServiceGetReplyArgs, newFeedbackServiceGetReplyResult, false),
		"GetUnreadFeedback":     kitex.NewMethodInfo(getUnreadFeedbackHandler, newFeedbackServiceGetUnreadFeedbackArgs, newFeedbackServiceGetUnreadFeedbackResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "feedback",
		"ServiceFilePath": `idl/feedback.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.7.2",
		Extra:           extra,
	}
	return svcInfo
}

func getFeedbackCategoriesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {

	realResult := result.(*feedback.FeedbackServiceGetFeedbackCategoriesResult)
	success, err := handler.(feedback.FeedbackService).GetFeedbackCategories(ctx)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFeedbackServiceGetFeedbackCategoriesArgs() interface{} {
	return feedback.NewFeedbackServiceGetFeedbackCategoriesArgs()
}

func newFeedbackServiceGetFeedbackCategoriesResult() interface{} {
	return feedback.NewFeedbackServiceGetFeedbackCategoriesResult()
}

func getFeedbackHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*feedback.FeedbackServiceGetFeedbackArgs)
	realResult := result.(*feedback.FeedbackServiceGetFeedbackResult)
	success, err := handler.(feedback.FeedbackService).GetFeedback(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFeedbackServiceGetFeedbackArgs() interface{} {
	return feedback.NewFeedbackServiceGetFeedbackArgs()
}

func newFeedbackServiceGetFeedbackResult() interface{} {
	return feedback.NewFeedbackServiceGetFeedbackResult()
}

func addFeedbackHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*feedback.FeedbackServiceAddFeedbackArgs)
	realResult := result.(*feedback.FeedbackServiceAddFeedbackResult)
	success, err := handler.(feedback.FeedbackService).AddFeedback(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFeedbackServiceAddFeedbackArgs() interface{} {
	return feedback.NewFeedbackServiceAddFeedbackArgs()
}

func newFeedbackServiceAddFeedbackResult() interface{} {
	return feedback.NewFeedbackServiceAddFeedbackResult()
}

func addReplyHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*feedback.FeedbackServiceAddReplyArgs)
	realResult := result.(*feedback.FeedbackServiceAddReplyResult)
	success, err := handler.(feedback.FeedbackService).AddReply(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFeedbackServiceAddReplyArgs() interface{} {
	return feedback.NewFeedbackServiceAddReplyArgs()
}

func newFeedbackServiceAddReplyResult() interface{} {
	return feedback.NewFeedbackServiceAddReplyResult()
}

func getReplyHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*feedback.FeedbackServiceGetReplyArgs)
	realResult := result.(*feedback.FeedbackServiceGetReplyResult)
	success, err := handler.(feedback.FeedbackService).GetReply(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFeedbackServiceGetReplyArgs() interface{} {
	return feedback.NewFeedbackServiceGetReplyArgs()
}

func newFeedbackServiceGetReplyResult() interface{} {
	return feedback.NewFeedbackServiceGetReplyResult()
}

func getUnreadFeedbackHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {

	realResult := result.(*feedback.FeedbackServiceGetUnreadFeedbackResult)
	success, err := handler.(feedback.FeedbackService).GetUnreadFeedback(ctx)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFeedbackServiceGetUnreadFeedbackArgs() interface{} {
	return feedback.NewFeedbackServiceGetUnreadFeedbackArgs()
}

func newFeedbackServiceGetUnreadFeedbackResult() interface{} {
	return feedback.NewFeedbackServiceGetUnreadFeedbackResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetFeedbackCategories(ctx context.Context) (r *feedback.GetFeedbackCategoryResp, err error) {
	var _args feedback.FeedbackServiceGetFeedbackCategoriesArgs
	var _result feedback.FeedbackServiceGetFeedbackCategoriesResult
	if err = p.c.Call(ctx, "GetFeedbackCategories", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFeedback(ctx context.Context, req *feedback.GetFeedbackReq) (r *feedback.GetFeedbackResp, err error) {
	var _args feedback.FeedbackServiceGetFeedbackArgs
	_args.Req = req
	var _result feedback.FeedbackServiceGetFeedbackResult
	if err = p.c.Call(ctx, "GetFeedback", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AddFeedback(ctx context.Context, req *feedback.AddFeedBackReq) (r *feedback.AddFeedbackResp, err error) {
	var _args feedback.FeedbackServiceAddFeedbackArgs
	_args.Req = req
	var _result feedback.FeedbackServiceAddFeedbackResult
	if err = p.c.Call(ctx, "AddFeedback", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AddReply(ctx context.Context, req *feedback.AddReplyReq) (r *feedback.AddReplyResp, err error) {
	var _args feedback.FeedbackServiceAddReplyArgs
	_args.Req = req
	var _result feedback.FeedbackServiceAddReplyResult
	if err = p.c.Call(ctx, "AddReply", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetReply(ctx context.Context, req *feedback.GetReplyReq) (r *feedback.GetReplyResp, err error) {
	var _args feedback.FeedbackServiceGetReplyArgs
	_args.Req = req
	var _result feedback.FeedbackServiceGetReplyResult
	if err = p.c.Call(ctx, "GetReply", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUnreadFeedback(ctx context.Context) (r *feedback.GetUnreadFeedbackResp, err error) {
	var _args feedback.FeedbackServiceGetUnreadFeedbackArgs
	var _result feedback.FeedbackServiceGetUnreadFeedbackResult
	if err = p.c.Call(ctx, "GetUnreadFeedback", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
