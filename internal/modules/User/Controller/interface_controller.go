package Controller

import (
	"github.com/vadim-shalnev/GOlibrary/internal/infrastructures/responder"
	"github.com/vadim-shalnev/GOlibrary/internal/modules/User/Service"
	"net/http"
)

type AuthController struct {
	Servicer  Service.Servicer
	Responder responder.Responder
}
type Auther interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	RentBook(w http.ResponseWriter, r *http.Request)
	ReturnBook(w http.ResponseWriter, r *http.Request)
}

func NewAuthController(servicer Service.Servicer, responder responder.Responder) *AuthController {
	return &AuthController{Servicer: servicer, Responder: responder}
}
