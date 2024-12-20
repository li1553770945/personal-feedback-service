package repo

import (
	"github.com/li1553770945/personal-feedback-service/biz/internal/domain"
	"gorm.io/gorm"
)

type IRepository interface {
	FindAllFeedbackCategories() (*[]domain.FeedbackCategoryEntity, error)
	SaveFeedback(entity *domain.FeedbackEntity) error
	FindFeedbackByUUID(uuid string) (*domain.FeedbackEntity, error)
	FindReplyByFeedbackID(messageId uint) (*domain.ReplyEntity, error)
	SaveReply(entity *domain.ReplyEntity) error
	FindFeedbackByID(messageId uint) (*domain.FeedbackEntity, error)
	GetUnreadFeedback() (*[]domain.FeedbackEntity, error)
}

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) IRepository {
	err := db.AutoMigrate(&domain.FeedbackCategoryEntity{})
	if err != nil {
		panic("迁移消息类别模型失败：" + err.Error())
	}
	err = db.AutoMigrate(&domain.FeedbackEntity{})
	if err != nil {
		panic("迁移消息模型失败：" + err.Error())
	}
	err = db.AutoMigrate(&domain.ReplyEntity{})
	if err != nil {
		panic("迁移回复模型失败：" + err.Error())
	}
	return &Repository{
		DB: db,
	}
}
