package pg

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository"
	"github.com/Kosfedev/learn_go/internal/repository/question/pg/converter"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/question/pg/model"
)

type repo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) repository.QuestionRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(_ context.Context, newQuestion *model.NewQuestion) (int, error) {
	newQuestionRepo := converter.NewQuestionToPGSQL(newQuestion)
	query := `
		INSERT INTO question (text, type, reference_answer) 
		VALUES ($1, $2, $3)
		RETURNING id;`

	var questionId int
	err := r.db.QueryRow(query, newQuestionRepo.Text, newQuestionRepo.Type, newQuestionRepo.ReferenceAnswer).Scan(&questionId)
	if err != nil {
		return 0, err
	}

	nArgs := 3
	query = `INSERT INTO question_option (question_id, text, is_correct) VALUES`
	// TODO: нужна транзакция
	if len(newQuestion.Options) > 0 {
		var values []interface{}

		for i, option := range newQuestion.Options {
			query += fmt.Sprintf(" ($%d, $%d, $%d),", i*nArgs+1, i*nArgs+2, i*nArgs+3)
			newOptionRepo := converter.NewQuestionOptionToPGSQL(questionId, option)
			values = append(values, newOptionRepo.QuestionId, newOptionRepo.Text, newOptionRepo.IsCorrect)
		}

		query = strings.TrimSuffix(query, ",")

		_, err = r.db.Query(query, values...)
		if err != nil {
			return 0, err
		}
	}

	return questionId, nil
}

func (r *repo) Get(_ context.Context, id int) (*model.Question, error) {
	questionRepo := &modelRepo.Question{}
	query := `SELECT * FROM question WHERE id = $1`
	row := r.db.QueryRow(query, id)
	err := row.Scan(&questionRepo.Id, &questionRepo.Text, &questionRepo.Type, &questionRepo.ReferenceAnswer, &questionRepo.CreatedAt, &questionRepo.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	questionServ := converter.QuestionFromPGSQL(questionRepo)

	questionOptionsServ := make([]*model.QuestionOption, 0)
	query = `SELECT * FROM question_option WHERE question_id = $1`
	rows, err := r.db.Query(query, id)

	for rows.Next() {
		questionOptionRepo := &modelRepo.QuestionOption{}
		err = rows.Scan(&questionOptionRepo.Id, &questionOptionRepo.QuestionId, &questionOptionRepo.Text, &questionOptionRepo.IsCorrect)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}

			return nil, err
		}
		questionOptionsServ = append(questionOptionsServ, converter.QuestionOptionFromPGSQL(questionOptionRepo))
	}

	questionServ.Options = questionOptionsServ

	return questionServ, nil
}

func (r *repo) Update(_ context.Context, id int, updatedQuestion *model.UpdatedQuestion) error {
	updatedQuestionRepo := converter.UpdatedQuestionToPGSQL(updatedQuestion)
	index := 2
	values := []interface{}{id}
	query := "UPDATE question SET"

	if updatedQuestionRepo.Text.Valid {
		values = append(values, updatedQuestionRepo.Text)
		query += fmt.Sprintf(" text = $%d,", index)
		index++
	}

	if updatedQuestionRepo.Type.Valid {
		values = append(values, updatedQuestionRepo.Type)
		query += fmt.Sprintf(" type = $%d,", index)
		index++
	}

	if updatedQuestionRepo.ReferenceAnswer.Valid {
		values = append(values, updatedQuestionRepo.ReferenceAnswer)
		query += fmt.Sprintf(" reference_answer = $%d,", index)
		index++
	}

	query = strings.TrimSuffix(query, ",")
	query += " WHERE id = $1"

	_, err := r.db.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Delete(_ context.Context, id int) error {
	query := `DELETE FROM question WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) AddOptions(ctx context.Context, questionId int, options []*model.NewQuestionOption) error {
	return nil
}

func (r *repo) DeleteOptions(ctx context.Context, ids int) error {
	return nil
}
