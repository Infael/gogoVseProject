package mail

import (
	"fmt"

	"github.com/Infael/gogoVseProject/model"
	"github.com/Infael/gogoVseProject/utils"
	gomail "gopkg.in/gomail.v2"
)

type StmpMailService struct {
	dialer *gomail.Dialer
}

func NewStmpMailService(dialer *gomail.Dialer) *StmpMailService {
	return &StmpMailService{
		dialer: dialer,
	}
}

func (sms *StmpMailService) SendMailToListOfEmails(emails []string, mailContent MailContent) error {
	// TODO: should store errors and return them
	for _, email := range emails {
		msg := gomail.NewMessage()
		msg.SetHeader("From", sms.dialer.Username)
		msg.SetHeader("To", email)
		msg.SetHeader("Subject", mailContent.Subject)
		msg.SetBody("text/html", mailContent.Html)

		if err := sms.dialer.DialAndSend(msg); err != nil {
			return utils.InternalServerError(err)
		}
	}

	return nil
}

func (sms *StmpMailService) SendMailNewsletterPost(newsletter model.NewsletterAll, post model.PostAll, subscribers model.SubscriberAllList) error {
	// TODO: add unsubscribe url
	html := fmt.Sprintf(
		"<h2>%s</h2><p>%s</p><hr/><p><strong>Created:</strong> %s, Unsubscribe from newsletter here: <a href=\"%s\">Unsubscribe</a></p>",
		post.Title,
		post.Body,
		post.CreatedAt,
		"www.dothisshitlater.com",
	)

	return sms.SendMailToListOfEmails(sms.getListOfEmailsFromSubsctibers(subscribers), MailContent{
		Subject: newsletter.Title,
		Html:    html,
	})
}

func (sms *StmpMailService) SendMailPasswordResetToken(user model.User, token string) error {
	// TODO: add correct url
	// TODO: html template
	html := fmt.Sprintf(
		"<h2>Password Reset Request</h2><p>We received a request to reset your password. If you did not request this, you can ignore this email.</p><hr/><p><strong>To reset your password, please click the link below:</strong> <br/><br/><a href=\"%s\">Reset Password</a></p>",
		"http://localhost:3000/password/reset/"+token,
	)

	return sms.SendMailToListOfEmails([]string{user.Email}, MailContent{
		Subject: "Password Reset Request",
		Html:    html,
	})
}

func (sms *StmpMailService) getListOfEmailsFromSubsctibers(listOfSubscribers model.SubscriberAllList) []string {
	emails := []string{}
	for _, subscriber := range listOfSubscribers.Subscribers {
		emails = append(emails, subscriber.Email)
	}

	return emails
}
