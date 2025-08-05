package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Kosfedev/learn_go/internal/model"
	desc "github.com/Kosfedev/learn_go/pkg/category_v1"
)

func CategoryToGRPC(category *model.Category) *desc.GetResponse {
	res := &desc.GetResponse{
		Id:        int64(category.ID),
		Name:      category.Name,
		DomainId:  int64(category.DomainID),
		CreatedAt: timestamppb.New(category.CreatedAt),
		UpdatedAt: nil,
	}

	if category.UpdatedAt != nil {
		res.UpdatedAt = timestamppb.New(*category.UpdatedAt)
	}

	return res
}

func NewCategoryFromGRPC(req *desc.CreateRequest) *model.NewCategory {
	newCategory := &model.NewCategory{Name: req.Name, DomainID: int64(req.DomainId)}

	return newCategory
}

func UpdatedCategoryFromGRPC(req *desc.UpdateRequest) *model.UpdatedCategory {
	updatedCategory := &model.UpdatedCategory{}

	if req.Name != nil {
		name := req.Name.GetValue()
		updatedCategory.Name = &name
	}

	if req.DomainId != nil {
		domainId := int64(req.DomainId.GetValue())
		updatedCategory.DomainID = &domainId
	}

	return updatedCategory
}
