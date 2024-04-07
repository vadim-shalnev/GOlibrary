package Controllers

import (
	"encoding/json"
	"github.com/vadim-shalnev/GOlibrary/internal/entity"
	"net/http"
)

func (c *LibraryController) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.Service.GetUsers()
	if err != nil {
		c.Responder.HandleError(w, err)
	}
	c.Responder.SendJSONResponse(w, users)
}
func (c *LibraryController) GetAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := c.Service.GetAuthors()
	if err != nil {
		c.Responder.HandleError(w, err)
	}
	c.Responder.SendJSONResponse(w, authors)
}
func (c *LibraryController) AddAuthor(w http.ResponseWriter, r *http.Request) {
	var autor entity.Author
	err := json.NewDecoder(r.Body).Decode(&autor)
	if err != nil {
		c.Responder.HandleError(w, err)
	}
	err = c.Service.AddAuthor(autor)
	if err != nil {
		c.Responder.HandleError(w, err)
	}
	c.Responder.SendJSONResponse(w, "succes")
}
func (c *LibraryController) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := c.Service.GetBooks()
	if err != nil {
		c.Responder.HandleError(w, err)
	}
	c.Responder.SendJSONResponse(w, books)
}
func (c *LibraryController) AddBook(w http.ResponseWriter, r *http.Request) {
	var book entity.Books
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		c.Responder.HandleError(w, err)
	}
	err = c.Service.AddBook(book)
	if err != nil {
		c.Responder.HandleError(w, err)
	}
	c.Responder.SendJSONResponse(w, "succes")
}
