package repository

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

type QuestionRepository interface {
	Create(ctx context.Context, question *model.NewQuestion) (int64, error)
	Get(ctx context.Context, id int64) (*model.Question, error)
	Update(ctx context.Context, id int64, updatedQuestion *model.UpdatedQuestion) error
	Delete(ctx context.Context, id int64) error
	AddOptions(ctx context.Context, questionID int64, options []*model.NewQuestionOption) error
	DeleteOptions(ctx context.Context, ids []int64) error
}

type QuestionSetRepository interface {
	Create(ctx context.Context, question *model.NewQuestionSet) (int64, error)
	Get(ctx context.Context, id int64) (*model.QuestionSet, error)
	Update(ctx context.Context, id int64, updatedQuestion *model.UpdatedQuestionSet) error
	Delete(ctx context.Context, id int64) error
	// AddQuestions(ctx context.Context, questionSetID int64, questionIDs []int64) error
	// RemoveQuestions(ctx context.Context, ids []int64) error
}

type QuestionSubcategoryRepository interface {
	AddSubcategoriesToQuestion(ctx context.Context, questionID int64, subcategoryIDs []int64) error
	RemoveSubcategoriesFromQuestion(ctx context.Context, questionID int64, subcategoryIDs []int64) error
	ListSubcategoriesByQuestionID(ctx context.Context, questionID int64) ([]int64, error)
	// AddQuestionsToSubcategory(ctx context.Context, subcategoryID int64, questionIDs []int64) error
	// RemoveQuestionsFromSubcategory(ctx context.Context, subcategoryID int64, questionIDs []int64) error
	// ListQuestionsBySubcategoryID(ctx context.Context, subcategoryID int64) ([]int64, error)
}

type DomainRepository interface {
	Create(ctx context.Context, domain *model.NewDomain) (int64, error)
	Get(ctx context.Context, id int64) (*model.Domain, error)
	Update(ctx context.Context, id int64, updatedDomain *model.UpdatedDomain) error
	Delete(ctx context.Context, id int64) error
}

type CategoryRepository interface {
	Create(ctx context.Context, category *model.NewCategory) (int64, error)
	Get(ctx context.Context, id int64) (*model.Category, error)
	Update(ctx context.Context, id int64, updatedCategory *model.UpdatedCategory) error
	Delete(ctx context.Context, id int64) error
}

type SubcategoryRepository interface {
	Create(ctx context.Context, subcategory *model.NewSubcategory) (int64, error)
	Get(ctx context.Context, id int64) (*model.Subcategory, error)
	Update(ctx context.Context, id int64, updatedSubcategory *model.UpdatedSubcategory) error
	Delete(ctx context.Context, id int64) error
	ListByIDs(ctx context.Context, ids []int64) ([]*model.Subcategory, error)
}
