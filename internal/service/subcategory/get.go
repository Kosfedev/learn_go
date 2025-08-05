package subcategory

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qss *serv) Get(ctx context.Context, id int64) (*model.Subcategory, error) {
	category, err := qss.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return category, nil
}
