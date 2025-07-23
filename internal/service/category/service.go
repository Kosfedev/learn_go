package Category

import (
	"github.com/Kosfedev/learn_go/internal/service"
)

type serv struct{}

func NewService() service.CategoryService {
	return &serv{}
}
