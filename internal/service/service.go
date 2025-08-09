package service

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

type QuestionService interface {
	Create(ctx context.Context, newQuestion *model.NewQuestion) (int64, error)
	Get(ctx context.Context, id int64) (*model.Question, error)
	Update(ctx context.Context, id int64, updatedQuestion *model.UpdatedQuestion) error
	Delete(ctx context.Context, id int64) error
}

type SetService interface {
	Create(ctx context.Context, newSet *model.NewSet) (int64, error)
	Get(ctx context.Context, id int64) (*model.Set, error)
	Update(ctx context.Context, id int64, updatedSet *model.UpdatedSet) error
	Delete(ctx context.Context, id int64) error
}

type DomainService interface {
	Create(ctx context.Context, newDomain *model.NewDomain) (int64, error)
	Get(ctx context.Context, id int64) (*model.Domain, error)
	Update(ctx context.Context, id int64, updatedDomain *model.UpdatedDomain) error
	Delete(ctx context.Context, id int64) error
}

type CategoryService interface {
	Create(ctx context.Context, newCategory *model.NewCategory) (int64, error)
	Get(ctx context.Context, id int64) (*model.Category, error)
	Update(ctx context.Context, id int64, updatedCategory *model.UpdatedCategory) error
	Delete(ctx context.Context, id int64) error
}

type SubcategoryService interface {
	Create(ctx context.Context, newSubcategory *model.NewSubcategory) (int64, error)
	Get(ctx context.Context, id int64) (*model.Subcategory, error)
	Update(ctx context.Context, id int64, updatedSubcategory *model.UpdatedSubcategory) error
	Delete(ctx context.Context, id int64) error
}

type QuestionFormService interface {
	GetWithOptionsSetsSubcategories(ctx context.Context, questionID int64) (*model.QuestionForm, error)
}

type QuestionFormUpdaterService interface {
	CreateWithOptions(ctx context.Context, newQuestion *model.NewQuestionWithOptions) (int64, error)
	CreateOptions(ctx context.Context, questionID int64, newQuestionOptions []*model.NewQuestionOption) error
	DeleteOptions(ctx context.Context, optionIDs []int64) error
	AddSubcategories(ctx context.Context, questionID int64, subcategoryIDs []int64) error
	RemoveSubcategories(ctx context.Context, questionID int64, subcategoryIDs []int64) error
	AddSets(ctx context.Context, questionID int64, setIDs []int64) error
	RemoveSets(ctx context.Context, questionID int64, setIDs []int64) error
}
