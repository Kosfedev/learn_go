package tests

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/Kosfedev/learn_go/internal/service/questionset"
)

func TestDelete(t *testing.T) {
	questionSetService := questionset.NewService()
	id := int(gofakeit.Int64())

	t.Run("Delete placeholder implementation test", func(t *testing.T) {
		err := questionSetService.Delete(context.Background(), id)
		if err != nil {
			t.Errorf("QuestionSetService.Delete() error = %v\n", err)
		}
	})
}
