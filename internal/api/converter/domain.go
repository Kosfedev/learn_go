package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Kosfedev/learn_go/internal/model"
	desc "github.com/Kosfedev/learn_go/pkg/domain_v1"
)

func DomainToGRPC(domain *model.Domain) *desc.GetResponse {
	res := &desc.GetResponse{
		Id:        domain.ID,
		Name:      domain.Name,
		CreatedAt: timestamppb.New(domain.CreatedAt),
		UpdatedAt: nil,
	}

	if domain.UpdatedAt != nil {
		res.UpdatedAt = timestamppb.New(*domain.UpdatedAt)
	}

	return res
}

func NewDomainFromGRPC(req *desc.CreateRequest) *model.NewDomain {
	newDomain := &model.NewDomain{Name: req.Name}

	return newDomain
}

func UpdatedDomainFromGRPC(req *desc.UpdateRequest) *model.UpdatedDomain {
	updatedDomain := &model.UpdatedDomain{}

	if req.Name != nil {
		name := req.Name.GetValue()
		updatedDomain.Name = &name
	}

	return updatedDomain
}
