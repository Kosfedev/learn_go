package pg

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository"
)

type repo struct {
}

func NewRepository() repository.QuestionRepository {
	return &repo{}
}

func (r *repo) Create(ctx context.Context, question *model.NewQuestion) (int, error) {
	return 0, nil
}

func (r *repo) Get(ctx context.Context, id int) (*model.Question, error) {
	return nil, nil
}

func (r *repo) Update(ctx context.Context, id int, updatedQuestion *model.UpdatedQuestion) error {
	return nil
}

func (r *repo) Delete(ctx context.Context, id int) error {
	return nil
}

func (r *repo) AddOptions(ctx context.Context, questionId int, options []*model.NewQuestionOption) error {
	return nil
}

func (r *repo) DeleteOptions(ctx context.Context, ids int) error {
	return nil
}
