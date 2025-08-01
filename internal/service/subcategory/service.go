package Subategory

import (
	"github.com/Kosfedev/learn_go/internal/repository"
	"github.com/Kosfedev/learn_go/internal/service"
)

type serv struct {
	repo repository.SubcategoryRepository
}

func NewService(repo repository.SubcategoryRepository) service.SubcategoryService {
	return &serv{repo: repo}
}
