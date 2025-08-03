package question

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qs *serv) Get(ctx context.Context, questionId int) (*model.Question, error) {
	question, err := qs.repo.Get(ctx, questionId)
	if err != nil {
		return nil, err
	}

	return question, nil
}
