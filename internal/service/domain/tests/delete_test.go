package tests

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"

	"github.com/Kosfedev/learn_go/internal/repository/mocks"
	"github.com/Kosfedev/learn_go/internal/service/domain"
)

func TestDelete(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	id := gofakeit.Int64()
	mc := minimock.NewController(t)
	mockRepo := mocks.NewDomainRepositoryMock(mc)
	mockRepo.DeleteMock.Expect(ctx, id).Return(nil)
	domainService := domain.NewService(mockRepo)

	t.Run("Delete placeholder implementation test", func(t *testing.T) {
		err := domainService.Delete(ctx, id)
		if err != nil {
			t.Errorf("DomainService.Delete() error = %v\n", err)
		}
	})
}
