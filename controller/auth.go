package controller

import (
	"net/http"

	"github.com/Infael/gogoVseProject/controller/helpers"
	models "github.com/Infael/gogoVseProject/model"
	"github.com/Infael/gogoVseProject/service/auth"
	"github.com/Infael/gogoVseProject/utils"
)

type AuthController struct {
	authService auth.AuthService
}

// NewAuthController creates a new instance of the AuthController struct
func NewAuthController(authService *auth.AuthService) *AuthController {
	return &AuthController{authService: *authService}
}

// Login handles the POST /login route and login a new user with the provided credentials
func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	// Get the email and password from the request body
	var loginRequest models.LoginRequest

	if err := helpers.GetObjectFromJson(r, &loginRequest); err != nil {
		helpers.SendError(w, r, err)
		return
	}

	// Register the new user and get a custom token for the user
	customToken, err := c.authService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		helpers.SendError(w, r, err)
		return
	}

	var loginResponse = models.LoginResponse{
		Token: customToken,
	}

	// Return the custom token to the client
	if err := helpers.SendResponseStatusOk(w, loginResponse); err != nil {
		helpers.SendError(w, r, err)
		return
	}
}

// Register handles the POST /register route and creates a new user with the provided credentials
func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	// Get the email and password from the request body
	var registerRequest models.LoginRequest

	if err := helpers.GetObjectFromJson(r, &registerRequest); err != nil {
		helpers.SendError(w, r, utils.ErrorBadRequest(err))
		return
	}

	// Register the new user and get a custom token for the user
	err := c.authService.Register(registerRequest.Email, registerRequest.Password)
	if err != nil {
		helpers.SendError(w, r, err)
		return
	}

	// Return the custom token to the client
	if err := helpers.SendResponse(w, nil, http.StatusNoContent); err != nil {
		helpers.SendError(w, r, err)
		return
	}
}
