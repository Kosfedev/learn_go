package category

import (
	"context"
	"log"

	"github.com/Kosfedev/learn_go/internal/converter"
	desc "github.com/Kosfedev/learn_go/pkg/category_v1"
)

func (categoryImpl *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	newCategory := converter.NewCategoryFromGRPC(req)
	log.Printf("new category: %+v", newCategory)
	id, err := categoryImpl.categoryService.Create(ctx, newCategory)
	if err != nil {
		return nil, err
	}

	res := &desc.CreateResponse{
		Id: int64(id),
	}

	return res, nil
}
