package set

import (
	"context"

	desc "github.com/Kosfedev/learn_go/pkg/set_v1"

	"github.com/Kosfedev/learn_go/internal/api/converter"
)

func (questionSetImpl *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	updatedSet := converter.UpdatedSetFromGRPC(req)
	err := questionSetImpl.setService.Update(ctx, req.Id, updatedSet)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
