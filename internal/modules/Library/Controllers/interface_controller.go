package Controllers

import (
	"github.com/vadim-shalnev/GOlibrary/internal/infrastructures/responder"
	"github.com/vadim-shalnev/GOlibrary/internal/modules/User/Service"
	"net/http"
)

type LibraryController struct {
	Service   Service.Servicer
	Responder responder.Responder
}

type Librarer interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetAuthors(w http.ResponseWriter, r *http.Request)
	AddAuthor(w http.ResponseWriter, r *http.Request)
	GetBooks(w http.ResponseWriter, r *http.Request)
	AddBooks(w http.ResponseWriter, r *http.Request)
}

func NewLibraryController(service Service.Servicer, responder responder.Responder) *LibraryController {
	return &LibraryController{
		Service:   service,
		Responder: responder,
	}
}
