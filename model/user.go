package model

type User struct {
	Id           uint64 `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type UserList struct {
	Users []User `json:"users"`
}
