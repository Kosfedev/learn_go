package questionformupdater

import (
	"context"
)

func (s *serv) RemoveSubcategories(ctx context.Context, questionID int64, subcategoryIDs []int64) error {
	err := s.questionSubcategoryRepo.RemoveSubcategoriesFromQuestion(ctx, questionID, subcategoryIDs)
	if err != nil {
		return err
	}

	return nil
}
