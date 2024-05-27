//go:build wireinject
// +build wireinject

package main

import (
	"clean-arch/adapters"
	"clean-arch/application"
	"clean-arch/domain"
	"clean-arch/infrastructure/persistence"

	"github.com/google/wire"
)

// ProvideInMemoryUserRepository creates a new in-memory user repository.
func ProvideInMemoryUserRepository() domain.UserRepository {
	return persistence.NewInMemoryUserRepository()
}

// ProvideUserService creates a new UserService with the given UserRepository.
func ProvideUserService(repo domain.UserRepository) application.UserService {
	return application.NewUserService(repo)
}

// ProvideUserHandler creates a new UserHandler with the given UserService.
func ProvideUserHandler(service application.UserService) *adapters.UserHandler {
	return adapters.NewUserHandler(service)
}

// InitializeApp sets up the entire application with dependencies.
func InitializeApp() (*adapters.UserHandler, error) {
	wire.Build(
		ProvideInMemoryUserRepository,
		ProvideUserService,
		ProvideUserHandler,
	)
	return &adapters.UserHandler{}, nil
}
