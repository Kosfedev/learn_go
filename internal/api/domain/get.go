package domain

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/api/converter"
	desc "github.com/Kosfedev/learn_go/pkg/domain_v1"
)

func (domainImpl *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	domain, err := domainImpl.domainService.Get(ctx, int64(req.Id))
	if err != nil {
		return nil, err
	}

	res := converter.DomainToGRPC(domain)

	return res, nil
}
