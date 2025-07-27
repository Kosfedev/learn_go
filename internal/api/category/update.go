package category

import (
	"context"
	"log"

	"github.com/Kosfedev/learn_go/internal/converter"
	desc "github.com/Kosfedev/learn_go/pkg/category_v1"
)

func (categoryImpl *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	updatedCategory := converter.UpdatedCategoryFromGRPC(req)
	log.Printf("updated domain: %+v", updatedCategory)
	err := categoryImpl.categoryService.Update(ctx, int(req.Id), updatedCategory)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
