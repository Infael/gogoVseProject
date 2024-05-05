package password

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/Infael/gogoVseProject/model"
	"github.com/Infael/gogoVseProject/repository"
	"github.com/Infael/gogoVseProject/service/mail"
	"github.com/Infael/gogoVseProject/utils"
	"github.com/patrickmn/go-cache"
	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct {
	mailService    mail.MailService
	cache          *cache.Cache
	userRepository *repository.UserRepository
}

func NewPasswordService(mailService mail.MailService, cache *cache.Cache, userRepository *repository.UserRepository) *PasswordService {
	return &PasswordService{
		mailService:    mailService,
		cache:          cache,
		userRepository: userRepository,
	}
}

func (ps *PasswordService) SendPasswordResetToken(email string) error {
	user, err := ps.userRepository.GetUserByEmail(email)
	if err != nil {
		return err
	}

	// generate resetToken and store
	toBeHashed := time.Now().String() + strconv.FormatInt(int64(rand.Int()), 10)
	hash := sha256.New()
	hash.Write([]byte(toBeHashed))
	resetToken := hex.EncodeToString(hash.Sum(nil))
	if err != nil {
		return utils.InternalServerError(err)
	}

	ps.cache.Set(resetToken, user.Email, cache.DefaultExpiration)

	// send resetToken to user
	err = ps.mailService.SendMailPasswordResetToken(user, resetToken)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PasswordService) ResetPasswordWithToken(newPassword, token string) error {
	// check if token exists
	result, found := ps.cache.Get(token)
	if !found {
		return utils.ErrorNotFound(errors.New("Not Found SSS"))
	}

	// remove token
	ps.cache.Delete(token)

	// generate new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("XD!121: %s", err)
		return utils.InternalServerError(err)
	}

	// set new password
	user, err := ps.userRepository.GetUserByEmail(result.(string))
	if err != nil {
		log.Printf("XD!11: %s", err)
		return err
	}
	_, err = ps.userRepository.UpdateUser(&model.User{
		Email:        user.Email,
		PasswordHash: string(hashedPassword),
		Id:           user.Id,
	})

	return err
}
