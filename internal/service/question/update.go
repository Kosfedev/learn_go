package question

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qs *serv) Update(ctx context.Context, id int, updatedQuestion *model.UpdatedQuestion) error {
	err := qs.repo.Update(ctx, id, updatedQuestion)
	if err != nil {
		return err
	}

	return nil
}
