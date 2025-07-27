package domain

import (
	"context"
	"log"

	"github.com/Kosfedev/learn_go/internal/converter"
	desc "github.com/Kosfedev/learn_go/pkg/domain_v1"
)

func (domainImpl *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	newDomain := converter.NewDomainFromGRPC(req)
	log.Printf("new domain: %+v", newDomain)
	id, err := domainImpl.domainService.Create(ctx, newDomain)
	if err != nil {
		return nil, err
	}

	res := &desc.CreateResponse{
		Id: int64(id),
	}

	return res, nil
}
