package question

import (
	"context"
)

func (qs *serv) DeleteOptions(ctx context.Context, ids []int64) error {
	err := qs.questionRepo.DeleteOptions(ctx, ids)
	if err != nil {
		return err
	}

	return nil
}
