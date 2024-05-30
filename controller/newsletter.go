package controller

import (
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/Infael/gogoVseProject/controller/helpers"
	"github.com/Infael/gogoVseProject/model"
	"github.com/Infael/gogoVseProject/service/newsletter"
	"github.com/Infael/gogoVseProject/service/post"
	"github.com/Infael/gogoVseProject/service/subscriber"
	"github.com/Infael/gogoVseProject/service/user"
	"github.com/Infael/gogoVseProject/utils"
	"github.com/go-chi/chi/v5"
)

type NewsletterController struct {
	newsletterService *newsletter.NewsletterService
	postService       *post.PostService
	userService       *user.UserService
	subscriberService *subscriber.SubscriberService
}

func NewNewsletterController(newsletterService *newsletter.NewsletterService, postService *post.PostService, userService *user.UserService, subscriberService *subscriber.SubscriberService) *NewsletterController {
	return &NewsletterController{
		newsletterService: newsletterService,
		userService:       userService,
		postService:       postService,
		subscriberService: subscriberService,
	}
}

func (n *NewsletterController) getLoggedUser(r *http.Request) (*model.UserAll, error) {
	userMail := r.Context().Value("email")
	if userMail == nil {
		return nil, utils.ErrorBadRequest(errors.New("user not found in context"))
	}
	user, err := n.userService.GetUserByEmail(userMail.(string))
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (n *NewsletterController) Create(w http.ResponseWriter, r *http.Request) {
	var newNewsletter model.NewsletterCreate

	if err := helpers.GetObjectFromJson(r, &newNewsletter); err != nil {
		helpers.SendError(w, r, err)
		return
	}

	user, err := n.getLoggedUser(r)
	if err != nil {
		helpers.SendError(w, r, err)
		return
	}
	newNewsletter.Creator = user.Id

	if createdNewsletter, err := n.newsletterService.CreateNewsletter(newNewsletter); err != nil {
		helpers.SendError(w, r, err)
		return
	} else {
		helpers.SendResponseStatusOk(w, createdNewsletter)
	}
}

func (n *NewsletterController) List(w http.ResponseWriter, r *http.Request) {
	if newsletters, err := n.newsletterService.GetAllNewsletters(); err != nil {
		helpers.SendError(w, r, err)
		return
	} else {
		helpers.SendResponseStatusOk(w, newsletters)
	}
}

func (n *NewsletterController) GetById(w http.ResponseWriter, r *http.Request) {
	if id, err := helpers.GetIdFromRequest(r); err != nil {
		helpers.SendError(w, r, err)
		return
	} else {
		if newsletter, err := n.newsletterService.GetNewsletterById(*id); err != nil {
			helpers.SendError(w, r, err)
			return
		} else {
			helpers.SendResponseStatusOk(w, newsletter)
		}
	}

}

func (n *NewsletterController) UpdateById(w http.ResponseWriter, r *http.Request) {
	if id, err := helpers.GetIdFromRequest(r); err != nil {
		helpers.SendError(w, r, err)
		return
	} else {
		var updatedNewsletter model.NewsletterUpdate

		if err := helpers.GetObjectFromJson(r, &updatedNewsletter); err != nil {
			helpers.SendError(w, r, err)
			return
		}

		// Check if the user is allowed to update the newsletter
		user, err := n.getLoggedUser(r)
		if err != nil {
			helpers.SendError(w, r, err)
			return
		}
		newsletter, err := n.newsletterService.GetNewsletterById(*id)
		if err != nil {
			helpers.SendError(w, r, err)
			return
		}
		if newsletter.Creator != user.Id {
			helpers.SendError(w, r, utils.ErrorForbidden(errors.New("user is not allowed to update this newsletter")))
			return
		}

		// Update the newsletter
		if updatedNewsletter, err := n.newsletterService.UpdateNewsletter(*id, &updatedNewsletter); err != nil {
			helpers.SendError(w, r, err)
			return
		} else {
			helpers.SendResponseStatusOk(w, updatedNewsletter)
		}
	}
}

func (n *NewsletterController) DeleteById(w http.ResponseWriter, r *http.Request) {
	if id, err := helpers.GetIdFromRequest(r); err != nil {
		helpers.SendError(w, r, err)
		return
	} else {

		// Check if the user is allowed to delete the newsletter
		user, err := n.getLoggedUser(r)
		if err != nil {
			helpers.SendError(w, r, err)
			return
		}
		newsletter, err := n.newsletterService.GetNewsletterById(*id)
		if err != nil {
			helpers.SendError(w, r, err)
			return
		}
		if newsletter.Creator != user.Id {
			helpers.SendError(w, r, utils.ErrorForbidden(errors.New("user is not allowed to update this newsletter")))
			return
		}

		// Delete the newsletter
		if err := n.newsletterService.DeleteNewsletter(*id); err != nil {
			helpers.SendError(w, r, err)
			return
		} else {
			helpers.SendResponse(w, nil, http.StatusNoContent)
		}
	}
}

func (n *NewsletterController) CreatePost(w http.ResponseWriter, r *http.Request) {
	var newPostUpdate model.PostUpdate

	if err := helpers.GetObjectFromJson(r, &newPostUpdate); err != nil {
		helpers.SendError(w, r, err)
		return
	}

	id, err := helpers.GetIdFromRequest(r)
	if err != nil {
		helpers.SendError(w, r, err)
		return
	}

	newsletter, err := n.newsletterService.GetNewsletterById(*id)
	if err != nil {
		helpers.SendError(w, r, err)
		return
	}

	user, err := n.getLoggedUser(r)
	if err != nil {
		helpers.SendError(w, r, err)
		return
	}

	// Check if user is allowed to post to this newsletter
	if newsletter.Creator != user.Id {
		helpers.SendError(w, r, utils.ErrorForbidden(errors.New("user is not allowed to post to this newsletter")))
		return
	}

	newPost := model.PostAll{
		Title:        newPostUpdate.Title,
		Body:         newPostUpdate.Body,
		NewsletterId: *id,
	}

	if createdPost, err := n.postService.CreatePost(newPost); err != nil {
		helpers.SendError(w, r, err)
		return
	} else {
		helpers.SendResponseStatusOk(w, createdPost)
	}
}

func (n *NewsletterController) GetPosts(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetIdFromRequest(r)
	if err != nil {
		helpers.SendError(w, r, err)
		return
	}

	newsletter, err := n.newsletterService.GetNewsletterById(*id)
	if err != nil {
		helpers.SendError(w, r, err)
		return
	}

	user, err := n.getLoggedUser(r)
	if err != nil {
		helpers.SendError(w, r, err)
		return
	}

	// Check if user is allowed to post to this newsletter
	if newsletter.Creator != user.Id {
		helpers.SendError(w, r, utils.ErrorForbidden(errors.New("user is not allowed to post to this newsletter")))
		return
	}

	if posts, err := n.postService.GetAllPosts(*id); err != nil {
		helpers.SendError(w, r, err)
		return
	} else {
		helpers.SendResponseStatusOk(w, posts)
	}
}

func (n *NewsletterController) Subscribe(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetIdFromRequest(r)
	if err != nil {
		helpers.SendError(w, r, err)
		return
	}

	var subscribePayload model.Subscribe

	if err := helpers.GetObjectFromJson(r, &subscribePayload); err != nil {
		helpers.SendError(w, r, err)
		return
	}

	if _, err := n.subscriberService.SubscribeToNewsletter(subscribePayload.Email, *id); err != nil {
		helpers.SendError(w, r, err)
		return
	} else {
		helpers.SendResponseStatusOk(w, nil)
	}

}

func (n *NewsletterController) GetSubscribers(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetIdFromRequest(r)
	if err != nil {
		helpers.SendError(w, r, err)
		return
	}

	newsletter, err := n.newsletterService.GetNewsletterById(*id)
	if err != nil {
		helpers.SendError(w, r, err)
		return
	}

	user, err := n.getLoggedUser(r)
	if err != nil {
		helpers.SendError(w, r, err)
		return
	}

	// Check if user is allowed to post to this newsletter
	if newsletter.Creator != user.Id {
		helpers.SendError(w, r, utils.ErrorForbidden(errors.New("user is not allowed to post to this newsletter")))
		return
	}

	if posts, err := n.subscriberService.GetAllSubscribers(*id); err != nil {
		helpers.SendError(w, r, err)
		return
	} else {
		helpers.SendResponseStatusOk(w, posts)
	}
}

func (n *NewsletterController) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	newsletterId, err := helpers.GetIdFromRequest(r)
	if err != nil {
		helpers.SendError(w, r, err)
		return
	}

	subIdRaw := chi.URLParam(r, "subId")
	if subIdRaw == "" {
		helpers.SendError(w, r, utils.ErrorBadRequest(errors.New("subId is required")))
		return
	}
	subId, err := strconv.ParseUint(subIdRaw, 10, 64)
	if err != nil {
		helpers.SendError(w, r, utils.ErrorBadRequest(errors.New("subId must be a number")))
		return
	}

	if err := n.subscriberService.Unsubscribe(*newsletterId, subId); err != nil {
		helpers.SendError(w, r, err)
		return
	} else {
		if r.Method == http.MethodGet {
			url := os.Getenv("URL")
			helpers.Redirect(w, r, url+"/static/newsletter-unsubscribe.html")
		} else {
			helpers.SendResponseStatusOk(w, nil)
		}
	}

	helpers.SendResponse(w, nil, http.StatusOK)

}

func (n *NewsletterController) VerifySubscriber(w http.ResponseWriter, r *http.Request) {
	newsletterId, err := helpers.GetIdFromRequest(r)
	if err != nil {
		helpers.SendError(w, r, err)
		return
	}

	token := chi.URLParam(r, "token")
	if token == "" {
		helpers.SendError(w, r, utils.ErrorBadRequest(errors.New("token is required")))
		return
	}

	if err := n.subscriberService.VerifySubscriber(*newsletterId, token); err != nil {
		helpers.SendError(w, r, err)
		return
	} else {
		if r.Method == http.MethodGet {
			url := os.Getenv("URL")
			helpers.Redirect(w, r, url+"/static/newsletter-confirmation.html")
		} else {
			helpers.SendResponseStatusOk(w, nil)
		}
	}

	helpers.SendResponse(w, nil, http.StatusOK)
}
