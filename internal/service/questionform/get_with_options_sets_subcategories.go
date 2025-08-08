package questionform

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (s *serv) GetWithOptionsSetsSubcategories(ctx context.Context, questionID int64) (*model.QuestionForm, error) {
	questionForm, err := s.repo.GetWithOptionsSetsSubcategories(ctx, questionID)
	if err != nil {
		return nil, err
	}

	return questionForm, nil
}
