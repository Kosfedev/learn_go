package category

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qcs *serv) Create(ctx context.Context, newCategory *model.NewCategory) (int64, error) {
	id, err := qcs.repo.Create(ctx, newCategory)
	if err != nil {
		return 0, err
	}

	return id, nil
}
