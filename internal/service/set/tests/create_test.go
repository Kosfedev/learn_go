package tests

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/mocks"
	"github.com/Kosfedev/learn_go/internal/service/set"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	id := gofakeit.Int64()
	req := &model.NewSet{
		Name: gofakeit.Question(),
	}
	mc := minimock.NewController(t)
	mockRepo := mocks.NewQuestionSetRepositoryMock(mc)
	mockRepo.CreateMock.Expect(ctx, req).Return(id, nil)
	setService := set.NewService(mockRepo)

	t.Run("Create placeholder implementation test", func(t *testing.T) {
		_, err := setService.Create(context.Background(), req)
		if err != nil {
			t.Errorf("SetService.Create() error = %v\n", err)
		}
	})
}
