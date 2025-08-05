package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Kosfedev/learn_go/internal/model"
	desc "github.com/Kosfedev/learn_go/pkg/subcategory_v1"
)

func SubcategoryToGRPC(subcategory *model.Subcategory) *desc.GetResponse {
	res := &desc.GetResponse{
		Id:         subcategory.ID,
		Name:       subcategory.Name,
		CategoryId: subcategory.CategoryID,
		CreatedAt:  timestamppb.New(subcategory.CreatedAt),
		UpdatedAt:  nil,
	}

	if subcategory.UpdatedAt != nil {
		res.UpdatedAt = timestamppb.New(*subcategory.UpdatedAt)
	}

	return res
}

func NewSubcategoryFromGRPC(req *desc.CreateRequest) *model.NewSubcategory {
	newSubcategory := &model.NewSubcategory{Name: req.Name, CategoryID: req.CategoryId}

	return newSubcategory
}

func UpdatedSubcategoryFromGRPC(req *desc.UpdateRequest) *model.UpdatedSubcategory {
	updatedSubcategory := &model.UpdatedSubcategory{}

	if req.Name != nil {
		name := req.Name.GetValue()
		updatedSubcategory.Name = &name
	}

	if req.CategoryId != nil {
		domainId := req.CategoryId.GetValue()
		updatedSubcategory.CategoryID = &domainId
	}

	return updatedSubcategory
}
