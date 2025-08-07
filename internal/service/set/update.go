package set

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qss *serv) Update(ctx context.Context, id int64, updatedSet *model.UpdatedSet) error {
	err := qss.repo.Update(ctx, id, updatedSet)
	if err != nil {
		return err
	}

	return nil
}
