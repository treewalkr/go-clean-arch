package application

import (
	"clean-arch/domain"
)

type UserService interface {
	GetUser(id int) (*domain.User, error)
	CreateUser(user *domain.User) error
}

type userService struct {
	userRepo domain.UserRepository
}

func NewUserService(userRepo domain.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) GetUser(id int) (*domain.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *userService) CreateUser(user *domain.User) error {
	return s.userRepo.Create(user)
}
