package repo

import "github.com/li1553770945/personal-feedback-service/biz/internal/domain"

func (Repo *Repository) FindAllMessageCategory() (*[]domain.MessageCategoryEntity, error) {
	//TODO implement me
	panic("implement me")
}

func (Repo *Repository) SaveMessage(entity *domain.MessageEntity) error {
	//TODO implement me
	panic("implement me")
}

func (Repo *Repository) FindMessageByUUID(uuid string) (*domain.MessageEntity, error) {
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

func (Repo *Repository) FindMessageByID(messageId uint) (*domain.MessageEntity, error) {
	//TODO implement me
	panic("implement me")
}

func (Repo *Repository) GetUnreadMsg() (*[]domain.MessageEntity, error) {
	//TODO implement me
	panic("implement me")
}
