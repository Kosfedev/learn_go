package Category

import (
	"context"
	"time"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qcs *serv) Get(_ context.Context, categoryId int) (*model.Category, error) {
	now := time.Now()
	dayBefore := now.AddDate(0, 0, -1)

	return &model.Category{
		Id:        categoryId,
		Name:      gofakeit.Name(),
		DomainId:  int(gofakeit.Int64()),
		CreatedAt: dayBefore,
		UpdatedAt: &now,
	}, nil
}
