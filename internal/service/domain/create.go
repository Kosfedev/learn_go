package domain

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qds *serv) Create(_ context.Context, _ *model.NewQuestionsDomain) (int, error) {
	return int(gofakeit.Int64()), nil
}
