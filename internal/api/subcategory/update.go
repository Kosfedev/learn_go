package subcategory

import (
	"context"
	"log"

	"github.com/Kosfedev/learn_go/internal/api/converter"
	desc "github.com/Kosfedev/learn_go/pkg/subcategory_v1"
)

func (subcategoryImpl *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	updatedSubcategory := converter.UpdatedSubcategoryFromGRPC(req)
	log.Printf("updated subcategory: %+v", updatedSubcategory)
	err := subcategoryImpl.subcategoryService.Update(ctx, int64(req.Id), updatedSubcategory)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
