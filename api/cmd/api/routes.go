package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	v1Router := chi.NewRouter()

	v1Router.NotFound(app.notFound)
	v1Router.MethodNotAllowed(app.methodNotAllowed)

	v1Router.Use(middleware.Logger)
	v1Router.Use(middleware.Recoverer)
	v1Router.Use(middleware.CleanPath)
	v1Router.Use(middleware.StripSlashes)
	v1Router.Use(middleware.RealIP)
	v1Router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	v1Router.Get("/status", app.status)
	v1Router.Route("/api/v1", func(router chi.Router) {
		router.Post("/auth/token", app.authToken)
		router.Post("/users", app.createUser)
	})
	// mux.Post("/auth/token", app.authToken)
	// mux.Post("/users", app.createUser)

	return v1Router
}
