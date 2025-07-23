package service

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

type QuestionService interface {
	Create(ctx context.Context, newQuestion *model.NewQuestion) (int, error)
	Get(ctx context.Context, questionId int) (*model.Question, error)
	Update(ctx context.Context, questionId int, updatedQuestion *model.UpdatedQuestion) error
	Delete(ctx context.Context, questionId int) error
}

type QuestionSetService interface {
	Create(ctx context.Context, newQuestionSet *model.NewQuestionSet) (int, error)
	Get(ctx context.Context, questionSetId int) (*model.QuestionSet, error)
	Update(ctx context.Context, questionSetId int, updatedQuestionSet *model.UpdatedQuestionSet) error
	Delete(ctx context.Context, questionSetId int) error
}
