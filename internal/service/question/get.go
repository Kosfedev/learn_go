package question

import (
	"context"
	"time"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qs *serv) Get(_ context.Context, questionId int) (*model.Question, error) {
	now := time.Now()
	dayBefore := now.AddDate(0, 0, -1)
	options := []*model.QuestionOption{{
		Id:        0,
		Text:      gofakeit.Quote(),
		IsCorrect: false,
	}, {
		Id:        1,
		Text:      gofakeit.Quote(),
		IsCorrect: true,
	}, {
		Id:        2,
		Text:      gofakeit.Quote(),
		IsCorrect: false,
	}}

	return &model.Question{
		Id:              questionId,
		Text:            gofakeit.Question(),
		Type:            0,
		Options:         options,
		ReferenceAnswer: nil,
		CreatedAt:       dayBefore,
		UpdatedAt:       &now,
	}, nil
}
