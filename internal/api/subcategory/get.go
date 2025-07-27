package subcategory

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/converter"
	desc "github.com/Kosfedev/learn_go/pkg/subcategory_v1"
)

func (subcategoryImpl *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	subcategory, err := subcategoryImpl.subcategoryService.Get(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	res := converter.SubcategoryToGRPC(subcategory)

	return res, nil
}
