package question

import (
	"context"
)

func (qs *serv) DeleteOptions(ctx context.Context, ids []int64) error {
	err := qs.questionOptionRepo.DeleteList(ctx, ids)
	if err != nil {
		return err
	}

	return nil
}
