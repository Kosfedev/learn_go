package question

import (
	"context"
)

func (qs *serv) Delete(ctx context.Context, id int64) error {
	err := qs.questionRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
