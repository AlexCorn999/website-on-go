package main

import (
	"net/http"

	"github.com/AlexCorn999/website-on-go/pkg/config"
	"github.com/AlexCorn999/website-on-go/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
