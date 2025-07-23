package Subategory

import (
	"github.com/Kosfedev/learn_go/internal/service"
)

type serv struct{}

func NewService() service.QuestionsSubcategoryService {
	return &serv{}
}
