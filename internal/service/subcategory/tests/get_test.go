package tests

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/mocks"
	"github.com/Kosfedev/learn_go/internal/service/subcategory"
)

func TestGet(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	id := int(gofakeit.Int64())
	categoryId := int(gofakeit.Int64())
	res := &model.Subcategory{
		Id:         id,
		Name:       gofakeit.Name(),
		CategoryId: categoryId,
		CreatedAt:  time.Time{},
	}
	mc := minimock.NewController(t)
	mockRepo := mocks.NewSubcategoryRepositoryMock(mc)
	mockRepo.GetMock.Expect(ctx, id).Return(res, nil)
	subcategoryService := subcategory.NewService(mockRepo)

	t.Run("Get placeholder implementation test", func(t *testing.T) {
		_, err := subcategoryService.Get(ctx, id)
		if err != nil {
			t.Errorf("SubcategoryService.Get() error = %v\n", err)
		}
	})
}
