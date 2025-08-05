package tests

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"

	"github.com/Kosfedev/learn_go/internal/repository/mocks"
	"github.com/Kosfedev/learn_go/internal/service/question"
)

func TestDeleteOptions(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	optionIds := make([]int, 10)
	gofakeit.Slice(&optionIds)
	mc := minimock.NewController(t)
	mockQuestionSubcategoryRepo := mocks.NewQuestionSubcategoryRepositoryMock(mc)
	mockSubcategoryRepo := mocks.NewSubcategoryRepositoryMock(mc)
	mockQuestionRepo := mocks.NewQuestionRepositoryMock(mc)
	mockQuestionRepo.DeleteOptionsMock.Expect(ctx, optionIds).Return(nil)
	questionService := question.NewService(mockQuestionRepo, mockQuestionSubcategoryRepo, mockSubcategoryRepo)

	t.Run("Delete options placeholder implementation test", func(t *testing.T) {
		err := questionService.DeleteOptions(context.Background(), optionIds)
		if err != nil {
			t.Errorf("QuestionService.DeleteOptions() error = %v\n", err)
		}
	})
}
