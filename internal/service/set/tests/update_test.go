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

func TestUpdate(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	id := gofakeit.Int64()
	name := gofakeit.Question()
	req := &model.UpdatedSet{
		Name: &name,
	}
	mc := minimock.NewController(t)
	mockRepo := mocks.NewSetRepositoryMock(mc)
	mockRepo.UpdateMock.Expect(ctx, id, req).Return(nil)
	setService := set.NewService(mockRepo)

	t.Run("Update placeholder implementation test", func(t *testing.T) {
		err := setService.Update(context.Background(), id, req)
		if err != nil {
			t.Errorf("SetService.Update() error = %v\n", err)
		}
	})
}
