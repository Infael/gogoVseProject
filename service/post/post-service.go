package post

import (
	"github.com/Infael/gogoVseProject/model"
	"github.com/Infael/gogoVseProject/repository"
	"github.com/Infael/gogoVseProject/service/mail"
)

type PostService struct {
	mailService    mail.MailService
	postRepository repository.PostRepository
}

func NewPostService(mailService mail.MailService, postRepository *repository.PostRepository) *PostService {
	return &PostService{
		mailService:    mailService,
		postRepository: *postRepository,
	}
}

func (p *PostService) CreatePost(post model.PostAll) (*model.PostAll, error) {
	newPost, err := p.postRepository.CreatePost(&post)
	if err != nil {
		return nil, err
	}

	return &newPost, nil
}

// func (p *PostService) GetAllPosts() (*model.PostAllList, error) {
// 	if newsletters, err := n.newsletterRepository.GetAllNewsletters(); err != nil {
// 		return nil, err
// 	} else {
// 		allNewsletters := model.NewsletterAllList{
// 			Newsletters: newsletters,
// 		}
// 		return &allNewsletters, nil
// 	}
// }
