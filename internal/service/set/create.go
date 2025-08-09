package set

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (s *serv) Create(ctx context.Context, newSet *model.NewSet) (int64, error) {
	id, err := s.repo.Create(ctx, newSet)
	if err != nil {
		return 0, err
	}

	return id, nil
}
