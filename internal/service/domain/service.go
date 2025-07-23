package domain

import (
	"github.com/Kosfedev/learn_go/internal/service"
)

type serv struct{}

func NewService() service.QuestionsDomainService {
	return &serv{}
}
