package set

import (
	"github.com/Kosfedev/learn_go/internal/repository"
	"github.com/Kosfedev/learn_go/internal/service"
)

type serv struct {
	repo repository.SetRepository
}

func NewService(repo repository.SetRepository) service.SetService {
	return &serv{repo: repo}
}
