package questionset

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qss *serv) Update(ctx context.Context, id int64, updatedQuestionSet *model.UpdatedQuestionSet) error {
	err := qss.repo.Update(ctx, id, updatedQuestionSet)
	if err != nil {
		return err
	}

	return nil
}
