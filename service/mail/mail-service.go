package mail

import "github.com/Infael/gogoVseProject/model"

type MailContent struct {
	Subject string
	Html    string
}

type MailService interface {
	SendMailToListOfUsers(users []model.User, mailContent MailContent) error

	SendMailLastNewsletterPost(newsletter model.Newsletter) error

	SendMailPasswordResetToken(user model.User, token string) error
}
