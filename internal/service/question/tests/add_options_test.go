package tests

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/mocks"
	"github.com/Kosfedev/learn_go/internal/service/question"
)

func TestAddOptions(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	id := gofakeit.Int64()
	req := []*model.NewQuestionOption{
		{
			Text:      gofakeit.Question(),
			IsCorrect: false,
		},
		{
			Text:      gofakeit.Question(),
			IsCorrect: true,
		},
	}
	mc := minimock.NewController(t)
	mockQuestionSubcategoryRepo := mocks.NewQuestionSubcategoryRepositoryMock(mc)
	mockSubcategoryRepo := mocks.NewSubcategoryRepositoryMock(mc)
	mockQuestionRepo := mocks.NewQuestionRepositoryMock(mc)
	mockQuestionRepo.AddOptionsMock.Expect(ctx, id, req).Return(nil)
	questionService := question.NewService(mockQuestionRepo, mockQuestionSubcategoryRepo, mockSubcategoryRepo)

	t.Run("Add options placeholder implementation test", func(t *testing.T) {
		err := questionService.AddOptions(context.Background(), id, req)
		if err != nil {
			t.Errorf("QuestionService.AddOptions() error = %v\n", err)
		}
	})
}
