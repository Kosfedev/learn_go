package questionset

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qss *serv) Create(_ context.Context, newQuestionSet *model.NewQuestionSet) (int, error) {
	return int(gofakeit.Int64()), nil
}
