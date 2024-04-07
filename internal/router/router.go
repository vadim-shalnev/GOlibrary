package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/vadim-shalnev/GOlibrary/internal/modules"
	"net/http"
)

func New_Router(controllers *modules.Controllers) http.Handler {
	r := chi.NewRouter()
	r.Post("/api/register", controllers.User.Register)
	r.Post("/api/login", controllers.User.Login)
	r.Route("/api/user", func(r chi.Router) {
		r.Post("/rent_book{login}", controllers.User.RentBook)
		r.Post("/return_book{login}", controllers.User.ReturnBook)
	})
	r.Route("/api/library", func(r chi.Router) {
		r.Get("/users", controllers.Library.GetUsers)
		r.Get("/authors", controllers.Library.GetAuthors)
		r.Post("/authors", controllers.Library.AddAuthor)
		r.Get("/books", controllers.Library.GetBooks)
		r.Post("/books", controllers.Library.AddBook)
	})
	return r
}
