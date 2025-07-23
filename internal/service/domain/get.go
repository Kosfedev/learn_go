package domain

import (
	"context"
	"time"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qds *serv) Get(_ context.Context, domainId int) (*model.QuestionsDomain, error) {
	now := time.Now()
	dayBefore := now.AddDate(0, 0, -1)

	return &model.QuestionsDomain{
		Id:        domainId,
		Name:      gofakeit.Name(),
		CreatedAt: dayBefore,
		UpdatedAt: &now,
	}, nil
}
