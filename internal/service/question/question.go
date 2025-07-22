package question

import (
	"context"
	"log"
	"time"

	"github.com/brianvoe/gofakeit"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/service"
)

type serv struct{}

func NewService() service.QuestionService {
	return &serv{}
}

func (qs *serv) Create(_ context.Context, newQuestion *model.NewQuestion) (int, error) {
	log.Printf("Creating question with data: %+v\n", newQuestion)

	return int(gofakeit.Int64()), nil
}

func (qs *serv) Get(_ context.Context, questionId int) (*model.Question, error) {
	log.Println("Getting question with id:", questionId)

	now := time.Now()
	dayBefore := now.AddDate(0, 0, -1)
	options := []model.QuestionOption{{
		Id:      0,
		Text:    gofakeit.Quote(),
		Correct: false,
	}, {
		Id:      1,
		Text:    gofakeit.Quote(),
		Correct: true,
	}, {
		Id:      2,
		Text:    gofakeit.Quote(),
		Correct: false,
	}}

	return &model.Question{
		Id:              questionId,
		Text:            gofakeit.Question(),
		Type:            0,
		Options:         &options,
		ReferenceAnswer: nil,
		CreatedAt:       dayBefore,
		UpdatedAt:       &now,
	}, nil
}

func (qs *serv) Update(_ context.Context, updatedQuestion *model.UpdatedQuestion) error {
	log.Printf("Updating question with data: %+v\n", updatedQuestion)
	return nil
}

func (qs *serv) Delete(_ context.Context, questionId int) error {
	log.Println("Deleting question with id:", questionId)
	return nil
}
