package modules

import (
	lService "github.com/vadim-shalnev/GOlibrary/internal/modules/Library/Service"
	uService "github.com/vadim-shalnev/GOlibrary/internal/modules/User/Service"
)

type Services struct {
	User    uService.Servicer
	Library lService.LServicer
}

func NewServices(repositories Repositoryes) Services {
	return Services{
		User:    uService.NewAuthController(repositories.User),
		Library: lService.NewLibraryService(repositories.Library),
	}
}
