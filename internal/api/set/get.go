package set

import (
	"context"

	desc "github.com/Kosfedev/learn_go/pkg/set_v1"

	"github.com/Kosfedev/learn_go/internal/api/converter"
)

func (questionSetImpl *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	set, err := questionSetImpl.setService.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	res := converter.SetToGRPC(set)

	return res, nil
}
