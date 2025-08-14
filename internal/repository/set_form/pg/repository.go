package pg

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository"
	"github.com/Kosfedev/learn_go/internal/repository/set_form/pg/converter"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/set_form/pg/model"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.SetFormRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) GetWithQuestions(ctx context.Context, setID int64) (*model.SetForm, error) {
	setFormRepo := &modelRepo.SetForm{}
	queryRaw := `SELECT set.id AS "set.id", set.created_at AS "set.created_at", set.updated_at AS "set.updated_at",
		 (SELECT json_agg(json_build_object(
      'id', q.id,
      'text', q.text,
      'type', q.type,
      'reference_answer', q.reference_answer,
      'created_at', q.created_at,
      'updated_at', q.updated_at
    )) FROM question AS q JOIN question_set AS q_set ON q_set.question_id = q.id WHERE q_set.set_id = set.id) AS questions
		FROM set
		WHERE set.id = $1`

	query := db.Query{
		Name:     "set_form.get_with_questions",
		QueryRaw: queryRaw,
	}

	err := r.db.DB().ScanOne(ctx, setFormRepo, query, setID)
	if err != nil {
		return nil, err
	}

	setForm := converter.SetFormFromPGSQL(setFormRepo)

	return setForm, nil
}
