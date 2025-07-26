package Subategory

import (
	"context"
	"time"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qss *serv) Get(_ context.Context, categoryId int) (*model.Subcategory, error) {
	now := time.Now()
	dayBefore := now.AddDate(0, 0, -1)

	return &model.Subcategory{
		Id:         categoryId,
		Name:       gofakeit.Name(),
		CategoryId: int(gofakeit.Int64()),
		CreatedAt:  dayBefore,
		UpdatedAt:  &now,
	}, nil
}
