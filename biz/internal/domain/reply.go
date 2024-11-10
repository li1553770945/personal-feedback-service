package domain

import "github.com/li1553770945/personal-feedback-service/biz/internal/do"

type ReplyEntity struct {
	do.BaseModel
	Content   string `json:"content"`
	Message   MessageEntity
	MessageID uint `json:"message_id"`
	HaveRead  bool `json:"have_read"`
}
