package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Kosfedev/learn_go/internal/model"
	desc "github.com/Kosfedev/learn_go/pkg/subcategory_v1"
)

func NewSubcategoryFromGRPC(req *desc.CreateRequest) *model.NewSubcategory {
	newSubcategory := &model.NewSubcategory{Name: req.Name, CategoryId: int(req.CategoryId)}

	return newSubcategory
}

func UpdatedSubcategoryFromGRPC(req *desc.UpdateRequest) *model.UpdatedSubcategory {
	updatedSubcategory := &model.UpdatedSubcategory{}

	if req.Name != nil {
		name := req.Name.GetValue()
		updatedSubcategory.Name = &name
	}

	if req.CategoryId != nil {
		domainId := int(req.CategoryId.GetValue())
		updatedSubcategory.CategoryId = &domainId
	}

	return updatedSubcategory
}

func SubcategoryToGRPC(subcategory *model.Subcategory) *desc.GetResponse {
	res := &desc.GetResponse{
		Id:         int64(subcategory.Id),
		Name:       subcategory.Name,
		CategoryId: int64(subcategory.CategoryId),
		CreatedAt:  timestamppb.New(subcategory.CreatedAt),
		UpdatedAt:  nil,
	}

	if subcategory.UpdatedAt != nil {
		res.UpdatedAt = timestamppb.New(*subcategory.UpdatedAt)
	}

	return res
}
