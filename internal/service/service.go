package service

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

type QuestionService interface {
	Create(ctx context.Context, newQuestion *model.NewQuestion) (int, error)
	Get(ctx context.Context, id int) (*model.Question, error)
	Update(ctx context.Context, id int, updatedQuestion *model.UpdatedQuestion) error
	Delete(ctx context.Context, id int) error
	AddOptions(ctx context.Context, questionId int, newQuestionOptions []*model.NewQuestionOption) error
	DeleteOptions(ctx context.Context, ids []int) error
	AddSubcategories(ctx context.Context, questionId int, subcategoryIds []int) error
	RemoveSubcategories(ctx context.Context, questionId int, subcategoryIds []int) error
}

type QuestionSetService interface {
	Create(ctx context.Context, newQuestionSet *model.NewQuestionSet) (int, error)
	Get(ctx context.Context, id int) (*model.QuestionSet, error)
	Update(ctx context.Context, id int, updatedQuestionSet *model.UpdatedQuestionSet) error
	Delete(ctx context.Context, id int) error
}

type DomainService interface {
	Create(ctx context.Context, newDomain *model.NewDomain) (int, error)
	Get(ctx context.Context, id int) (*model.Domain, error)
	Update(ctx context.Context, id int, updatedDomain *model.UpdatedDomain) error
	Delete(ctx context.Context, id int) error
}

type CategoryService interface {
	Create(ctx context.Context, newCategory *model.NewCategory) (int, error)
	Get(ctx context.Context, id int) (*model.Category, error)
	Update(ctx context.Context, id int, updatedCategory *model.UpdatedCategory) error
	Delete(ctx context.Context, id int) error
}

type SubcategoryService interface {
	Create(ctx context.Context, newSubcategory *model.NewSubcategory) (int, error)
	Get(ctx context.Context, id int) (*model.Subcategory, error)
	Update(ctx context.Context, id int, updatedSubcategory *model.UpdatedSubcategory) error
	Delete(ctx context.Context, id int) error
}
