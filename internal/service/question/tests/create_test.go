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

func TestCreate(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	id := int(gofakeit.Int64())
	refAnswer := gofakeit.Quote()
	req := &model.NewQuestion{
		Text:            gofakeit.Question(),
		Type:            3,
		ReferenceAnswer: &refAnswer,
	}
	mc := minimock.NewController(t)
	mockRepo := mocks.NewQuestionRepositoryMock(mc)
	mockRepo.CreateMock.Expect(ctx, req).Return(id, nil)
	questionService := question.NewService(mockRepo)

	t.Run("Create placeholder implementation test", func(t *testing.T) {
		_, err := questionService.Create(context.Background(), req)
		if err != nil {
			t.Errorf("QuestionService.Create() error = %v\n", err)
		}
	})
}
