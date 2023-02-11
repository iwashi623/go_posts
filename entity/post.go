package entity

import "time"

type PostID int64
type PostStatus string

const (
	PostStatusDraft       PostStatus = "draft"
	PostStatusPublished   PostStatus = "published"
	PostStatusUnPublished PostStatus = "unpublished"
)

type Post struct {
	ID        PostID     `json:"id"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	Status    PostStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
}

type Posts []*Post
