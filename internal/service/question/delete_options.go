package question

import (
	"context"
)

func (qs *serv) DeleteOptions(ctx context.Context, ids []int) error {
	err := qs.repo.DeleteOptions(ctx, ids)
	if err != nil {
		return err
	}

	return nil
}
