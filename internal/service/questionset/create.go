package questionset

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qss *serv) Create(ctx context.Context, newQuestionSet *model.NewQuestionSet) (int64, error) {
	id, err := qss.repo.Create(ctx, newQuestionSet)
	if err != nil {
		return 0, err
	}

	return id, nil
}
