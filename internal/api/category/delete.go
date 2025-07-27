package category

import (
	"context"

	desc "github.com/Kosfedev/learn_go/pkg/category_v1"
)

func (categoryImpl *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*desc.DeleteResponse, error) {
	err := categoryImpl.categoryService.Delete(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	return nil, nil
}
