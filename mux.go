package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/iwashi623/go_posts/handler"
	"github.com/iwashi623/go_posts/store"
)

func NewMux() http.Handler {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})

	v := validator.New()
	cp := &handler.CreatePost{Store: store.Posts, Validator: v}
	mux.Post("/posts", cp.ServeHTTP)
	lp := &handler.ListPost{Store: store.Posts}
	mux.Get("/posts", lp.ServeHTTP)
	return mux
}
