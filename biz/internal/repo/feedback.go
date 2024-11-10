package repo

import "github.com/li1553770945/personal-feedback-service/biz/internal/domain"

func (Repo *Repository) FindAllMessageCategory() (*[]domain.FeedbackCategoryEntity, error) {
	//TODO implement me
	panic("implement me")
}

func (Repo *Repository) SaveMessage(entity *domain.FeedbackEntity) error {
	//TODO implement me
	panic("implement me")
}

func (Repo *Repository) FindFeedbackByUUID(uuid string) (*domain.FeedbackEntity, error) {
	//TODO implement me
	panic("implement me")
}

func (Repo *Repository) FindReplyByMessageID(messageId uint) (*domain.ReplyEntity, error) {
	//TODO implement me
	panic("implement me")
}

func (Repo *Repository) SaveReply(entity *domain.ReplyEntity) error {
	//TODO implement me
	panic("implement me")
}

func (Repo *Repository) FindMessageByID(messageId uint) (*domain.FeedbackEntity, error) {
	//TODO implement me
	panic("implement me")
}

func (Repo *Repository) GetUnreadMsg() (*[]domain.FeedbackEntity, error) {
	//TODO implement me
	panic("implement me")
}
