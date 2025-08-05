package tests

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/mocks"
	"github.com/Kosfedev/learn_go/internal/service/domain"
)

func TestGet(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	id := gofakeit.Int64()
	res := &model.Domain{
		ID:        id,
		Name:      gofakeit.Name(),
		CreatedAt: time.Time{},
	}
	mc := minimock.NewController(t)
	mockRepo := mocks.NewDomainRepositoryMock(mc)
	mockRepo.GetMock.Expect(ctx, id).Return(res, nil)
	domainService := domain.NewService(mockRepo)

	t.Run("Get placeholder implementation test", func(t *testing.T) {
		_, err := domainService.Get(ctx, id)
		if err != nil {
			t.Errorf("DomainService.Get() error = %v\n", err)
		}
	})
}
