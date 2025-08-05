package question

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qs *serv) AddOptions(ctx context.Context, questionId int, options []*model.NewQuestionOption) error {
	err := qs.questionRepo.AddOptions(ctx, questionId, options)
	if err != nil {
		return err
	}

	return nil
}
