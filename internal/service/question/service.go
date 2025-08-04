package question

import (
	"github.com/Kosfedev/learn_go/internal/repository"
	"github.com/Kosfedev/learn_go/internal/service"
)

type serv struct {
	questionRepo    repository.QuestionRepository
	subcategoryRepo repository.SubcategoryRepository
}

func NewService(questionRepo repository.QuestionRepository, subcategoryRepo repository.SubcategoryRepository) service.QuestionService {
	return &serv{
		questionRepo:    questionRepo,
		subcategoryRepo: subcategoryRepo,
	}
}
