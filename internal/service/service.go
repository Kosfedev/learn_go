package service

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

type QuestionService interface {
	Create(ctx context.Context, newQuestion *model.NewQuestion) (int, error)
	Get(ctx context.Context, questionId int) (*model.Question, error)
	Update(ctx context.Context, updatedQuestion *model.UpdatedQuestion) error
	Delete(ctx context.Context, questionId int) error
}
