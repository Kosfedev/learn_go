package domain

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qds *serv) Update(_ context.Context, _ int, _ *model.UpdatedDomain) error {
	return nil
}
