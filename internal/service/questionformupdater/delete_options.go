package questionformupdater

import (
	"context"
)

func (s *serv) DeleteOptions(ctx context.Context, optionIDs []int64) error {
	err := s.questionOptionRepo.DeleteList(ctx, optionIDs)
	if err != nil {
		return err
	}

	return nil
}
