package category

import (
	"github.com/Kosfedev/learn_go/internal/service"
	desc "github.com/Kosfedev/learn_go/pkg/subcategory_v1"
)

type Implementation struct {
	desc.UnimplementedSubcategoryV1Server
	subcategoryService service.SubcategoryService
}

func NewImplementation(subcategoryService service.SubcategoryService) *Implementation {
	return &Implementation{
		subcategoryService: subcategoryService,
	}
}
