package Category

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qcs *serv) Get(ctx context.Context, id int) (*model.Category, error) {
	category, err := qcs.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return category, nil
}
