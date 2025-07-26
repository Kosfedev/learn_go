package question

import (
	"context"
	"errors"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qs *serv) Create(_ context.Context, newQuestion *model.NewQuestion) (int, error) {
	if newQuestion == nil {
		return 0, errors.New("newQuestion cannot be nil")
	}

	if !newQuestion.Type.IsValid() {
		return 0, errors.New("invalid question type")
	}

	return int(gofakeit.Int64()), nil
}
