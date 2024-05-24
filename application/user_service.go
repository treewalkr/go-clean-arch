package application

import (
	"clean-arch/domain"
	"errors"
)

type UserService interface {
	GetUser(id int) (*domain.User, error)
	CreateUser(user *domain.User) error
}

type userService struct {
	users  map[int]*domain.User
	nextID int
}

func NewUserService() UserService {
	return &userService{
		users:  make(map[int]*domain.User),
		nextID: 1,
	}
}

func (s *userService) GetUser(id int) (*domain.User, error) {
	user, exists := s.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *userService) CreateUser(user *domain.User) error {
	user.ID = s.nextID
	s.users[s.nextID] = user
	s.nextID++
	return nil
}
