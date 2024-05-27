package newsletter

import (
	"github.com/Infael/gogoVseProject/model"
	"github.com/Infael/gogoVseProject/repository"
)

type NewsletterService struct {
	newsletterRepository *repository.NewsletterRepository
}

func NewNewsletterService(newsletterRepository *repository.NewsletterRepository) *NewsletterService {
	return &NewsletterService{
		newsletterRepository: newsletterRepository,
	}
}

func (n *NewsletterService) CreateNewsletter(newsletter model.NewsletterCreate) (*model.NewsletterAll, error) {

	if newNewsletter, err := n.newsletterRepository.CreateNewsletter(newsletter); err != nil {
		return nil, err
	} else {
		return &newNewsletter, nil
	}
}

func (n *NewsletterService) UpdateNewsletter(id uint64, newsletter *model.NewsletterUpdate) (*model.NewsletterAll, error) {
	oldNewsletter, err := n.newsletterRepository.GetNewsletterById(id)
	if err != nil {
		return nil, err
	}
	oldNewsletter.Title = newsletter.Title
	oldNewsletter.Description = newsletter.Description

	if updatedNewsletter, err := n.newsletterRepository.UpdateNewsletter(&oldNewsletter); err != nil {
		return nil, err
	} else {
		return &updatedNewsletter, nil
	}
}

func (n *NewsletterService) DeleteNewsletter(id uint64) error {
	return n.newsletterRepository.DeleteNewsletter(id)
}

func (n *NewsletterService) GetNewsletterById(id uint64) (*model.NewsletterAll, error) {
	if newsletter, err := n.newsletterRepository.GetNewsletterById(id); err != nil {
		return nil, err
	} else {
		return &newsletter, nil
	}
}

func (n *NewsletterService) GetAllNewsletters() (*model.NewsletterAllList, error) {
	if newsletters, err := n.newsletterRepository.GetAllNewsletters(); err != nil {
		return nil, err
	} else {
		allNewsletters := model.NewsletterAllList{
			Newsletters: newsletters,
		}
		return &allNewsletters, nil
	}
}
