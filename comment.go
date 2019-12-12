package goshiki

import "time"

type Comment struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	CommentableID   int       `json:"commentable_id"`
	CommentableType string    `json:"commentable_type"`
	Body            string    `json:"body"`
	BodyHTML        string    `json:"html_body"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	IsOfftopic      bool      `json:"is_offtopic"`
	IsSummary       bool      `json:"is_summary"`
	CanBeEdited     bool      `json:"can_be_edited"`
	User            User      `json:"user"`
}
