package modules

import (
	"database/sql"
	librar "github.com/vadim-shalnev/GOlibrary/internal/modules/Library/Repository"
	userr "github.com/vadim-shalnev/GOlibrary/internal/modules/User/Repository"
)

type Repositoryes struct {
	User    userr.Repositorier
	Library librar.Repository
}

func NewRepositories(db *sql.DB) Repositoryes {
	return Repositoryes{
		User:    userr.NewAuthRepository(db),
		Library: librar.NewLibraryRepository(db),
	}
}
