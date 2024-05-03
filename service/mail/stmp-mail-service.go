package mail

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Infael/gogoVseProject/model"
	gomail "gopkg.in/gomail.v2"
)

type StmpMailService struct {
	dialer       *gomail.Dialer
	sendersEmail string
}

func NewStmpMailService() *StmpMailService {
	provider := os.Getenv("STMP_PROVIDER")
	port, err := strconv.Atoi(os.Getenv("STMP_PORT"))
	if err != nil {
		panic(err)
	}
	user := os.Getenv("STMP_MAIL")
	pwd := os.Getenv("STMP_PWD")

	return &StmpMailService{
		dialer:       gomail.NewDialer(provider, port, user, pwd),
		sendersEmail: user,
	}
}

func (sms *StmpMailService) SendMailToListOfUser(users []model.User, mailContent MailContent) error {
	for _, user := range users {
		msg := gomail.NewMessage()
		msg.SetHeader("From", sms.sendersEmail)
		msg.SetHeader("To", user.Email)
		msg.SetHeader("Subject", mailContent.Subject)
		msg.SetBody("text/html", mailContent.Html)

		if err := sms.dialer.DialAndSend(msg); err != nil {
			return err
		}
	}

	return nil
}

func (rms *StmpMailService) SendMailLastNewsletterPost(newsletter model.Newsletter) error {
	lastNewsletterPost := newsletter.Posts[len(newsletter.Posts)-1]

	// TODO: add unsubscribe url
	html := fmt.Sprintf(
		"<h2>%s</h2><p>%s</p><hr/><p><strong>Created:</strong> %s, Unsubscribe from newsletter here: <a href=\"%s\">%s</a></p>",
		lastNewsletterPost.Title,
		lastNewsletterPost.Body,
		lastNewsletterPost.CreatedAt,
		"www.dothisshitlater.com",
		"www.dothisshitlater.com",
	)

	return rms.SendMailToListOfUser(newsletter.Subscribers, MailContent{
		Subject: newsletter.Title,
		Html:    html,
	})
}
