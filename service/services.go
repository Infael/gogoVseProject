package service

import (
	"github.com/Infael/gogoVseProject/repository"
	"github.com/Infael/gogoVseProject/service/auth"
	"github.com/Infael/gogoVseProject/service/mail"
	"github.com/Infael/gogoVseProject/service/password"
	"github.com/Infael/gogoVseProject/service/user"
	"github.com/patrickmn/go-cache"
	"gopkg.in/gomail.v2"
)

type Services struct {
	AuthService     auth.AuthService
	PasswordService password.PasswordService
	MailService     mail.MailService
	UserService     user.UserService
}

func NewServices(repositories *repository.Repositories, cache *cache.Cache, dialer *gomail.Dialer) *Services {
	mailService := mail.NewStmpMailService(dialer)

	return &Services{
		AuthService: auth.NewJwtAuthService(repositories.UserRepository),
		PasswordService: *password.NewPasswordService(
			mailService,
			cache,
			repositories.UserRepository,
		),
		UserService: *user.NewUserService(repositories.UserRepository),
	}
}
