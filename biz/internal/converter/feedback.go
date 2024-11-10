package converter

import (
	"github.com/li1553770945/personal-feedback-service/biz/constant"
	"github.com/li1553770945/personal-feedback-service/biz/internal/domain"
	"github.com/li1553770945/personal-feedback-service/kitex_gen/base"
	"github.com/li1553770945/personal-feedback-service/kitex_gen/feedback"
)

func GetSuccessBaseResp() *base.BaseResp {
	return &base.BaseResp{
		Code: constant.Success,
	}
}

func FeedbackCategoryEntityToDTO(entities *[]domain.FeedbackCategoryEntity) *feedback.GetFeedbackCategoryResp {
	var data []*feedback.GetFeedbackCategoryRespData
	for _, entity := range *entities {
		data = append(data, &feedback.GetFeedbackCategoryRespData{
			Id:    entity.ID, // 假设 ID 是 uint 类型，需要转换为 int64
			Name:  entity.Name,
			Value: entity.Value,
		})
	}

	// 返回 GetFeedbackCategoryResp 结构体
	return &feedback.GetFeedbackCategoryResp{
		BaseResp: GetSuccessBaseResp(),
		Data:     data,
	}
}

func FeedbackEntityToDto(entity *domain.FeedbackEntity) *feedback.GetFeedbackResp {
	return &feedback.GetFeedbackResp{
		BaseResp: GetSuccessBaseResp(),
		Title:    entity.Title,
		Content:  entity.Content,
		Name:     entity.Name,
		Contact:  &entity.Contact,
	}
}

func UnreadFeedbackEntitiesToDTO(entities *[]domain.FeedbackEntity) *feedback.GetUnreadFeedbackResp {
	var data []*feedback.UnreadFeedbackData
	for _, entity := range *entities {
		data = append(data, &feedback.UnreadFeedbackData{
			Id:    int64(entity.ID), // 假设 ID 是 uint 类型，需要转换为 int64
			Title: entity.Title,
		})
	}
	return &feedback.GetUnreadFeedbackResp{
		BaseResp: GetSuccessBaseResp(),
		Data:     data,
	}
}
