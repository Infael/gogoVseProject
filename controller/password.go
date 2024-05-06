package controller

import (
	"net/http"

	"github.com/Infael/gogoVseProject/controller/helpers"
	model "github.com/Infael/gogoVseProject/model"
	"github.com/Infael/gogoVseProject/service/password"
	"github.com/Infael/gogoVseProject/utils"
)

type PasswordController struct {
	passwordService *password.PasswordService
}

func NewPasswordController(passwordService *password.PasswordService) *PasswordController {
	return &PasswordController{
		passwordService: passwordService,
	}
}

func (pc *PasswordController) SendResetPasswordEmail(w http.ResponseWriter, r *http.Request) {
	var sendResetPasswordEmailRequest model.SendResetPwdEmailRequest

	if err := helpers.GetObjectFromJson(r, &sendResetPasswordEmailRequest); err != nil {
		helpers.SendError(w, r, utils.ErrorBadRequest(err))
		return
	}

	if err := pc.passwordService.SendPasswordResetToken(sendResetPasswordEmailRequest.Email); err != nil && err.Error() != "user not found" {
		helpers.SendError(w, r, err)
		return
	}

	helpers.SendResponse(w, nil, http.StatusNoContent)
	return
}

func (pc *PasswordController) SetNewPasswordWithResetToken(w http.ResponseWriter, r *http.Request) {
	var sendResetPwpRequest model.ResetPwdRequest

	if err := helpers.GetObjectFromJson(r, &sendResetPwpRequest); err != nil {
		helpers.SendError(w, r, utils.ErrorBadRequest(err))
		return
	}

	resetToken := r.PathValue("token")
	if err := pc.passwordService.ResetPasswordWithToken(sendResetPwpRequest.NewPassword, resetToken); err != nil && err.Error() != "user not found" {
		helpers.SendError(w, r, err)
		return
	}

	helpers.SendResponse(w, nil, http.StatusNoContent)
	return
}
