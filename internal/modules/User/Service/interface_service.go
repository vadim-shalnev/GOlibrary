package Service

import (
	"github.com/vadim-shalnev/GOlibrary/internal/entity"
	"github.com/vadim-shalnev/GOlibrary/internal/modules/User/Repository"
)

type AuthService struct {
	Repository Repository.Repositorier
}

type Servicer interface {
	Register(regData entity.User) (string, error)
	Login(regData entity.LoginData) (string, error)
	RentBook(book entity.Books, login string) (entity.Books, error)
	ReturnBook(book entity.Books, login string) (string, error)
}

func NewAuthController(repository Repository.Repositorier) *AuthService {
	return &AuthService{
		Repository: repository,
	}
}
