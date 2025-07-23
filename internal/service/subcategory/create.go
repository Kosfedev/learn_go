package Subategory

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qss *serv) Create(_ context.Context, _ *model.NewSubcategory) (int, error) {
	return int(gofakeit.Int64()), nil
}
