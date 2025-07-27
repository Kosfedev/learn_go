package category

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/converter"
	desc "github.com/Kosfedev/learn_go/pkg/category_v1"
)

func (categoryImpl *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	category, err := categoryImpl.categoryService.Get(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	res := converter.CategoryToGRPC(category)

	return res, nil
}
