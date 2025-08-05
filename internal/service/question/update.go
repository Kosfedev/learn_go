package question

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qs *serv) Update(ctx context.Context, id int64, updatedQuestion *model.UpdatedQuestion) error {
	err := qs.questionRepo.Update(ctx, id, updatedQuestion)
	if err != nil {
		return err
	}

	return nil
}
