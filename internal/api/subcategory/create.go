package subcategory

import (
	"context"
	"log"

	"github.com/Kosfedev/learn_go/internal/converter"
	desc "github.com/Kosfedev/learn_go/pkg/subcategory_v1"
)

func (subcategoryImpl *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	newCategory := converter.NewSubcategoryFromGRPC(req)
	log.Printf("new subcategory: %+v", newCategory)
	id, err := subcategoryImpl.subcategoryService.Create(ctx, newCategory)
	if err != nil {
		return nil, err
	}

	res := &desc.CreateResponse{
		Id: int64(id),
	}

	return res, nil
}
