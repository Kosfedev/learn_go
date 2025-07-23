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

type DomainService interface {
	Create(ctx context.Context, newDomain *model.NewDomain) (int, error)
	Get(ctx context.Context, domainId int) (*model.Domain, error)
	Update(ctx context.Context, domainId int, updatedDomain *model.UpdatedDomain) error
	Delete(ctx context.Context, domainId int) error
}

type CategoryService interface {
	Create(ctx context.Context, newCategory *model.NewCategory) (int, error)
	Get(ctx context.Context, categoryId int) (*model.Category, error)
	Update(ctx context.Context, categoryId int, updatedCategory *model.UpdatedCategory) error
	Delete(ctx context.Context, categoryId int) error
}

type SubcategoryService interface {
	Create(ctx context.Context, newSubcategory *model.NewSubcategory) (int, error)
	Get(ctx context.Context, subcategoryId int) (*model.Subcategory, error)
	Update(ctx context.Context, subcategoryId int, updatedSubcategory *model.UpdatedSubcategory) error
	Delete(ctx context.Context, subcategoryId int) error
}
