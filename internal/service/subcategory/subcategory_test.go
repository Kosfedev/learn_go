package Subategory

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/Kosfedev/learn_go/internal/model"
)

func TestCreate(t *testing.T) {
	subcategoryService := NewService()
	newSubcategory := model.NewSubcategory{
		Name:       gofakeit.Name(),
		CategoryId: int(gofakeit.Int64()),
	}

	t.Run("Create placeholder implementation test", func(t *testing.T) {
		_, err := subcategoryService.Create(context.Background(), &newSubcategory)
		if err != nil {
			t.Errorf("SubcategoryService.Create() error = %v\n", err)
		}
	})
}

func TestGet(t *testing.T) {
	subcategoryService := NewService()
	id := int(gofakeit.Int64())

	t.Run("Get placeholder implementation test", func(t *testing.T) {
		_, err := subcategoryService.Get(context.Background(), id)
		if err != nil {
			t.Errorf("SubcategoryService.Create() error = %v\n", err)
		}
	})
}

func TestUpdate(t *testing.T) {
	subcategoryService := NewService()
	text := gofakeit.Question()
	id := int(gofakeit.Int64())
	updatedSubcategory := model.UpdatedSubcategory{
		Name: &text,
	}

	t.Run("Update placeholder implementation test", func(t *testing.T) {
		err := subcategoryService.Update(context.Background(), id, &updatedSubcategory)
		if err != nil {
			t.Errorf("SubcategoryService.Update() error = %v\n", err)
		}
	})
}

func TestDelete(t *testing.T) {
	subcategoryService := NewService()
	id := int(gofakeit.Int64())

	t.Run("Delete placeholder implementation test", func(t *testing.T) {
		err := subcategoryService.Delete(context.Background(), id)
		if err != nil {
			t.Errorf("SubcategoryService.Delete() error = %v\n", err)
		}
	})
}
