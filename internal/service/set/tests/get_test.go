package tests

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/mocks"
	"github.com/Kosfedev/learn_go/internal/service/set"
)

func TestGet(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	id := gofakeit.Int64()
	res := &model.Set{
		ID:          id,
		Name:        gofakeit.Question(),
		QuestionIDs: []int64{1, 2, 3, 4, 5},
		CreatedAt:   time.Time{},
	}
	mc := minimock.NewController(t)
	mockRepo := mocks.NewQuestionSetRepositoryMock(mc)
	mockRepo.GetMock.Expect(ctx, id).Return(res, nil)
	setService := set.NewService(mockRepo)

	t.Run("Get placeholder implementation test", func(t *testing.T) {
		_, err := setService.Get(ctx, id)
		if err != nil {
			t.Errorf("SetService.Create() error = %v\n", err)
		}
	})
}
