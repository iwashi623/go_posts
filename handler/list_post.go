package handler

import (
	"net/http"

	"github.com/iwashi623/go_posts/entity"
	"github.com/iwashi623/go_posts/store"
)

type ListPost struct {
	Store *store.PostStore
}

type post struct {
	ID     entity.PostID     `json:"id"`
	Title  string            `json:"title"`
	Status entity.PostStatus `json:"status"`
}

func (lp *ListPost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	posts := lp.Store.All()
	res := []post{}
	for _, p := range posts {
		res = append(res, post{
			ID:     p.ID,
			Title:  p.Title,
			Status: p.Status,
		})
	}
	RespondJSON(ctx, w, res, http.StatusOK)
}
