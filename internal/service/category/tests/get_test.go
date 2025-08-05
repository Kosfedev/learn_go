package tests

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/mocks"
	"github.com/Kosfedev/learn_go/internal/service/category"
)

func TestGet(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	id := int(gofakeit.Int64())
	domainID := int(gofakeit.Int64())
	res := &model.Category{
		ID:        id,
		Name:      gofakeit.Name(),
		DomainID:  domainID,
		CreatedAt: time.Time{},
	}
	mc := minimock.NewController(t)
	mockRepo := mocks.NewCategoryRepositoryMock(mc)
	mockRepo.GetMock.Expect(ctx, id).Return(res, nil)
	categoryService := category.NewService(mockRepo)

	t.Run("Get placeholder implementation test", func(t *testing.T) {
		_, err := categoryService.Get(ctx, id)
		if err != nil {
			t.Errorf("CategoryService.Get() error = %v\n", err)
		}
	})
}
