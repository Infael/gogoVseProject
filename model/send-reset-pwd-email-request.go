package model

type SendResetPwdEmailRequest struct {
	Email string `json:"email"  validate:"required"`
}
