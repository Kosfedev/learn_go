package question

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qs *serv) Update(_ context.Context, _ int, _ *model.UpdatedQuestion) error {
	return nil
}
