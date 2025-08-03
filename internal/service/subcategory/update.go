package subcategory

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qss *serv) Update(ctx context.Context, id int, updatedCategory *model.UpdatedSubcategory) error {
	err := qss.repo.Update(ctx, id, updatedCategory)
	if err != nil {
		return err
	}

	return nil
}
