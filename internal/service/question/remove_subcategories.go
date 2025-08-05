package question

import (
	"context"
)

func (qs *serv) RemoveSubcategories(ctx context.Context, questionID int64, subcategoryIDs []int64) error {
	err := qs.questionSubcategoryRepo.RemoveSubcategoriesFromQuestion(ctx, questionID, subcategoryIDs)
	if err != nil {
		return err
	}

	return nil
}
