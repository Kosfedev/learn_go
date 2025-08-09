package question

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (s *serv) Get(ctx context.Context, questionID int64) (*model.Question, error) {
	question, err := s.questionRepo.Get(ctx, questionID)
	if err != nil {
		return nil, err
	}

	return question, nil
}
