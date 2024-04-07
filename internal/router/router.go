package router

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func New_Router(controllers *controllers.Controllers) http.Handler {
	r := chi.NewRouter()
	r.Post("/api/register", controllers.Auth.Register)
	r.Post("/api/login", controllers.Auth.Login)
	r.Route("/api/user", func(r chi.Router) {
		r.Post("/rent_book{id}", controllers.Auth.RentBook)
		r.Post("/return_book{id}", controllers.Auth.ReturnBook)
	})
	r.Route("/api/library", func(r chi.Router) {
		r.Get("/users", controllers.Library.GetUsers)
		r.Get("/authors", controllers.Library.GetAuthors)
		r.Post("/authors", controllers.Library.AddAuthor)
		r.Get("/books", controllers.Library.GetBooks)
		r.Post("/books", controllers.Library.AddBooks)
	})
	return r
}
