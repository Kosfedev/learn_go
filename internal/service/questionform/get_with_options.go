package questionform

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (s *serv) GetWithOptions(ctx context.Context, questionID int64) (*model.QuestionWithOptions, error) {
	questionWithOptions, err := s.repo.GetWithOptions(ctx, questionID)
	if err != nil {
		return nil, err
	}

	return questionWithOptions, nil
}
