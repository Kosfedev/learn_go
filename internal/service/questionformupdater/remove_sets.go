package questionformupdater

import (
	"context"
)

func (s *serv) RemoveSets(ctx context.Context, questionID int64, setIDs []int64) error {
	err := s.questionSetRepo.RemoveQuestionFromSets(ctx, questionID, setIDs)
	if err != nil {
		return err
	}

	return nil
}
