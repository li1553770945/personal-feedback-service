package repo

import "github.com/li1553770945/personal-feedback-service/biz/internal/domain"

func (Repo *Repository) FindAllFeedbackCategories() (*[]domain.FeedbackCategoryEntity, error) {
	var entity []domain.FeedbackCategoryEntity
	err := Repo.DB.Where("can_use = ?", true).Find(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (Repo *Repository) SaveFeedback(entity *domain.FeedbackEntity) error {
	if entity.ID == 0 {
		err := Repo.DB.Create(&entity).Error
		Repo.DB.Preload("Category").Find(&entity)
		return err
	} else {
		err := Repo.DB.Save(&entity).Error
		return err
	}
}

func (Repo *Repository) FindFeedbackByUUID(uuid string) (*domain.FeedbackEntity, error) {
	var entity domain.FeedbackEntity
	err := Repo.DB.Preload("Category").Where("uuid = ?", uuid).Find(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (Repo *Repository) FindReplyByFeedbackID(feedbackId uint) (*domain.ReplyEntity, error) {
	var entity domain.ReplyEntity
	err := Repo.DB.Where("message_id = ?", feedbackId).Find(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (Repo *Repository) SaveReply(entity *domain.ReplyEntity) error {
	if entity.ID == 0 {
		err := Repo.DB.Create(&entity).Error
		return err
	} else {
		err := Repo.DB.Save(&entity).Error
		return err
	}
}

func (Repo *Repository) FindFeedbackByID(feedbackId uint) (*domain.FeedbackEntity, error) {
	var entity domain.FeedbackEntity
	err := Repo.DB.Where("id = ?", feedbackId).Find(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (Repo *Repository) GetUnreadFeedback() (*[]domain.FeedbackEntity, error) {
	var entity []domain.FeedbackEntity
	err := Repo.DB.Preload("Category").Select("ID", "Name", "CreatedAt", "Title", "UUID").Where("have_read = ?", false).Find(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
