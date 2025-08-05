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

func TestUpdate(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	id := gofakeit.Int64()
	text := gofakeit.Sentence(5)
	req := &model.UpdatedSubcategory{
		Name: &text,
	}
	mc := minimock.NewController(t)
	mockRepo := mocks.NewSubcategoryRepositoryMock(mc)
	mockRepo.UpdateMock.Expect(ctx, id, req).Return(nil)
	subcategoryService := subcategory.NewService(mockRepo)

	t.Run("Update placeholder implementation test", func(t *testing.T) {
		err := subcategoryService.Update(ctx, id, req)
		if err != nil {
			t.Errorf("SubcategoryService.Update() error = %v\n", err)
		}
	})
}
