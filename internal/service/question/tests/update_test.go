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

func TestUpdate(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	id := gofakeit.Int64()
	refAnswer := gofakeit.Quote()
	req := &model.UpdatedQuestion{
		ReferenceAnswer: &refAnswer,
	}
	mc := minimock.NewController(t)
	mockQuestionRepo := mocks.NewQuestionRepositoryMock(mc)
	mockQuestionRepo.UpdateMock.Expect(ctx, id, req).Return(nil)
	questionService := question.NewService(mockQuestionRepo)

	t.Run("Update placeholder implementation test", func(t *testing.T) {
		err := questionService.Update(context.Background(), id, req)
		if err != nil {
			t.Errorf("QuestionService.Update() error = %v\n", err)
		}
	})
}
