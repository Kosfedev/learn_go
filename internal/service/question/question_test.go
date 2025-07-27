package question

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/Kosfedev/learn_go/internal/model"
)

func TestCreate(t *testing.T) {
	questionService := NewService()
	refAnswer := gofakeit.Quote()
	newQuestion := model.NewQuestion{
		Text:            gofakeit.Question(),
		Type:            3,
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
	questionService := NewService()
	text := gofakeit.Question()
	id := int(gofakeit.Int64())
	updatedQuestion := model.UpdatedQuestion{
		Text: &text,
	}

	t.Run("Update placeholder implementation test", func(t *testing.T) {
		err := questionService.Update(context.Background(), id, &updatedQuestion)
		if err != nil {
			t.Errorf("QuestionService.Update() error = %v\n", err)
		}
	})
}

func TestDelete(t *testing.T) {
	questionService := NewService()
	id := int(gofakeit.Int64())

	t.Run("Delete placeholder implementation test", func(t *testing.T) {
		err := questionService.Delete(context.Background(), id)
		if err != nil {
			t.Errorf("QuestionService.Delete() error = %v\n", err)
		}
	})
}

func TestAddOptions(t *testing.T) {
	questionService := NewService()
	id := int(gofakeit.Int64())
	newOptions := []*model.NewQuestionOption{
		{
			Text:      gofakeit.Question(),
			IsCorrect: false,
		},
		{
			Text:      gofakeit.Question(),
			IsCorrect: true,
		},
	}

	t.Run("Add options placeholder implementation test", func(t *testing.T) {
		err := questionService.AddOptions(context.Background(), id, newOptions)
		if err != nil {
			t.Errorf("QuestionService.AddOptions() error = %v\n", err)
		}
	})
}

func TestDeleteOptions(t *testing.T) {
	questionService := NewService()
	optionIds := make([]int, 10)
	gofakeit.Slice(&optionIds)

	t.Run("Delete options placeholder implementation test", func(t *testing.T) {
		err := questionService.DeleteOptions(context.Background(), optionIds)
		if err != nil {
			t.Errorf("QuestionService.DeleteOptions() error = %v\n", err)
		}
	})
}
