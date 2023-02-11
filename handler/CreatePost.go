package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/iwashi623/go_posts/entity"
	"github.com/iwashi623/go_posts/store"
)

type CreatePost struct {
	PostStore *store.PostStore
	Validator *validator.Validate
}

func (cp *CreatePost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var b struct {
		Title string `json:"title" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
	}

	err := cp.Validator.Struct(b)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
	}

	p := &entity.Post{
		Title:     b.Title,
		Status:    entity.PostStatusDraft,
		CreatedAt: time.Now(),
	}
	id, err := cp.PostStore.Create(p)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
	}

	res := struct {
		ID entity.PostID `json:"id"`
	}{
		ID: id,
	}
	RespondJSON(ctx, w, res, http.StatusOK)
}
