package pg

import (
	"database/sql"

	"github.com/Kosfedev/learn_go/internal/repository"
)

type repo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) repository.QuestionSetRepository {
	return &repo{
		db: db,
	}
}
