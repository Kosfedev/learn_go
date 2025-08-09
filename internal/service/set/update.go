package set

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (s *serv) Update(ctx context.Context, id int64, updatedSet *model.UpdatedSet) error {
	err := s.repo.Update(ctx, id, updatedSet)
	if err != nil {
		return err
	}

	return nil
}
