package domain

import (
	"context"
)

func (qds *serv) Delete(ctx context.Context, id int) error {
	err := qds.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
