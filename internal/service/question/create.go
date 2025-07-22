package question

import (
	"context"

	"github.com/brianvoe/gofakeit"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qs *serv) Create(_ context.Context, _ *model.NewQuestion) (int, error) {
	return int(gofakeit.Int64()), nil
}
