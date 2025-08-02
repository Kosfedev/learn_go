package question

import (
	"github.com/Kosfedev/learn_go/internal/repository"
	"github.com/Kosfedev/learn_go/internal/service"
)

type serv struct {
	repo repository.QuestionRepository
}

func NewService(repo repository.QuestionRepository) service.QuestionService {
	return &serv{
		repo: repo,
	}
}
