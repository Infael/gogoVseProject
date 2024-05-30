package controller

import (
	"net/http"
	"time"

	"github.com/Infael/gogoVseProject/controller/helpers"
	"github.com/Infael/gogoVseProject/model"
	"github.com/Infael/gogoVseProject/service/user"
)

type UserController struct {
	userService *user.UserService
}

func NewUserController(userService *user.UserService) *UserController {
	return &UserController{userService: userService}
}

func (n *UserController) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	email := ctx.Value("email").(string)
	oldUser, err := n.userService.GetUserByEmail(email)
	if err != nil {
		helpers.SendError(w, r, err)
		return
	}

	var userUpdateData model.UserUpdate

	err = helpers.GetObjectFromJson(r, &userUpdateData)
	if err != nil {
		helpers.SendError(w, r, err)
		return
	}

	oldUser.Email = userUpdateData.Email

	_, err = n.userService.UpdateUser(oldUser.Id, &oldUser)
	if err != nil {
		helpers.SendError(w, r, err)
		return
	}

	helpers.SendResponse(w, nil, http.StatusNoContent)
}

func (userController *UserController) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	email := ctx.Value("email").(string)
	if err := userController.userService.ScheduleUserDeletion(email, 60*time.Second); err != nil {
		helpers.SendError(w, r, err)
		return
	}
	helpers.SendResponse(w, nil, http.StatusNoContent)
}

func (n *UserController) CancelUserDeletion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	email := ctx.Value("email").(string)
	if err := n.userService.CancelUserDeletion(email); err != nil {
		helpers.SendError(w, r, err)
		return
	}
	helpers.SendResponse(w, nil, http.StatusNoContent)
}
