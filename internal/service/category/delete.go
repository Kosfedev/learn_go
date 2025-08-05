package category

import (
	"context"
)

func (qcs *serv) Delete(ctx context.Context, id int64) error {
	err := qcs.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
