package modules

import "net/http"

type Facade struct {
	c Controllers
}

type Facader interface {
	reg(w http.ResponseWriter, r *http.Request)
	log(w http.ResponseWriter, r *http.Request)
	Rent(w http.ResponseWriter, r *http.Request)
	Ret(w http.ResponseWriter, r *http.Request)
	Getu(w http.ResponseWriter, r *http.Request)
	Geta(w http.ResponseWriter, r *http.Request)
	Adda(w http.ResponseWriter, r *http.Request)
	Getb(w http.ResponseWriter, r *http.Request)
	Addb(w http.ResponseWriter, r *http.Request)
}

func (f Facade) reg(w http.ResponseWriter, r *http.Request) {
	f.c.User.Register(w, r)
}
func (f Facade) log(w http.ResponseWriter, r *http.Request) {
	f.c.User.Login(w, r)
}
func (f Facade) Rent(w http.ResponseWriter, r *http.Request) {
	f.c.User.RentBook(w, r)
}
func (f Facade) Ret(w http.ResponseWriter, r *http.Request) {
	f.c.User.ReturnBook(w, r)
}
func (f Facade) Getu(w http.ResponseWriter, r *http.Request) {
	f.c.Library.GetUsers(w, r)
}
func (f Facade) Geta(w http.ResponseWriter, r *http.Request) {
	f.c.Library.GetAuthors(w, r)
}
func (f Facade) Adda(w http.ResponseWriter, r *http.Request) {
	f.c.Library.AddAuthor(w, r)
}
func (f Facade) Getb(w http.ResponseWriter, r *http.Request) {
	f.c.Library.GetBooks(w, r)
}
func (f Facade) Addb(w http.ResponseWriter, r *http.Request) {
	f.c.Library.AddBook(w, r)
}
