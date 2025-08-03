package questionset

import (
	"github.com/Kosfedev/learn_go/internal/repository"
	"github.com/Kosfedev/learn_go/internal/service"
)

type serv struct {
	repo repository.QuestionSetRepository
}

func NewService(repo repository.QuestionSetRepository) service.QuestionSetService {
	return &serv{repo: repo}
}
