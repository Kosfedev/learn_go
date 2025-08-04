package repository

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

type QuestionRepository interface {
	Create(ctx context.Context, question *model.NewQuestion) (int, error)
	Get(ctx context.Context, id int) (*model.Question, error)
	Update(ctx context.Context, id int, updatedQuestion *model.UpdatedQuestion) error
	Delete(ctx context.Context, id int) error
	AddOptions(ctx context.Context, questionId int, options []*model.NewQuestionOption) error
	DeleteOptions(ctx context.Context, ids []int) error
	AddSubcategories(ctx context.Context, questionId int, subcategoryIds []int) error
	RemoveSubcategories(ctx context.Context, questionId int, subcategoryIds []int) error
	ListSubcategoriesByQuestionId(ctx context.Context, questionId int) ([]int, error)
}

type QuestionSetRepository interface {
	Create(ctx context.Context, question *model.NewQuestionSet) (int, error)
	Get(ctx context.Context, id int) (*model.QuestionSet, error)
	Update(ctx context.Context, id int, updatedQuestion *model.UpdatedQuestionSet) error
	Delete(ctx context.Context, id int) error
	// AddQuestions(ctx context.Context, questionSetId int, questionIds []int) error
	// RemoveQuestions(ctx context.Context, ids []int) error
}

type DomainRepository interface {
	Create(ctx context.Context, domain *model.NewDomain) (int, error)
	Get(ctx context.Context, id int) (*model.Domain, error)
	Update(ctx context.Context, id int, updatedDomain *model.UpdatedDomain) error
	Delete(ctx context.Context, id int) error
}

type CategoryRepository interface {
	Create(ctx context.Context, category *model.NewCategory) (int, error)
	Get(ctx context.Context, id int) (*model.Category, error)
	Update(ctx context.Context, id int, updatedCategory *model.UpdatedCategory) error
	Delete(ctx context.Context, id int) error
}

type SubcategoryRepository interface {
	Create(ctx context.Context, subcategory *model.NewSubcategory) (int, error)
	Get(ctx context.Context, id int) (*model.Subcategory, error)
	Update(ctx context.Context, id int, updatedSubcategory *model.UpdatedSubcategory) error
	Delete(ctx context.Context, id int) error
	ListByIds(ctx context.Context, ids []int) ([]*model.Subcategory, error)
}
