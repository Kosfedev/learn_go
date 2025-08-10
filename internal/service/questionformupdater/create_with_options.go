package questionformupdater

import (
	"context"
	"errors"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (s *serv) CreateWithOptions(ctx context.Context, newQuestion *model.NewQuestionWithOptions) (int64, error) {
	if newQuestion == nil {
		return 0, errors.New("newQuestion cannot be nil")
	}

	if !newQuestion.Question.Type.IsValid() {
		return 0, errors.New("invalid question type")
	}

	var id int64
	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var err error
		id, err = s.questionRepo.Create(ctx, newQuestion.Question)
		if err != nil {
			return err
		}

		err = s.questionOptionRepo.CreateList(ctx, id, newQuestion.Options)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
