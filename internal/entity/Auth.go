package entity

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
