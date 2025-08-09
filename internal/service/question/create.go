package question

import (
	"context"
	"errors"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (s *serv) Create(ctx context.Context, newQuestion *model.NewQuestion) (int64, error) {
	if newQuestion == nil {
		return 0, errors.New("newQuestion cannot be nil")
	}

	if !newQuestion.Type.IsValid() {
		return 0, errors.New("invalid question type")
	}

	id, err := s.questionRepo.Create(ctx, newQuestion)
	if err != nil {
		return 0, err
	}

	return id, nil
}
