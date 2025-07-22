package question

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit"

	"github.com/Kosfedev/learn_go/internal/model"
)

func TestCreate(t *testing.T) {
	questionService := NewService()
	refAnswer := gofakeit.Quote()

	newQuestion := model.NewQuestion{
		Text:            gofakeit.Question(),
		Type:            3,
		Options:         nil,
		ReferenceAnswer: &refAnswer,
	}

	t.Run("Create placeholder implementation test", func(t *testing.T) {
		_, err := questionService.Create(context.Background(), &newQuestion)
		if err != nil {
			t.Errorf("QuestionService.Create() error = %v\n", err)
		}
	})
}

func TestGet(t *testing.T) {
	questionService := NewService()
	id := int(gofakeit.Int64())

	t.Run("Get placeholder implementation test", func(t *testing.T) {
		_, err := questionService.Get(context.Background(), id)
		if err != nil {
			t.Errorf("QuestionService.Create() error = %v\n", err)
		}
	})
}

func TestUpdate(t *testing.T) {
	text := gofakeit.Question()
	questionService := NewService()
	id := int(gofakeit.Int64())
	updatedQuestion := model.UpdatedQuestion{
		Text: &text,
	}

	t.Run("Get placeholder implementation test", func(t *testing.T) {
		err := questionService.Update(context.Background(), id, &updatedQuestion)
		if err != nil {
			t.Errorf("QuestionService.Update() error = %v\n", err)
		}
	})
}

func TestDelete(t *testing.T) {
	questionService := NewService()
	id := int(gofakeit.Int64())

	t.Run("Get placeholder implementation test", func(t *testing.T) {
		err := questionService.Delete(context.Background(), id)
		if err != nil {
			t.Errorf("QuestionService.Delete() error = %v\n", err)
		}
	})
}
