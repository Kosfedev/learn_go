package question

import (
	"github.com/Kosfedev/learn_go/internal/repository"
	"github.com/Kosfedev/learn_go/internal/service"
)

type serv struct {
	questionRepo repository.QuestionRepository
}

func NewService(
	questionRepo repository.QuestionRepository,
) service.QuestionService {
	return &serv{
		questionRepo: questionRepo,
	}
}
