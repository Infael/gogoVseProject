package controller

import (
	"fmt"
	"net/http"
)

type User struct{}

func (n *User) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting user account...")
}

func (n *User) ResetPassword(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Resetting user password...")
}

func (n *User) GetNewsletters(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting user newsletters...")
}
