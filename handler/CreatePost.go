package handler

import "github.com/iwashi623/go_posts/store"

type CreatePost struct {
	PostStore *store.PostStore
	Validator *validator.Validate
}
