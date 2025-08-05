package subcategory

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qss *serv) Create(ctx context.Context, newSubcategory *model.NewSubcategory) (int64, error) {
	id, err := qss.repo.Create(ctx, newSubcategory)
	if err != nil {
		return 0, err
	}

	return id, nil
}
