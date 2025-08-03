package category

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qcs *serv) Update(ctx context.Context, id int, updatedCategory *model.UpdatedCategory) error {
	err := qcs.repo.Update(ctx, id, updatedCategory)
	if err != nil {
		return err
	}

	return nil
}
