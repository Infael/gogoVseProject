package mail

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"text/template"

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
		msg.SetHeader("From", os.Getenv("STMP_MAIL"))
		msg.SetHeader("To", email)
		msg.SetHeader("Subject", mailContent.Subject)
		if mailContent.UnsubscribeUrl != "" {
			msg.SetHeader("list-unsubscribe", mailContent.UnsubscribeUrl)
		}
		msg.SetBody("text/html", mailContent.Html)

		if err := sms.dialer.DialAndSend(msg); err != nil {
			return utils.InternalServerError(err)
		}
	}

	return nil
}

func (sms *StmpMailService) SendMailNewsletterPost(newsletter model.NewsletterAll, post model.PostAll, subscribers model.SubscriberAllList) error {
	for _, subscriber := range subscribers.Subscribers {
		urlHostname := os.Getenv("URL")
		unsubscribeUrl := fmt.Sprintf(urlHostname+"/newsletters/%s/subscribers/unsubscribe/%s", strconv.FormatUint(newsletter.Id, 10), strconv.FormatUint(subscriber.Id, 10))

		html, err := sms.parseTemplate("newsletter-template",
			map[string]string{
				"UnsubscribeLink": unsubscribeUrl,
				"Title":           post.Title,
				"Body":            post.Body,
			},
		)

		if err != nil {
			return err
		}

		sms.SendMailToListOfEmails([]string{subscriber.Email}, MailContent{
			Subject:        post.Title + " - " + newsletter.Title,
			Html:           html,
			UnsubscribeUrl: unsubscribeUrl,
		})
	}

	return nil
}

func (sms *StmpMailService) SendMailPasswordResetToken(user model.User, token string) error {
	// TODO: add correct url
	// TODO: html template
	url := os.Getenv("URL")

	html, err := sms.parseTemplate("password-reset",
		map[string]string{
			"ResetLink": url + "/password/reset/" + token,
		},
	)

	if err != nil {
		return err
	}

	return sms.SendMailToListOfEmails([]string{user.Email}, MailContent{
		Subject: "Password Reset Request",
		Html:    html,
	})
}

func (sms *StmpMailService) SendMailSubscriptionConfirmation(subscriber model.SubscriberAll, newsletter model.NewsletterAll, token string) error {
	urlHostname := os.Getenv("URL")
	url := fmt.Sprintf(urlHostname+"/newsletters/%s/subscribers/verify/%s", strconv.FormatUint(newsletter.Id, 10), token)

	html, err := sms.parseTemplate("newsletter-subscription-confirmation",
		map[string]string{
			"ConfirmationLink": url,
		},
	)

	if err != nil {
		return err
	}

	return sms.SendMailToListOfEmails([]string{subscriber.Email}, MailContent{
		Subject: "Welcome to Our Newsletter!",
		Html:    html,
	})
}

func (sms *StmpMailService) parseTemplate(templateName string, data map[string]string) (string, error) {
	rootDir, _ := os.Getwd()
	path := filepath.Join(filepath.Join(rootDir, "templates", templateName+".html"))
	t, err := template.ParseFiles(path)
	if err != nil {
		return "", err
	}

	var doc bytes.Buffer
	err = t.Execute(&doc, data)
	if err != nil {
		return "", nil
	}

	return doc.String(), err
}
