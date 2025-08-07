package set

import (
	"context"

	desc "github.com/Kosfedev/learn_go/pkg/set_v1"

	"github.com/Kosfedev/learn_go/internal/api/converter"
)

func (questionSetImpl *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	set := converter.NewSetFromGRPC(req)
	id, err := questionSetImpl.setService.Create(ctx, set)
	if err != nil {
		return nil, err
	}

	res := &desc.CreateResponse{
		Id: id,
	}

	return res, nil
}
