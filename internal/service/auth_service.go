package service

import "go-gin-web/internal/repository"

type AuthService interface {
	login() string
}

type AuthServiceImpl struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &AuthServiceImpl{
		userRepository: userRepository,
	}
}

func (service AuthServiceImpl) login() string {
	return ""
}
