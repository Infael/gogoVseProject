package service

import (
	"github.com/Infael/gogoVseProject/repository"
	"github.com/Infael/gogoVseProject/service/auth"
	"github.com/Infael/gogoVseProject/service/mail"
	"github.com/Infael/gogoVseProject/service/password"
	"github.com/redis/go-redis/v9"
	"gopkg.in/gomail.v2"
)

type Services struct {
	AuthService     auth.AuthService
	PasswordService password.PasswordService
	MailService     mail.MailService
}

func NewServices(repositories *repository.Repositories, redisClient *redis.Client, dialer *gomail.Dialer) *Services {
	mailService := mail.NewStmpMailService(dialer)

	return &Services{
		AuthService: auth.NewJwtAuthService(repositories.UserRepository),
		PasswordService: *password.NewPasswordService(
			mailService,
			redisClient,
			repositories.UserRepository,
		),
	}
}
