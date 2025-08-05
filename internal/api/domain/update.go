package domain

import (
	"context"
	"log"

	"github.com/Kosfedev/learn_go/internal/api/converter"
	desc "github.com/Kosfedev/learn_go/pkg/domain_v1"
)

func (domainImpl *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	updatedDomain := converter.UpdatedDomainFromGRPC(req)
	log.Printf("updated domain: %+v", updatedDomain)
	err := domainImpl.domainService.Update(ctx, int64(req.Id), updatedDomain)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
