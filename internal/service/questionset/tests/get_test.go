package tests

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/Kosfedev/learn_go/internal/service/questionset"
)

func TestGet(t *testing.T) {
	questionSetService := questionset.NewService()
	id := int(gofakeit.Int64())

	t.Run("Get placeholder implementation test", func(t *testing.T) {
		_, err := questionSetService.Get(context.Background(), id)
		if err != nil {
			t.Errorf("QuestionSetService.Create() error = %v\n", err)
		}
	})
}
