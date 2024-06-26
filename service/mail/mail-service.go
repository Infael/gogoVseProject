package mail

import "github.com/Infael/gogoVseProject/model"

type MailContent struct {
	Subject        string
	Html           string
	UnsubscribeUrl string
}

type MailService interface {
	SendMailToListOfEmails(emails []string, mailContent MailContent) error

	SendMailNewsletterPost(newsletter model.NewsletterAll, post model.PostAll, subscribers model.SubscriberAllList) error

	SendMailPasswordResetToken(user model.User, token string) error

	SendMailSubscriptionConfirmation(subscriber model.SubscriberAll, newsletter model.NewsletterAll, token string) error
}
