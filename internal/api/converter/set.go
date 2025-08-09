package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Kosfedev/learn_go/internal/model"
	desc "github.com/Kosfedev/learn_go/pkg/set_v1"
)

func SetToGRPC(set *model.Set) *desc.GetResponse {
	setGRPC := &desc.Set{
		Id:        set.ID,
		Name:      set.Name,
		CreatedAt: timestamppb.New(set.CreatedAt),
	}

	if set.UpdatedAt != nil {
		setGRPC.UpdatedAt = timestamppb.New(*set.UpdatedAt)
	}

	res := &desc.GetResponse{
		Data: setGRPC,
	}

	return res
}

func NewSetFromGRPC(req *desc.CreateRequest) *model.NewSet {
	newSet := &model.NewSet{
		Name: req.Name,
	}

	return newSet
}

func UpdatedSetFromGRPC(req *desc.UpdateRequest) *model.UpdatedSet {
	updatedSet := &model.UpdatedSet{}

	if req.Name != nil {
		name := req.Name.GetValue()
		updatedSet.Name = &name
	}

	return updatedSet
}
