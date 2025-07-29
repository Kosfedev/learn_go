package question

import (
	"context"
)

func (qs *serv) Delete(ctx context.Context, id int) error {
	err := qs.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
