package tests

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"

	"github.com/Kosfedev/learn_go/internal/repository/mocks"
	"github.com/Kosfedev/learn_go/internal/service/set"
)

func TestDelete(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	id := gofakeit.Int64()
	mc := minimock.NewController(t)
	mockRepo := mocks.NewQuestionSetRepositoryMock(mc)
	mockRepo.DeleteMock.Expect(ctx, id).Return(nil)
	setService := set.NewService(mockRepo)

	t.Run("Delete placeholder implementation test", func(t *testing.T) {
		err := setService.Delete(context.Background(), id)
		if err != nil {
			t.Errorf("SetService.Delete() error = %v\n", err)
		}
	})
}
