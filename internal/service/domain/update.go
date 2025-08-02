package domain

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qds *serv) Update(ctx context.Context, id int, updatedDomain *model.UpdatedDomain) error {
	err := qds.repo.Update(ctx, id, updatedDomain)
	if err != nil {
		return err
	}

	return nil
}
