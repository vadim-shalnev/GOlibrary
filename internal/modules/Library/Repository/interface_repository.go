package Repository

import (
	"database/sql"
	"github.com/vadim-shalnev/GOlibrary/internal/entity"
)

type LibraryRepository struct {
	DB *sql.DB
}

type Repository interface {
	GetUsers() ([]entity.UserBooks, error)
	GetAuthors() ([]entity.Author, error)
	GetBooks() ([]entity.Books, error)
	AddAuthor(author entity.Author) error
	AddBook(book entity.Books) error
}

func NewLibraryRepository(db *sql.DB) *LibraryRepository {
	return &LibraryRepository{
		DB: db,
	}
}
