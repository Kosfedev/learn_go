package domain

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (s *serv) Create(ctx context.Context, newDomain *model.NewDomain) (int64, error) {
	id, err := s.repo.Create(ctx, newDomain)
	if err != nil {
		return 0, err
	}

	return id, err
}
