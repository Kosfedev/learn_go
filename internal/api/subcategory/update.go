package subcategory

import (
	"context"
	"log"

	"github.com/Kosfedev/learn_go/internal/converter"
	desc "github.com/Kosfedev/learn_go/pkg/subcategory_v1"
)

func (subcategoryImpl *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	updatedSubcategory := converter.UpdatedSubcategoryFromGRPC(req)
	log.Printf("updated subcategory: %+v", updatedSubcategory)
	err := subcategoryImpl.subcategoryService.Update(ctx, int(req.Id), updatedSubcategory)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
