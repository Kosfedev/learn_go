package questionform

import (
	"github.com/Kosfedev/learn_go/internal/repository"
	"github.com/Kosfedev/learn_go/internal/service"
)

type serv struct {
	repo repository.QuestionFormRepository
}

func NewService(repo repository.QuestionFormRepository) service.QuestionFormService {
	return &serv{
		repo: repo,
	}
}
