package controller

import (
	"fmt"
	"net/http"

	"github.com/Infael/gogoVseProject/service/password"
)

type PasswordController struct {
	passwordService *password.PasswordService
}

func (pc *PasswordController) ResetPassword(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Resetting user password...")
}

func (pc *PasswordController) SetNewPasswordWithResetToken(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Resetting user password...")
}
