package domain

import (
	"github.com/Kosfedev/learn_go/internal/service"
	desc "github.com/Kosfedev/learn_go/pkg/domain_v1"
)

type Implementation struct {
	desc.UnimplementedDomainV1Server
	domainService service.DomainService
}

func NewImplementation(domainService service.DomainService) *Implementation {
	return &Implementation{
		domainService: domainService,
	}
}
