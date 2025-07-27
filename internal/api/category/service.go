package category

import (
	"github.com/Kosfedev/learn_go/internal/service"
	desc "github.com/Kosfedev/learn_go/pkg/category_v1"
)

type Implementation struct {
	desc.UnimplementedCategoryV1Server
	categoryService service.CategoryService
}

func NewImplementation(categoryService service.CategoryService) *Implementation {
	return &Implementation{
		categoryService: categoryService,
	}
}
