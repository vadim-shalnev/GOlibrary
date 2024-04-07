package entity

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
