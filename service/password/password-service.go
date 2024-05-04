package password

import (
	"context"
	"errors"
	"time"

	"github.com/Infael/gogoVseProject/model"
	"github.com/Infael/gogoVseProject/service/mail"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct {
	mailService  mail.MailService
	redisStorage *redis.Client
}

func (ps *PasswordService) SendPasswordResetToken(email string) error {
	// TODO: check if user exist in DB !
	var hashedPassword, _ = bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)
	var user = model.User{
		Id:           123,
		PasswordHash: string(hashedPassword),
		Email:        "test",
	}

	// generate resetToken and store
	toBeHashed := user.PasswordHash + user.Email + time.Now().String()
	resetToken, err := bcrypt.GenerateFromPassword([]byte(toBeHashed), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("internal server error")
	}
	ps.redisStorage.SetNX(context.Background(), string(resetToken), user.Email, time.Duration(15*float64(time.Minute)))

	// send resetToken to user
	// wtf
	ps.mailService.SendMailPasswordResetToken(user, string(resetToken))

	return nil
}

func (ps *PasswordService) ResetPasswordWithToken(newPassword, token string) error {
	// check if token exists
	result := ps.redisStorage.Get(context.Background(), token)

	// TODO: remove token

	// TODO: generate new pass

	// TODO: send password

	return nil
}
