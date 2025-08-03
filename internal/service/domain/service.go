package domain

import (
	"github.com/Kosfedev/learn_go/internal/repository"
	"github.com/Kosfedev/learn_go/internal/service"
)

type serv struct {
	repo repository.DomainRepository
}

func NewService(repo repository.DomainRepository) service.DomainService {
	return &serv{
		repo: repo,
	}
}
