package set

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (s *serv) ListSearch(ctx context.Context, listSearchOptions *model.SetListSearchOptions) (*model.SetListWithTotal, error) {
	setList, err := s.repo.ListSearch(ctx, listSearchOptions)
	if err != nil {
		return nil, err
	}

	return setList, nil
}
