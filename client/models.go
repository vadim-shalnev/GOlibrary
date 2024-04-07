package main

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
type Books struct {
	Title  string `json:"book_name"`
	Author Author `json:"author"`
}
type Author struct {
	FullName string `json:"full_name"`
}

type UserBooks struct {
	Books []Books `json:"books"`
	Users User    `json:"user"`
}
