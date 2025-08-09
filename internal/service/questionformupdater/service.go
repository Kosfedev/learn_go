package questionformupdater

import (
	"github.com/Kosfedev/learn_go/internal/repository"
	"github.com/Kosfedev/learn_go/internal/service"
)

type serv struct {
	questionRepo            repository.QuestionRepository
	questionOptionRepo      repository.QuestionOptionRepository
	questionSubcategoryRepo repository.QuestionSubcategoryRepository
}

func NewService(
	questionRepo repository.QuestionRepository,
	questionOptionRepo repository.QuestionOptionRepository,
	questionSubcategoryRepo repository.QuestionSubcategoryRepository,
) service.QuestionFormUpdaterService {
	return &serv{
		questionRepo:            questionRepo,
		questionOptionRepo:      questionOptionRepo,
		questionSubcategoryRepo: questionSubcategoryRepo,
	}
}
