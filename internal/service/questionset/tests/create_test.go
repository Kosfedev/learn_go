package tests

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/service/questionset"
)

func TestCreate(t *testing.T) {
	questionSetService := questionset.NewService()
	questionIds := make([]int, 10)
	gofakeit.Slice(&questionIds)
	newQuestionSet := model.NewQuestionSet{
		Name:        gofakeit.Name(),
		QuestionIds: questionIds,
	}

	t.Run("Create placeholder implementation test", func(t *testing.T) {
		_, err := questionSetService.Create(context.Background(), &newQuestionSet)
		if err != nil {
			t.Errorf("QuestionSetService.Create() error = %v\n", err)
		}
	})
}
