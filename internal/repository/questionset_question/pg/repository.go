package pg

import (
	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/repository"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.QuestionQuestionSetRepository {
	return &repo{
		db: db,
	}
}
