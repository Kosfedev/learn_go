package question

import (
	"context"
)

func (qs *serv) AddSubcategories(ctx context.Context, questionId int, subcategoryIds []int) error {
	err := qs.repo.AddSubcategories(ctx, questionId, subcategoryIds)
	if err != nil {
		return err
	}

	return nil
}
