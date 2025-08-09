package questionformupdater

import (
	"context"
)

func (s *serv) AddSets(ctx context.Context, questionID int64, setIDs []int64) error {
	err := s.questionSetRepo.AddQuestionToSets(ctx, questionID, setIDs)
	if err != nil {
		return err
	}

	return nil
}
