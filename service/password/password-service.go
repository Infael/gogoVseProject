package password

import (
	"context"
	"time"

	"github.com/Infael/gogoVseProject/model"
	"github.com/Infael/gogoVseProject/repository"
	"github.com/Infael/gogoVseProject/service/mail"
	"github.com/Infael/gogoVseProject/utils"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct {
	mailService    mail.MailService
	redisStorage   *redis.Client
	userRepository *repository.UserRepository
}

func NewPasswordService(mailService mail.MailService, redisStorage *redis.Client, userRepository *repository.UserRepository) *PasswordService {
	return &PasswordService{
		mailService:    mailService,
		redisStorage:   redisStorage,
		userRepository: userRepository,
	}
}

func (ps *PasswordService) SendPasswordResetToken(email string) error {
	user, err := ps.userRepository.GetUserByEmail(email)
	if err != nil {
		return err
	}

	// generate resetToken and store
	toBeHashed := user.PasswordHash + user.Email + time.Now().String()
	resetToken, err := bcrypt.GenerateFromPassword([]byte(toBeHashed), bcrypt.DefaultCost)
	if err != nil {
		return utils.InternalServerError(err)
	}

	err = ps.redisStorage.Set(context.Background(), string(resetToken), user.Email, 1200).Err()
	if err != nil {
		return utils.InternalServerError(err)
	}

	// send resetToken to user
	ps.mailService.SendMailPasswordResetToken(user, string(resetToken))

	return nil
}

func (ps *PasswordService) ResetPasswordWithToken(newPassword, token string) error {
	// check if token exists
	result, err := ps.redisStorage.Get(context.Background(), token).Result()
	if err != nil {
		return utils.InternalServerError(err)
	}

	// remove token
	err = ps.redisStorage.Del(context.Background(), token).Err()
	if err != nil {
		return utils.InternalServerError(err)
	}

	// generate new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return utils.InternalServerError(err)
	}

	// set new password
	user, err := ps.userRepository.GetUserByEmail(result)
	if err != nil {
		return err
	}
	_, err = ps.userRepository.UpdateUser(&model.User{
		Email:        user.Email,
		PasswordHash: string(hashedPassword),
		Id:           user.Id,
	})

	return err
}
