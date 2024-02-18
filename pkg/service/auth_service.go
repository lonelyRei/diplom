package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/lonelyRei/diplomApi/entities"
	"github.com/lonelyRei/diplomApi/pkg/repository"
)

const (
	salt = "cskdcmksckldsmkcd3991&&&!/"
)

type AuthService struct {
	repo repository.Authorization
}

func (s *AuthService) CreateUser(user entities.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) Test(test entities.Test) (int, error) {
	return s.repo.Test(test)
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
