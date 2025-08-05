package domain

import (
	"context"
)

func (qds *serv) Delete(ctx context.Context, id int64) error {
	err := qds.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
