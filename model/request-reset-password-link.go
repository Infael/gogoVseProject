package model

type RequestResetPasswordLink struct {
	Email string `json:"email"  validate:"required"`
}
