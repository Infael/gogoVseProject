package service

import (
	"github.com/Infael/gogoVseProject/repository"
	"github.com/Infael/gogoVseProject/service/auth"
	"github.com/Infael/gogoVseProject/service/mail"
	"github.com/Infael/gogoVseProject/service/newsletter"
	"github.com/Infael/gogoVseProject/service/password"
	"github.com/Infael/gogoVseProject/service/post"
	"github.com/Infael/gogoVseProject/service/subscriber"
	"github.com/Infael/gogoVseProject/service/user"
	"github.com/patrickmn/go-cache"
	"gopkg.in/gomail.v2"
)

type Services struct {
	AuthService       auth.AuthService
	PasswordService   password.PasswordService
	MailService       mail.MailService
	UserService       user.UserService
	NewsletterService newsletter.NewsletterService
	PostService       post.PostService
	SubscriberService subscriber.SubscriberService
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
		UserService:       *user.NewUserService(repositories.UserRepository),
		NewsletterService: *newsletter.NewNewsletterService(repositories.NewsletterRepository),
		PostService:       *post.NewPostService(mailService, repositories.PostRepository, repositories.NewsletterRepository, repositories.SubscriberRepository),
		SubscriberService: *subscriber.NewSubscriberService(mailService, repositories.SubscriberRepository, repositories.NewsletterRepository),
	}
}
