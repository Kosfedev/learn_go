package tests

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"

	"github.com/Kosfedev/learn_go/internal/repository/mocks"
	"github.com/Kosfedev/learn_go/internal/service/subcategory"
)

func TestDelete(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	id := int(gofakeit.Int64())
	mc := minimock.NewController(t)
	mockRepo := mocks.NewSubcategoryRepositoryMock(mc)
	mockRepo.DeleteMock.Expect(ctx, id).Return(nil)
	subcategoryService := subcategory.NewService(mockRepo)

	t.Run("Delete placeholder implementation test", func(t *testing.T) {
		err := subcategoryService.Delete(ctx, id)
		if err != nil {
			t.Errorf("SubcategoryService.Delete() error = %v\n", err)
		}
	})
}
