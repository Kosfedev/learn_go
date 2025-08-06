package tests

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/mocks"
	"github.com/Kosfedev/learn_go/internal/service/question"
)

func TestGet(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	id := gofakeit.Int64()
	res := &model.Question{
		ID:              id,
		Text:            gofakeit.Question(),
		Type:            3,
		Options:         nil,
		ReferenceAnswer: nil,
		CreatedAt:       time.Time{},
		UpdatedAt:       nil,
	}
	mc := minimock.NewController(t)
	mockQuestionSubcategoryRepo := mocks.NewQuestionSubcategoryRepositoryMock(mc)
	mockSubcategoryRepo := mocks.NewSubcategoryRepositoryMock(mc)
	mockQuestionSetRepo := mocks.NewQuestionSetRepositoryMock(mc)
	mockSetRepo := mocks.NewSetRepositoryMock(mc)
	mockQuestionRepo := mocks.NewQuestionRepositoryMock(mc)
	mockQuestionRepo.GetMock.Expect(ctx, id).Return(res, nil)
	questionService := question.NewService(mockQuestionRepo, mockQuestionSetRepo, mockQuestionSubcategoryRepo, mockSetRepo, mockSubcategoryRepo)

	t.Run("Get placeholder implementation test", func(t *testing.T) {
		_, err := questionService.Get(ctx, id)
		if err != nil {
			t.Errorf("QuestionService.Create() error = %v\n", err)
		}
	})
}
