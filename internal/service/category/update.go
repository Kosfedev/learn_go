package Category

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qcs *serv) Update(_ context.Context, _ int, _ *model.UpdatedQuestionsCategory) error {
	return nil
}
