package tests

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"

	"github.com/Kosfedev/learn_go/internal/repository/mocks"
	"github.com/Kosfedev/learn_go/internal/service/question"
)

func TestDelete(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	id := int(gofakeit.Int64())
	mc := minimock.NewController(t)
	mockRepo := mocks.NewQuestionRepositoryMock(mc)
	mockRepo.DeleteMock.Expect(ctx, id).Return(nil)
	questionService := question.NewService(mockRepo)

	t.Run("Delete placeholder implementation test", func(t *testing.T) {
		err := questionService.Delete(context.Background(), id)
		if err != nil {
			t.Errorf("QuestionService.Delete() error = %v\n", err)
		}
	})
}
