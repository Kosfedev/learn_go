package question

import (
	"context"
)

func (qs *serv) RemoveSubcategories(ctx context.Context, questionID int, subcategoryIDs []int) error {
	err := qs.questionSubcategoryRepo.RemoveSubcategoriesFromQuestion(ctx, questionID, subcategoryIDs)
	if err != nil {
		return err
	}

	return nil
}
