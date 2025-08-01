package Category

import (
	"context"
)

func (qcs *serv) Delete(ctx context.Context, id int) error {
	err := qcs.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
