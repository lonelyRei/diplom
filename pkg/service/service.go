package service

import "github.com/lonelyRei/diplomApi/pkg/repository"

type Service struct {
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
