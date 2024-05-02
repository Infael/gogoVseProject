package controller

import (
	"net/http"

	"github.com/Infael/gogoVseProject/auth"
	"github.com/Infael/gogoVseProject/controller/helpers"
	models "github.com/Infael/gogoVseProject/model"
)


type AuthController struct{
	authService *auth.AuthService
}


// NewAuthController creates a new instance of the AuthController struct
func NewAuthController() *AuthController {
	return &AuthController{authService: &auth.AuthService{}}
}

// Login handles the POST /login route and login a new user with the provided credentials
func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	// Get the email and password from the request body
	var loginRequest models.LoginRequest

	if err := helpers.GetObjectFromJson[models.LoginRequest](r, &loginRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Register the new user and get a custom token for the user
	customToken, err := c.authService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var loginResponse = models.LoginResponse{
		Token: customToken,
	}

	// Return the custom token to the client
    if err := helpers.SendResponseStatusOk(w, loginResponse); err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
        return
    }
}

// Register handles the POST /register route and creates a new user with the provided credentials
func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	// Get the email and password from the request body
	var registerRequest models.LoginRequest

	if err := helpers.GetObjectFromJson[models.LoginRequest](r, &registerRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


	// Register the new user and get a custom token for the user
	err := c.authService.Register(registerRequest.Email, registerRequest.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return the custom token to the client
    if err := helpers.SendResponse(w, nil, 204); err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
        return
    }
}
