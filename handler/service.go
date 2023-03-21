package handler

import (
	"context"

	"github.com/iwashi623/go_posts/entity"
)

// go:generate go run github.com/matryer/moq -out moq_test.go . ListPostsService,CreatePostService
type ListPostsService interface {
	ListPosts(ctx context.Context) (entity.Posts, error)
}

type CreatePostService interface {
	CreatePost(ctx context.Context, title string, body string) (entity.Post, error)
}
