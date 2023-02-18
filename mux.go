package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/iwashi623/go_posts/clock"
	"github.com/iwashi623/go_posts/config"
	"github.com/iwashi623/go_posts/handler"
	"github.com/iwashi623/go_posts/store"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})

	v := validator.New()
	fmt.Println(cfg)
	db, cleanup, err := store.NewDB(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}
	r := store.Repository{Clocker: clock.RealClocker{}}

	cp := &handler.CreatePost{DB: db, Repo: &r, Validator: v}
	mux.Post("/posts", cp.ServeHTTP)

	lp := &handler.ListPost{DB: db, Repo: &r}
	mux.Get("/posts", lp.ServeHTTP)

	return mux, cleanup, nil
}
