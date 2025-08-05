package tests

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/mocks"
	"github.com/Kosfedev/learn_go/internal/service/subcategory"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	id := int(gofakeit.Int64())
	categoryID := int(gofakeit.Int64())
	req := &model.NewSubcategory{
		Name:       gofakeit.Name(),
		CategoryID: categoryID,
	}
	mc := minimock.NewController(t)
	mockRepo := mocks.NewSubcategoryRepositoryMock(mc)
	mockRepo.CreateMock.Expect(ctx, req).Return(id, nil)
	subcategoryService := subcategory.NewService(mockRepo)

	t.Run("Create placeholder implementation test", func(t *testing.T) {
		_, err := subcategoryService.Create(ctx, req)
		if err != nil {
			t.Errorf("SubcategoryService.Create() error = %v\n", err)
		}
	})
}
