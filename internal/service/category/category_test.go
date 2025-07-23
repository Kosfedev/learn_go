package Category

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/Kosfedev/learn_go/internal/model"
)

func TestCreate(t *testing.T) {
	categoryService := NewService()
	newCategory := model.NewQuestionsCategory{
		Name:     gofakeit.Name(),
		DomainId: int(gofakeit.Int64()),
	}

	t.Run("Create placeholder implementation test", func(t *testing.T) {
		_, err := categoryService.Create(context.Background(), &newCategory)
		if err != nil {
			t.Errorf("QuestionsCategoryService.Create() error = %v\n", err)
		}
	})
}

func TestGet(t *testing.T) {
	categoryService := NewService()
	id := int(gofakeit.Int64())

	t.Run("Get placeholder implementation test", func(t *testing.T) {
		_, err := categoryService.Get(context.Background(), id)
		if err != nil {
			t.Errorf("QuestionsCategoryService.Create() error = %v\n", err)
		}
	})
}

func TestUpdate(t *testing.T) {
	categoryService := NewService()
	text := gofakeit.Question()
	id := int(gofakeit.Int64())
	updatedCategory := model.UpdatedQuestionsCategory{
		Name: &text,
	}

	t.Run("Update placeholder implementation test", func(t *testing.T) {
		err := categoryService.Update(context.Background(), id, &updatedCategory)
		if err != nil {
			t.Errorf("QuestionsCategoryService.Update() error = %v\n", err)
		}
	})
}

func TestDelete(t *testing.T) {
	categoryService := NewService()
	id := int(gofakeit.Int64())

	t.Run("Delete placeholder implementation test", func(t *testing.T) {
		err := categoryService.Delete(context.Background(), id)
		if err != nil {
			t.Errorf("QuestionsCategoryService.Delete() error = %v\n", err)
		}
	})
}
