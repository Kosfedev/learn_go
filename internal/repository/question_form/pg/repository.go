package pg

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository"
	"github.com/Kosfedev/learn_go/internal/repository/question_form/pg/converter"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/question_form/pg/model"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.QuestionFormRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) GetWithOptionsSetsSubcategories(ctx context.Context, questionID int64) (*model.QuestionForm, error) {
	questionFormRepo := &modelRepo.QuestionForm{}
	queryRaw := `SELECT q.id AS "question.id", q.text AS "question.text", q.type AS "question.type", q.reference_answer AS "question.reference_answer", q.created_at AS "question.created_at", q.updated_at AS "question.updated_at", 
		 (SELECT json_agg(qo) FROM question_option AS qo WHERE qo.question_id = q.id) AS question_options,
		 (SELECT json_agg(json_build_object(
      'id', s.id,
      'name', s.name,
      'created_at', s.created_at,
      'updated_at', s.updated_at
    )) FROM set AS s JOIN question_set AS q_set ON q_set.question_id = q.id WHERE q_set.set_id = s.id) AS sets,
		 (SELECT json_agg(json_build_object(
      'id', sub.id,
      'name', sub.name,
      'category_id', sub.category_id,
      'created_at', sub.created_at,
      'updated_at', sub.updated_at
    )) FROM subcategory AS sub JOIN question_subcategory AS q_sub ON q_sub.question_id = q.id WHERE q_sub.subcategory_id = sub.id) AS subcategories
		FROM question AS q
		WHERE q.id = $1`

	query := db.Query{
		Name:     "question_form.get_with_options_sets_subcategories",
		QueryRaw: queryRaw,
	}

	err := r.db.DB().ScanOne(ctx, questionFormRepo, query, questionID)
	if err != nil {
		return nil, err
	}

	questionForm := converter.QuestionFormFromPGSQL(questionFormRepo)

	return questionForm, nil
}

func (r *repo) GetWithOptions(ctx context.Context, questionID int64) (*model.QuestionWithOptions, error) {
	questionWithOptionsRepo := &modelRepo.QuestionWithOptions{}
	queryRaw := `SELECT q.id AS "question.id", q.text AS "question.text", q.type AS "question.type", q.reference_answer AS "question.reference_answer", q.created_at AS "question.created_at", q.updated_at AS "question.updated_at", 
		 (SELECT json_agg(qo) FROM question_option AS qo WHERE qo.question_id = q.id) AS question_options
		FROM question AS q
		WHERE q.id = $1`

	query := db.Query{
		Name:     "question_form.get_with_options",
		QueryRaw: queryRaw,
	}

	err := r.db.DB().ScanOne(ctx, questionWithOptionsRepo, query, questionID)
	if err != nil {
		return nil, err
	}

	questionForm := converter.QuestionWithOptionsFromPGSQL(questionWithOptionsRepo)

	return questionForm, nil
}
