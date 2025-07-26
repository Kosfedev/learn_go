package questionset

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/Kosfedev/learn_go/internal/model"
)

func TestCreate(t *testing.T) {
	questionSetService := NewService()
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

func TestGet(t *testing.T) {
	questionSetService := NewService()
	id := int(gofakeit.Int64())

	t.Run("Get placeholder implementation test", func(t *testing.T) {
		_, err := questionSetService.Get(context.Background(), id)
		if err != nil {
			t.Errorf("QuestionSetService.Create() error = %v\n", err)
		}
	})
}

func TestUpdate(t *testing.T) {
	questionSetService := NewService()
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

func TestDelete(t *testing.T) {
	questionSetService := NewService()
	id := int(gofakeit.Int64())

	t.Run("Delete placeholder implementation test", func(t *testing.T) {
		err := questionSetService.Delete(context.Background(), id)
		if err != nil {
			t.Errorf("QuestionSetService.Delete() error = %v\n", err)
		}
	})
}
