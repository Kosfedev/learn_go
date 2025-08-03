package tests

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/mocks"
	"github.com/Kosfedev/learn_go/internal/service/questionset"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	id := int(gofakeit.Int64())
	req := &model.NewQuestionSet{
		Name: gofakeit.Question(),
	}
	mc := minimock.NewController(t)
	mockRepo := mocks.NewQuestionSetRepositoryMock(mc)
	mockRepo.CreateMock.Expect(ctx, req).Return(id, nil)
	questionSetService := questionset.NewService(mockRepo)

	t.Run("Create placeholder implementation test", func(t *testing.T) {
		_, err := questionSetService.Create(context.Background(), req)
		if err != nil {
			t.Errorf("QuestionSetService.Create() error = %v\n", err)
		}
	})
}
