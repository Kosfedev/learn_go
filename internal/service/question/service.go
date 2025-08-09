package question

import (
	"github.com/Kosfedev/learn_go/internal/repository"
	"github.com/Kosfedev/learn_go/internal/service"
)

// TODO: вынести в агрегирующий сервис?
type serv struct {
	questionRepo            repository.QuestionRepository
	questionOptionRepo      repository.QuestionOptionRepository
	questionSubcategoryRepo repository.QuestionSubcategoryRepository
	subcategoryRepo         repository.SubcategoryRepository
}

func NewService(
	questionRepo repository.QuestionRepository,
	questionOptionRepo repository.QuestionOptionRepository,
	questionSubcategoryRepo repository.QuestionSubcategoryRepository,
	subcategoryRepo repository.SubcategoryRepository,
) service.QuestionService {
	return &serv{
		questionRepo:            questionRepo,
		questionOptionRepo:      questionOptionRepo,
		questionSubcategoryRepo: questionSubcategoryRepo,
		subcategoryRepo:         subcategoryRepo,
	}
}
