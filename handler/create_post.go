package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/iwashi623/go_posts/entity"
	"github.com/iwashi623/go_posts/store"
	"github.com/jmoiron/sqlx"
)

type CreatePost struct {
	DB        *sqlx.DB
	Repo      *store.Repository
	Validator *validator.Validate
}

func (cp *CreatePost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var b struct {
		Title string `json:"title" validate:"required"`
		Body  string `json:"body" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	err := cp.Validator.Struct(b)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	p := &entity.Post{
		Title:  b.Title,
		Body:   b.Body,
		Status: entity.PostStatusDraft,
	}

	if err := cp.Repo.CreatePost(ctx, cp.DB, p); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	res := struct {
		ID entity.PostID `json:"id"`
	}{
		ID: p.ID,
	}
	RespondJSON(ctx, w, res, http.StatusOK)
}
