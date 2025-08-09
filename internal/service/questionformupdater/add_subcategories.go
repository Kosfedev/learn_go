package questionformupdater

import (
	"context"
)

func (s *serv) AddSubcategories(ctx context.Context, questionID int64, subcategoryIDs []int64) error {
	err := s.questionSubcategoryRepo.AddSubcategoriesToQuestion(ctx, questionID, subcategoryIDs)
	if err != nil {
		return err
	}

	return nil
}
