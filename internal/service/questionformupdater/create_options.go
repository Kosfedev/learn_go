package questionformupdater

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (s *serv) CreateOptions(ctx context.Context, questionID int64, options []*model.NewQuestionOption) error {
	err := s.questionOptionRepo.CreateList(ctx, questionID, options)
	if err != nil {
		return err
	}

	return nil
}
