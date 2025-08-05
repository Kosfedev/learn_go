package question

import (
	"context"
)

func (qs *serv) AddSubcategories(ctx context.Context, questionID int64, subcategoryIDs []int64) error {
	err := qs.questionSubcategoryRepo.AddSubcategoriesToQuestion(ctx, questionID, subcategoryIDs)
	if err != nil {
		return err
	}

	return nil
}
