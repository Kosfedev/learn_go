package set

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.Set, error) {
	set, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return set, nil
}
