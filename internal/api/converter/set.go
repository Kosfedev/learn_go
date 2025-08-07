package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Kosfedev/learn_go/internal/model"
	desc "github.com/Kosfedev/learn_go/pkg/set_v1"
)

func SetToGRPC(set *model.Set) *desc.GetResponse {
	res := &desc.GetResponse{
		Id:          set.ID,
		Name:        set.Name,
		QuestionIds: set.QuestionIDs,
		CreatedAt:   timestamppb.New(set.CreatedAt),
	}

	if set.UpdatedAt != nil {
		res.UpdatedAt = timestamppb.New(*set.UpdatedAt)
	}

	return res
}

func NewSetFromGRPC(req *desc.CreateRequest) *model.NewSet {
	newSet := &model.NewSet{
		Name:        req.Name,
		QuestionIDs: req.QuestionIds,
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
