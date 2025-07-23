package domain

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/Kosfedev/learn_go/internal/model"
)

func TestCreate(t *testing.T) {
	questionService := NewService()
	newDomain := model.NewDomain{
		Name: gofakeit.Name(),
	}

	t.Run("Create placeholder implementation test", func(t *testing.T) {
		_, err := questionService.Create(context.Background(), &newDomain)
		if err != nil {
			t.Errorf("DomainService.Create() error = %v\n", err)
		}
	})
}

func TestGet(t *testing.T) {
	questionService := NewService()
	id := int(gofakeit.Int64())

	t.Run("Get placeholder implementation test", func(t *testing.T) {
		_, err := questionService.Get(context.Background(), id)
		if err != nil {
			t.Errorf("DomainService.Create() error = %v\n", err)
		}
	})
}

func TestUpdate(t *testing.T) {
	questionService := NewService()
	text := gofakeit.Question()
	id := int(gofakeit.Int64())
	updatedDomain := model.UpdatedDomain{
		Name: &text,
	}

	t.Run("Update placeholder implementation test", func(t *testing.T) {
		err := questionService.Update(context.Background(), id, &updatedDomain)
		if err != nil {
			t.Errorf("DomainService.Update() error = %v\n", err)
		}
	})
}

func TestDelete(t *testing.T) {
	questionService := NewService()
	id := int(gofakeit.Int64())

	t.Run("Delete placeholder implementation test", func(t *testing.T) {
		err := questionService.Delete(context.Background(), id)
		if err != nil {
			t.Errorf("DomainService.Delete() error = %v\n", err)
		}
	})
}
