package tests

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/mocks"
	"github.com/Kosfedev/learn_go/internal/service/category"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	id := int(gofakeit.Int64())
	domainID := int(gofakeit.Int64())
	req := &model.NewCategory{
		Name:     gofakeit.Name(),
		DomainID: domainID,
	}
	mc := minimock.NewController(t)
	mockRepo := mocks.NewCategoryRepositoryMock(mc)
	mockRepo.CreateMock.Expect(ctx, req).Return(id, nil)
	categoryService := category.NewService(mockRepo)

	t.Run("Create placeholder implementation test", func(t *testing.T) {
		_, err := categoryService.Create(ctx, req)
		if err != nil {
			t.Errorf("CategoryService.Create() error = %v\n", err)
		}
	})
}
