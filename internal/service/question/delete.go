package question

import (
	"context"
)

func (s *serv) Delete(ctx context.Context, id int64) error {
	err := s.questionRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
