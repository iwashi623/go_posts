package handler

import (
	"net/http"

	"github.com/iwashi623/go_posts/entity"
	"github.com/iwashi623/go_posts/store"
	"github.com/jmoiron/sqlx"
)

type ListPost struct {
	DB   *sqlx.DB
	Repo *store.Repository
}

type post struct {
	ID     entity.PostID     `json:"id"`
	Title  string            `json:"title"`
	Body   string            `json:"body"`
	Status entity.PostStatus `json:"status"`
}

func (lp *ListPost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	posts, err := lp.Repo.ListPosts(ctx, lp.DB)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	res := []post{}
	for _, p := range posts {
		res = append(res, post{
			ID:     p.ID,
			Title:  p.Title,
			Body:   p.Body,
			Status: p.Status,
		})
	}
	RespondJSON(ctx, w, res, http.StatusOK)
}
