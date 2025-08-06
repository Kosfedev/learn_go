package set

import (
	"context"

	desc "github.com/Kosfedev/learn_go/pkg/set_v1"
)

func (questionSetImpl *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*desc.DeleteResponse, error) {
	err := questionSetImpl.setService.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
