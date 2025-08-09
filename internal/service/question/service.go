package question

import (
	"github.com/Kosfedev/learn_go/internal/repository"
	"github.com/Kosfedev/learn_go/internal/service"
)

type serv struct {
	questionRepo       repository.QuestionRepository
	questionOptionRepo repository.QuestionOptionRepository
}

func NewService(
	questionRepo repository.QuestionRepository,
	questionOptionRepo repository.QuestionOptionRepository,
) service.QuestionService {
	return &serv{
		questionRepo:       questionRepo,
		questionOptionRepo: questionOptionRepo,
	}
}
