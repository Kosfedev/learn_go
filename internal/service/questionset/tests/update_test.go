package tests

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/mocks"
	"github.com/Kosfedev/learn_go/internal/service/questionset"
)

func TestUpdate(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	id := int64(gofakeit.Int64())
	name := gofakeit.Question()
	req := &model.UpdatedQuestionSet{
		Name: &name,
	}
	mc := minimock.NewController(t)
	mockRepo := mocks.NewQuestionSetRepositoryMock(mc)
	mockRepo.UpdateMock.Expect(ctx, id, req).Return(nil)
	questionSetService := questionset.NewService(mockRepo)

	t.Run("Update placeholder implementation test", func(t *testing.T) {
		err := questionSetService.Update(context.Background(), id, req)
		if err != nil {
			t.Errorf("QuestionSetService.Update() error = %v\n", err)
		}
	})
}
