package post

import (
	"github.com/Infael/gogoVseProject/model"
	"github.com/Infael/gogoVseProject/repository"
	"github.com/Infael/gogoVseProject/service/mail"
)

type PostService struct {
	mailService          mail.MailService
	postRepository       repository.PostRepository
	newsletterRepository repository.NewsletterRepository
	subscriberRepository repository.SubscriberRepository
}

func NewPostService(mailService mail.MailService, postRepository *repository.PostRepository, newsletterRepository *repository.NewsletterRepository, subscriberRepository *repository.SubscriberRepository) *PostService {
	return &PostService{
		mailService:          mailService,
		postRepository:       *postRepository,
		newsletterRepository: *newsletterRepository,
		subscriberRepository: *subscriberRepository,
	}
}

func (p *PostService) CreatePost(post model.PostAll) (*model.PostAll, error) {

	newsletter, err := p.newsletterRepository.GetNewsletterById(post.NewsletterId)
	if err != nil {
		return nil, err
	}

	subscribers, err := p.subscriberRepository.GetAllSubscribersOfNewsletters(newsletter.Id)
	if err != nil {
		return nil, err
	}

	p.mailService.SendMailNewsletterPost(newsletter, post, model.SubscriberAllList{
		Subscribers: subscribers,
	})

	newPost, err := p.postRepository.CreatePost(&post)
	if err != nil {
		return nil, err
	}

	return &newPost, nil
}

func (p *PostService) GetAllPosts(newsletterId uint64) (*model.PostAllList, error) {
	if posts, err := p.postRepository.GetAllPostsOfNewsletters(newsletterId); err != nil {
		return nil, err
	} else {
		allPosts := model.PostAllList{
			Posts: posts,
		}
		return &allPosts, nil
	}
}
