package service

import (
	"github.com/lonelyRei/diplomApi/entities"
	"github.com/lonelyRei/diplomApi/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func (s *AuthService) CreateUser(user entities.User) (int, error) {
	// todo hash password before send
	return s.repo.CreateUser(user)
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}
