package set

import (
	"context"

	desc "github.com/Kosfedev/learn_go/pkg/set_v1"

	"github.com/Kosfedev/learn_go/internal/api/converter"
)

func (questionSetImpl *Implementation) ListSearch(ctx context.Context, req *desc.ListSearchRequest) (*desc.ListSearchResponse, error) {
	setListSearchOptions := converter.SetListSearchOptionsFromGRPC(req)
	setListWithTotal, err := questionSetImpl.setService.ListSearch(ctx, setListSearchOptions)
	if err != nil {
		return nil, err
	}

	res := converter.SetListResultsToGRPC(req, setListWithTotal)

	return res, nil
}
