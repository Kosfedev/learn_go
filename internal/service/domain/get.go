package domain

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qds *serv) Get(ctx context.Context, id int64) (*model.Domain, error) {
	domain, err := qds.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return domain, nil
}
