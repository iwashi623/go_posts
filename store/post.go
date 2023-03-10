package store

import (
	"context"

	"github.com/iwashi623/go_posts/entity"
)

func (r *Repository) ListPosts(ctx context.Context, db Queryer) (entity.Posts, error) {
	posts := entity.Posts{}
	sql := `SELECT * FROM post`
	if err := db.SelectContext(ctx, &posts, sql); err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *Repository) CreatePost(ctx context.Context, db Execer, p *entity.Post) error {
	p.CreatedAt = r.Clocker.Now()
	p.UpdatedAt = r.Clocker.Now()
	sql := `INSERT INTO post
	        (title, body, status, created_at, updated_at) 
	        VALUES (?, ?, ?, ?, ?)`

	result, err := db.ExecContext(
		ctx, sql, p.Title, p.Body, p.Status, p.CreatedAt, p.UpdatedAt,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = entity.PostID(id)
	return nil
}
