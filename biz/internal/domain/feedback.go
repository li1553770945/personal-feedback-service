package domain

import "github.com/li1553770945/personal-feedback-service/biz/internal/do"

type MessageCategoryEntity struct {
	do.BaseModel
	Name   string `json:"name"`
	Value  string `json:"value"`
	CanUse bool   `json:"can_use"`
}
type MessageEntity struct {
	do.BaseModel
	Category   MessageCategoryEntity `json:"category"`
	CategoryID int                   `json:"category_id"`
	Title      string                `json:"title"`
	Message    string                `json:"message"`
	Name       string                `json:"name"`
	Contact    string                `json:"contact"`
	HaveRead   bool                  `json:"have_read"`
	UUID       string                `json:"uuid"`
}
