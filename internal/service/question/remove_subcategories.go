package question

import (
	"context"
)

func (qs *serv) RemoveSubcategories(ctx context.Context, questionId int, subcategoryIds []int) error {
	err := qs.repo.RemoveSubcategories(ctx, questionId, subcategoryIds)
	if err != nil {
		return err
	}

	return nil
}
