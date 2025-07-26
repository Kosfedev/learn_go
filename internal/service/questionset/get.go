package questionset

import (
	"context"
	"time"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (qss *serv) Get(_ context.Context, questionSetId int) (*model.QuestionSet, error) {
	updatedAt := time.Now().AddDate(0, 0, -1)
	questionIds := make([]int, 10)
	gofakeit.Slice(&questionIds)
	questionSet := model.QuestionSet{
		Id:          questionSetId,
		Name:        gofakeit.Name(),
		QuestionIds: questionIds,
		CreatedAt:   updatedAt.AddDate(0, 0, -1),
		UpdatedAt:   &updatedAt,
	}

	return &questionSet, nil
}
