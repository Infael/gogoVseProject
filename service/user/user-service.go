package user

import (
	"strconv"
	"sync"
	"time"

	"github.com/Infael/gogoVseProject/model"
	"github.com/Infael/gogoVseProject/repository"
)

type UserService struct {
	repository *repository.UserRepository

	mu         sync.Mutex
	deleteJobs map[string]*time.Timer
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{
		repository: repository,
		deleteJobs: make(map[string]*time.Timer),
	}
}

func (s *UserService) GetUserByEmail(email string) (model.UserAll, error) {
	return s.repository.GetUserByEmail(email)
}

func (s *UserService) UpdateUser(userId uint64, updatedUser *model.UserAll) (model.UserAll, error) {
	return s.repository.UpdateUser(updatedUser)
}

func (s *UserService) DeleteUser(user model.UserAll) error {
	return s.repository.DeleteUser(user.Id)
}

func (s *UserService) ScheduleUserDeletion(email string, delay time.Duration) error {
	user, err := s.GetUserByEmail(email)
	if err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	userId := strconv.FormatUint(user.Id, 10)

	if timer, exists := s.deleteJobs[userId]; exists {
		timer.Stop()
	}

	s.deleteJobs[userId] = time.AfterFunc(delay, func() {
		s.DeleteUser(user)
	})

	return nil
}

func (s *UserService) CancelUserDeletion(email string) error {
	user, err := s.GetUserByEmail(email)
	if err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	userId := strconv.FormatUint(user.Id, 10)

	if timer, exists := s.deleteJobs[userId]; exists {
		timer.Stop()
		delete(s.deleteJobs, userId)
	}

	return nil
}
