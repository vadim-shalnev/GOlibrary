package modules

import (
	"github.com/vadim-shalnev/GOlibrary/internal/infrastructures/responder"
	lController "github.com/vadim-shalnev/GOlibrary/internal/modules/Library/Controllers"
	uController "github.com/vadim-shalnev/GOlibrary/internal/modules/User/Controller"
)

type Controllers struct {
	User    uController.Auther
	Library lController.Librarer
}

func NewControllers(services Services, responder responder.Responder) *Controllers {
	return &Controllers{
		User:    uController.NewAuthController(services.User, responder),
		Library: lController.NewLibraryController(services.Library, responder),
	}
}
