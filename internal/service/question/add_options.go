package question

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qs *serv) AddOptions(_ context.Context, _ int, _ []*model.NewQuestionOption) error {
	return nil
}
