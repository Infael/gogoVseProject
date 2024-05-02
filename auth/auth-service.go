package auth

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"

	// TODO: proper import
	models "github.com/Infael/gogoVseProject/model"
)


type AuthService struct {}

func (s *AuthService) Login(email, password string) (string, error) {
	// Get the user from the database
	// TODO: call DB to get user !!!
	var hashedPassword, _ = bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)
	var user = models.User{
		Id:  123,
		PasswordHash: string(hashedPassword),
		Email: "test",
	 }
   
	// Check if the provided password matches the user's password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
	 return "", errors.New("Invalid password or email.")
	}
   
	// Generate a Firebase custom token for the user
	token, err := CreateToken(user.Email)
	if err != nil {
	 log.Printf("failed to generate custom token: %v", err)
	 return "", errors.New("internal server error")
	}
   
	return token, nil
}

// Register creates a new user with the provided credentials and returns token
func (s *AuthService) Register(email, password string) (error) {
	// Check if the user with the email already exists
	// TODO: call DB to get user !!!
	var user = models.User{
		Id:  0,
		PasswordHash:   "",
		Email: "",
	}
	
   
	// TODO: fix condition when DB will be added
	if user.Id != 0 {
	 return errors.New("user with email already exists")
	}
   
	// Generate a hash of the user's password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
	 log.Printf("failed to hash password: %v", err)
	 return errors.New("internal server error")
	}
	log.Printf("password hashed: %v", hashedPassword)
   
	// Create a new user in the database
	// TODO: create user call DB to create user !!!
	var newUser = models.User{
		Id: uint64(10),
		PasswordHash: string(hashedPassword),
		Email: email,
	}
	log.Printf("new user created in DB: %d, %s, %s", newUser.Id, newUser.Email, newUser.PasswordHash)

	return nil
}
