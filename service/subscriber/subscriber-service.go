package subscriber

import (
	"github.com/Infael/gogoVseProject/model"
	"github.com/Infael/gogoVseProject/repository"
	"github.com/Infael/gogoVseProject/service/mail"
)

type SubscriberService struct {
	mailService          mail.MailService
	subscriberRepository repository.SubscriberRepository
	newsletterRepository repository.NewsletterRepository
}

func NewSubscriberService(mailService mail.MailService, subscriberRepository *repository.SubscriberRepository, newsletterRepository *repository.NewsletterRepository) *SubscriberService {
	return &SubscriberService{
		mailService:          mailService,
		subscriberRepository: *subscriberRepository,
		newsletterRepository: *newsletterRepository,
	}
}

func (p *SubscriberService) SubscribeToNewsletter(email string, newsletterId uint64) (*model.SubscriberAll, error) {
	subscriber, err := p.subscriberRepository.CreateOrFindSubscriber(email)
	if err != nil {
		return nil, err
	}

	newsletterSubscriber, err := p.subscriberRepository.SubscribeToNewsletter(newsletterId, subscriber.Id)
	if err != nil {
		return nil, err
	}

	newsletter, err := p.newsletterRepository.GetNewsletterById(newsletterId)
	if err != nil {
		return nil, err
	}

	if !newsletterSubscriber.Verified {
		p.mailService.SendMailSubscriptionConfirmation(subscriber, newsletter, newsletterSubscriber.Token)
	}

	return &subscriber, nil
}

func (p *SubscriberService) Unsubscribe(newsletterId uint64, subscriberId uint64) error {
	err := p.subscriberRepository.UnsubscribeFromNewsletter(newsletterId, subscriberId)
	return err
}

func (p *SubscriberService) VerifySubscriber(newsletterId uint64, token string) error {
	err := p.subscriberRepository.VerifySubscriber(newsletterId, token)
	return err
}

func (p *SubscriberService) GetAllSubscribers(newsletterId uint64) (*model.SubscriberAllList, error) {
	if subscribers, err := p.subscriberRepository.GetAllSubscribersOfNewsletters(newsletterId); err != nil {
		return nil, err
	} else {
		allSubscribers := model.SubscriberAllList{
			Subscribers: subscribers,
		}
		return &allSubscribers, nil
	}
}
