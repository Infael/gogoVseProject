package auth

import (
	"errors"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	models "github.com/Infael/gogoVseProject/model"
	"github.com/Infael/gogoVseProject/repository"
	"github.com/Infael/gogoVseProject/utils"
)

type JwtAuthService struct {
	repository *repository.UserRepository
}

func NewJwtAuthService(repository *repository.UserRepository) *JwtAuthService {
	return &JwtAuthService{repository: repository}
}

func (s *JwtAuthService) Login(email, password string) (string, error) {
	// Get the user from the database
	user, err := s.repository.GetUserByEmail(email)

	if err != nil {
		return "", err
	}

	// Check if the provided password matches the user's password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", utils.NewError(errors.New("invalid credentials"), http.StatusUnauthorized)
	}

	// Generate a Firebase custom token for the user
	token, err := CreateToken(user.Email)
	if err != nil {
		log.Printf("failed to generate custom token: %v", err)
		return "", utils.InternalServerError(err)
	}

	return token, nil
}

// Register creates a new user with the provided credentials and returns token
func (s *JwtAuthService) Register(email, password string) error {
	// Generate a hash of the user's password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return utils.InternalServerError(err)
	}

	// Create a new user in the database
	newUser := models.User{
		PasswordHash: string(hashedPassword),
		Email:        email,
	}

	_, error := s.repository.CreateUser(&newUser)

	if error != nil {
		return error
	}

	log.Printf("new user created in DB: %d, %s, %s", newUser.Id, newUser.Email, newUser.PasswordHash)

	return nil
}
