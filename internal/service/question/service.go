package question

import (
	"github.com/Kosfedev/learn_go/internal/repository"
	"github.com/Kosfedev/learn_go/internal/service"
)

// TODO: вынести в агрегирующий сервис?
type serv struct {
	questionRepo            repository.QuestionRepository
	questionSetRepo         repository.QuestionSetRepository
	questionSubcategoryRepo repository.QuestionSubcategoryRepository
	setRepo                 repository.SetRepository
	subcategoryRepo         repository.SubcategoryRepository
}

func NewService(
	questionRepo repository.QuestionRepository,
	questionSetRepo repository.QuestionSetRepository,
	questionSubcategoryRepo repository.QuestionSubcategoryRepository,
	setRepo repository.SetRepository,
	subcategoryRepo repository.SubcategoryRepository,
) service.QuestionService {
	return &serv{
		questionRepo:            questionRepo,
		questionSetRepo:         questionSetRepo,
		questionSubcategoryRepo: questionSubcategoryRepo,
		setRepo:                 setRepo,
		subcategoryRepo:         subcategoryRepo,
	}
}
