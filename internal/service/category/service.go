package Category

import (
	"github.com/Kosfedev/learn_go/internal/repository"
	"github.com/Kosfedev/learn_go/internal/service"
)

type serv struct {
	repo repository.CategoryRepository
}

func NewService(repo repository.CategoryRepository) service.CategoryService {
	return &serv{repo: repo}
}
