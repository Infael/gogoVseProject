package model

type ResetPasswordAttempt struct {
	NewPassword string `json:"new_password"  validate:"required"`
}
