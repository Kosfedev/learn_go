package question

import (
	"context"
)

func (qs *serv) AddSubcategories(ctx context.Context, questionID int, subcategoryIDs []int) error {
	err := qs.questionSubcategoryRepo.AddSubcategoriesToQuestion(ctx, questionID, subcategoryIDs)
	if err != nil {
		return err
	}

	return nil
}
