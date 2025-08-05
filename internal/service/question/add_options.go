package question

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qs *serv) AddOptions(ctx context.Context, questionID int64, options []*model.NewQuestionOption) error {
	err := qs.questionRepo.AddOptions(ctx, questionID, options)
	if err != nil {
		return err
	}

	return nil
}
