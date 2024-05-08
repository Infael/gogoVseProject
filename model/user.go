package model

type UserAll struct {
	Id           uint64 `json:"id" validate:"required"`
	Email        string `json:"email" validate:"required"`
	PasswordHash string `json:"password_hash" validate:"required"`
}

type User struct {
	Id    uint64 `json:"id" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type UserUpdate struct {
	Email string `json:"email" validate:"required"`
}

type UserList struct {
	Users []User `json:"users" validate:"required"`
}
