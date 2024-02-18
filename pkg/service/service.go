package service

import (
	"github.com/lonelyRei/diplomApi/entities"
	"github.com/lonelyRei/diplomApi/pkg/repository"
)

type Authorization interface {
	CreateUser(user entities.User) (int, error)
	Test(test entities.Test) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (*TokenClaims, error)
}

type Service struct {
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
