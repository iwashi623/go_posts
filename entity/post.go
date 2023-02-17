package entity

import "time"

type PostID int
type PostStatus string

const (
	PostStatusDraft     PostStatus = "draft"
	PostStatusPublished PostStatus = "published"
	PostStatusPrivate   PostStatus = "private"
)

type Post struct {
	ID        PostID     `json:"id" db:"id"`
	Title     string     `json:"title" db:"title"`
	Body      string     `json:"body" db:"body"`
	Status    PostStatus `json:"status" db:"status"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
}

type Posts []*Post
