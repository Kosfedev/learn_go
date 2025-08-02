package domain

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qds *serv) Create(ctx context.Context, newDomain *model.NewDomain) (int, error) {
	id, err := qds.repo.Create(ctx, newDomain)
	if err != nil {
		return 0, err
	}

	return id, err
}
