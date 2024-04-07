package Repository

import (
	"database/sql"
	"github.com/vadim-shalnev/GOlibrary/internal/entity"
)

type AuthRepository struct {
	DB *sql.DB
}

type Repositorier interface {
	RegisterUser(regData entity.User) (string, error)
	CheckUserInfo(loginData entity.LoginData) (string, error)
	RentBook(book entity.Books, login string) (entity.Books, error)
	ReturnBook(book entity.Books, login string) (string, error)
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		DB: db,
	}
}
