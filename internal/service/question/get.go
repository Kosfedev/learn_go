package question

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qs *serv) Get(ctx context.Context, questionID int64) (*model.Question, error) {
	question, err := qs.questionRepo.Get(ctx, questionID)
	if err != nil {
		return nil, err
	}

	return question, nil
}
