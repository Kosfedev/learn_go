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

type QuestionsDomainService interface {
	Create(ctx context.Context, newDomain *model.NewQuestionsDomain) (int, error)
	Get(ctx context.Context, domainId int) (*model.QuestionsDomain, error)
	Update(ctx context.Context, domainId int, updatedDomain *model.UpdatedQuestionsDomain) error
	Delete(ctx context.Context, domainId int) error
}

type QuestionsCategoryService interface {
	Create(ctx context.Context, newCategory *model.NewQuestionsCategory) (int, error)
	Get(ctx context.Context, categoryId int) (*model.QuestionsCategory, error)
	Update(ctx context.Context, categoryId int, updatedCategory *model.UpdatedQuestionsCategory) error
	Delete(ctx context.Context, categoryId int) error
}

type QuestionsSubcategoryService interface {
	Create(ctx context.Context, newSubcategory *model.NewQuestionsSubcategory) (int, error)
	Get(ctx context.Context, subcategoryId int) (*model.QuestionsSubcategory, error)
	Update(ctx context.Context, subcategoryId int, updatedSubcategory *model.UpdatedQuestionsSubcategory) error
	Delete(ctx context.Context, subcategoryId int) error
}
