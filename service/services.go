package service

import (
	"github.com/Infael/gogoVseProject/repository"
	"github.com/Infael/gogoVseProject/service/auth"
)

type Services struct {
	AuthService auth.AuthService
}

func NewServices(repositories *repository.Repositories) *Services {
	return &Services{
		AuthService: auth.NewJwtAuthService(repositories.UserRepository),
	}
}
