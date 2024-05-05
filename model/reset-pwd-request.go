package model

type ResetPwdRequest struct {
	NewPassword string `json:"new_password"  validate:"required"`
}
