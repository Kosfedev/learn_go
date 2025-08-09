package domain

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (s *serv) Update(ctx context.Context, id int64, updatedDomain *model.UpdatedDomain) error {
	err := s.repo.Update(ctx, id, updatedDomain)
	if err != nil {
		return err
	}

	return nil
}
