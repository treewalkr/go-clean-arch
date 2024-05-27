package persistence

import (
	"clean-arch/domain"
	"errors"
)

type InMemoryUserRepository struct {
	users  map[int]*domain.User
	nextID int
}

func NewInMemoryUserRepository() domain.UserRepository {
	return &InMemoryUserRepository{
		users:  make(map[int]*domain.User),
		nextID: 1,
	}
}

func (r *InMemoryUserRepository) GetByID(id int) (*domain.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *InMemoryUserRepository) Create(user *domain.User) error {
	user.ID = r.nextID
	r.users[r.nextID] = user
	r.nextID++
	return nil
}
