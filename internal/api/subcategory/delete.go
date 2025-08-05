package subcategory

import (
	"context"

	desc "github.com/Kosfedev/learn_go/pkg/subcategory_v1"
)

func (subcategoryImpl *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*desc.DeleteResponse, error) {
	err := subcategoryImpl.subcategoryService.Delete(ctx, int64(req.Id))
	if err != nil {
		return nil, err
	}

	return nil, nil
}
