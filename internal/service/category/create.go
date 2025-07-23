package Category

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qcs *serv) Create(_ context.Context, _ *model.NewQuestionsCategory) (int, error) {
	return int(gofakeit.Int64()), nil
}
