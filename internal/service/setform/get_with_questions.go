package setform

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (s *serv) GetWithQuestions(ctx context.Context, setID int64) (*model.SetForm, error) {
	setForm, err := s.repo.GetWithQuestions(ctx, setID)
	if err != nil {
		return nil, err
	}

	return setForm, nil
}
