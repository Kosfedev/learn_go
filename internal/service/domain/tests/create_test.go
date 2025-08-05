package tests

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/mocks"
	"github.com/Kosfedev/learn_go/internal/service/domain"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	id := gofakeit.Int64()
	req := &model.NewDomain{
		Name: gofakeit.Name(),
	}
	mc := minimock.NewController(t)
	mockRepo := mocks.NewDomainRepositoryMock(mc)
	mockRepo.CreateMock.Expect(ctx, req).Return(id, nil)
	domainService := domain.NewService(mockRepo)

	t.Run("Create placeholder implementation test", func(t *testing.T) {
		_, err := domainService.Create(ctx, req)
		if err != nil {
			t.Errorf("DomainService.Create() error = %v\n", err)
		}
	})
}
