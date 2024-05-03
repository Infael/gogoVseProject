package mail

import (
	"fmt"
	"os"

	"github.com/Infael/gogoVseProject/model"
	resend "github.com/resend/resend-go/v2"
)

type ResendMailService struct {
	client *resend.Client
}

func NewResendMailService() *ResendMailService {
	return &ResendMailService{
		client: resend.NewClient(os.Getenv("RESEND_API_KEY")),
	}
}

func (rms *ResendMailService) SendMailToListOfUser(users []model.User, mailContent MailContent) error {
	emails := []*resend.SendEmailRequest{}

	for _, user := range users {
		emails = append(emails, &resend.SendEmailRequest{
			From:    "newsletter@resend.dev",
			To:      []string{user.Email},
			Subject: mailContent.Subject,
			Html:    mailContent.Html,
		})
	}

	_, err := rms.client.Batch.Send(emails)
	return err
}

func (rms *ResendMailService) SendMailLastNewsletterPost(newsletter model.Newsletter) error {
	lastNewsletterPost := newsletter.Posts[len(newsletter.Posts)-1]

	// TODO: add unsubscribe url
	html := fmt.Sprintf("<h2>%s</h2><p>%s</p><p><strong>Created:</strong>%s, Unsubscribe from newsletter here: %s</p>",
		lastNewsletterPost.Title,
		lastNewsletterPost.Body,
		lastNewsletterPost.CreatedAt,
		"www.dothisshitlater.com",
	)

	return rms.SendMailToListOfUser(newsletter.Subscribers, MailContent{
		Subject: newsletter.Title,
		Html:    html,
	})
}
