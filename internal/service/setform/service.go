package setform

import (
	"github.com/Kosfedev/learn_go/internal/repository"
	"github.com/Kosfedev/learn_go/internal/service"
)

type serv struct {
	repo repository.SetFormRepository
}

func NewService(repo repository.SetFormRepository) service.SetFormService {
	return &serv{
		repo: repo,
	}
}
