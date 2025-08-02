package tests

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/service/questionset"
)

func TestUpdate(t *testing.T) {
	questionSetService := questionset.NewService()
	id := int(gofakeit.Int64())
	questionIds := make([]int, 10)
	gofakeit.Slice(&questionIds)
	updatedQuestionSet := model.UpdatedQuestionSet{
		QuestionIds: &questionIds,
	}

	t.Run("Update placeholder implementation test", func(t *testing.T) {
		err := questionSetService.Update(context.Background(), id, &updatedQuestionSet)
		if err != nil {
			t.Errorf("QuestionSetService.Update() error = %v\n", err)
		}
	})
}
