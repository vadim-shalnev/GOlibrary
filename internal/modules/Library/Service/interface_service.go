package Service

import (
	"github.com/vadim-shalnev/GOlibrary/internal/entity"
	"github.com/vadim-shalnev/GOlibrary/internal/modules/Library/Repository"
	"log"
	"os"
)

type LibraryService struct {
	Repository Repository.Repository
	log        *log.Logger
}

type LServicer interface {
	GetUsers() ([]entity.UserBooks, error)
	GetAuthors() ([]entity.Author, error)
	GetBooks() ([]entity.Books, error)
	AddAuthor(author entity.Author) error
	AddBook(book entity.Books) error
}

func NewLibraryService(repository Repository.Repository) *LibraryService {
	return &LibraryService{
		Repository: repository,
		log:        log.New(os.Stdout, "library", log.LstdFlags),
	}
}
