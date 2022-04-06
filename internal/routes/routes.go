package routes

import (
	"net/http"

	"bitbucket.org/janpavtel/site/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func CreateRoutes(view *handlers.View) http.Handler {
	mux := chi.NewMux()

	mux.Get("/signup", view.SignUp)
	mux.Get("/users", view.ShowUsers)
	mux.Get("/users/{id}", view.ShowUser)
	mux.Post("/users", view.NewUserCreation)

	fileServer := http.FileServer(http.Dir("./www/static/"))

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
