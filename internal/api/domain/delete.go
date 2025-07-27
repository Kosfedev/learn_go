package domain

import (
	"context"

	desc "github.com/Kosfedev/learn_go/pkg/domain_v1"
)

func (domainImpl *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*desc.DeleteResponse, error) {
	err := domainImpl.domainService.Delete(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	return nil, nil
}
