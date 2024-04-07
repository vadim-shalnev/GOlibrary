package Controllers

import (
	"github.com/vadim-shalnev/GOlibrary/internal/infrastructures/responder"
	"github.com/vadim-shalnev/GOlibrary/internal/modules/Library/Service"
	"net/http"
)

type LibraryController struct {
	Service   Service.LServicer
	Responder responder.Responder
}

type Librarer interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetAuthors(w http.ResponseWriter, r *http.Request)
	AddAuthor(w http.ResponseWriter, r *http.Request)
	GetBooks(w http.ResponseWriter, r *http.Request)
	AddBook(w http.ResponseWriter, r *http.Request)
}

func NewLibraryController(service Service.LServicer, responder responder.Responder) *LibraryController {
	return &LibraryController{
		Service:   service,
		Responder: responder,
	}
}
