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
	AddOptions(ctx context.Context, id int64, options []*model.NewQuestionOption) error
	DeleteOptions(ctx context.Context, ids []int64) error
}

type SetRepository interface {
	Create(ctx context.Context, newSet *model.NewSet) (int64, error)
	Get(ctx context.Context, id int64) (*model.Set, error)
	Update(ctx context.Context, id int64, updatedSet *model.UpdatedSet) error
	Delete(ctx context.Context, id int64) error
}

type QuestionSetRepository interface {
	AddQuestionsToSet(ctx context.Context, setID int64, questionIDs []int64) error
	RemoveQuestionsFromSet(ctx context.Context, setID int64, questionIDs []int64) error
	ListQuestionsBySetID(ctx context.Context, setID int64) ([]int64, error)
	AddQuestionToSets(ctx context.Context, questionID int64, setIDs []int64) error
	RemoveQuestionFromSets(ctx context.Context, questionID int64, setIDs []int64) error
	ListSetsByQuestionID(ctx context.Context, questionID int64) ([]int64, error)
}

type QuestionSubcategoryRepository interface {
	AddSubcategoriesToQuestion(ctx context.Context, questionID int64, subcategoryIDs []int64) error
	RemoveSubcategoriesFromQuestion(ctx context.Context, questionID int64, subcategoryIDs []int64) error
	ListSubcategoriesByQuestionID(ctx context.Context, questionID int64) ([]int64, error)
	// TODO: AddQuestionsToSubcategory(ctx context.Context, subcategoryID int64, questionIDs []int64) error
	// TODO: RemoveQuestionsFromSubcategory(ctx context.Context, subcategoryID int64, questionIDs []int64) error
	// TODO: ListQuestionsBySubcategoryID(ctx context.Context, subcategoryID int64) ([]int64, error)
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
