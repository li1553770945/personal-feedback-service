package domain

import "github.com/li1553770945/personal-feedback-service/biz/internal/do"

type FeedbackCategoryEntity struct {
	do.BaseModel
	Name   string `json:"name"`
	CanUse bool   `json:"can_use"`
}
type FeedbackEntity struct {
	do.BaseModel
	Category   FeedbackCategoryEntity `json:"category"`
	CategoryID int                    `json:"category_id"`
	Title      string                 `json:"title"`
	Content    string                 `json:"content"`
	Name       string                 `json:"name"`
	Contact    string                 `json:"contact"`
	HaveRead   bool                   `json:"have_read"`
	UUID       string                 `json:"uuid"`
}
