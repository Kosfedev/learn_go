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

func TestUpdate(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	id := int64(gofakeit.Int64())
	text := gofakeit.Sentence(5)
	req := &model.UpdatedDomain{
		Name: &text,
	}
	mc := minimock.NewController(t)
	mockRepo := mocks.NewDomainRepositoryMock(mc)
	mockRepo.UpdateMock.Expect(ctx, id, req).Return(nil)
	domainService := domain.NewService(mockRepo)

	t.Run("Update placeholder implementation test", func(t *testing.T) {
		err := domainService.Update(ctx, id, req)
		if err != nil {
			t.Errorf("DomainService.Update() error = %v\n", err)
		}
	})
}
