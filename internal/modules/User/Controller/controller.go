package Controller

import (
	"encoding/json"
	"github.com/vadim-shalnev/GOlibrary/internal/entity"
	"net/http"
)

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var regData entity.User
	err := json.NewDecoder(r.Body).Decode(&regData)
	if err != nil {
		c.Responder.HandleError(w, err)
		return
	}
	success, err := c.Servicer.Register(regData)
	if err != nil {
		c.Responder.HandleError(w, err)
		return
	}
	c.Responder.SendJSONResponse(w, success)
}
func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var loginData entity.LoginData
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		c.Responder.HandleError(w, err)
		return
	}
	success, err := c.Servicer.Login(loginData)
	if err != nil {
		c.Responder.HandleError(w, err)
		return
	}
	c.Responder.SendJSONResponse(w, success)
}
func (c *AuthController) RentBook(w http.ResponseWriter, r *http.Request) {
	var book entity.Books
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		c.Responder.HandleError(w, err)
		return
	}
	login := chi.URLParam(r, "id")
	b, err := c.Servicer.RentBook(book, login)
	if err != nil {
		c.Responder.HandleError(w, err)
		return
	}
	c.Responder.SendJSONResponse(w, b)
}
func (c *AuthController) ReturnBook(w http.ResponseWriter, r *http.Request) {
	var book entity.Books
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		c.Responder.HandleError(w, err)
		return
	}
	login := chi.URLParam(r, "id")
	resp, err := c.Servicer.ReturnBook(book, login)
	if err != nil {
		c.Responder.HandleError(w, err)
		return
	}
	c.Responder.SendJSONResponse(w, resp)
}
