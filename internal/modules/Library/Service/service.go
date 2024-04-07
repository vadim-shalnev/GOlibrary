package Service

import "github.com/vadim-shalnev/GOlibrary/internal/entity"

func (s *LibraryService) GetUsers() ([]entity.UserBooks, error) {
	return s.Repository.GetUsers()
}

func (s *LibraryService) GetAuthors() ([]entity.Author, error) {
	return s.Repository.GetAuthors()
}

func (s *LibraryService) GetBooks() ([]entity.Books, error) {
	return s.Repository.GetBooks()
}

func (s *LibraryService) AddAuthor(author entity.Author) error {
	return s.Repository.AddAuthor(author)
}

func (s *LibraryService) AddBook(book entity.Books) error {
	return s.Repository.AddBook(book)
}
