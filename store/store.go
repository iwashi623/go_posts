package store

import (
	"errors"

	"github.com/iwashi623/go_posts/entity"
)

var (
	Posts       = &PostStore{Posts: map[entity.PostID]*entity.Post{}}
	ErrNotFound = errors.New("not found")
)

type PostStore struct {
	LastID entity.PostID
	Posts  map[entity.PostID]*entity.Post
}

func (ps *PostStore) Create(p *entity.Post) (entity.PostID, error) {
	ps.LastID++
	p.ID = ps.LastID
	ps.Posts[p.ID] = p
	return p.ID, nil
}

func (ps *PostStore) All(p *entity.Post) entity.Posts {
	posts := make([]*entity.Post, len(ps.Posts))
	for i, p := range ps.Posts {
		posts[i-1] = p
	}
	return posts
}
