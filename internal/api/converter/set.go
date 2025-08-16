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

func SetListSearchOptionsFromGRPC(req *desc.ListSearchRequest) *model.SetListSearchOptions {
	setListSearchOptions := &model.SetListSearchOptions{
		Filters: &model.SetListSearchFilters{Name: req.Filters.Name},
		Pagination: &model.Pagination{
			Offset: req.Pagination.Offset,
			Limit:  req.Pagination.Limit,
		},
		Sort: &model.Sort{
			SortBy:    req.Sort.SortBy,
			SortOrder: model.SortOrder(req.Sort.SortOrder),
		},
	}

	return setListSearchOptions
}

func SetListResultsToGRPC(req *desc.ListSearchRequest, setListWithTotal *model.SetListWithTotal) *desc.ListSearchResponse {
	res := &desc.ListSearchResponse{
		Data: &desc.ListSearchResponse_Data{
			Sets: SetsToGRPC(setListWithTotal.Sets),
		},
		Meta: &desc.ListSearchResponse_Meta{
			Filters: req.Filters,
			Pagination: &desc.PaginationResponse{
				Limit:  req.Pagination.Limit,
				Offset: req.Pagination.Offset,
				Total:  setListWithTotal.Total,
			},
			Sort: req.Sort,
		},
	}

	return res
}
