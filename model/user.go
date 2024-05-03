package model

import "net/http"

type User struct {
	Id           uint64 `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type UserList struct {
	Users []User `json:"users"`
}

func (user *User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (user *UserList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
