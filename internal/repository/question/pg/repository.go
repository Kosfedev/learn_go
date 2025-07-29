package pg

import (
	"context"
	"database/sql"
	"fmt"
	"log"

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

func (r *repo) Create(ctx context.Context, newQuestion *model.NewQuestion) (int, error) {
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

	if len(newQuestion.Options) > 0 {
		var values []interface{}
		for i, option := range newQuestion.Options {
			// TODO: вынести вне цикла? может ругаться на диалект
			if i == 0 {
				query = `
				INSERT INTO question_option (question_id, text, is_correct) 
				VALUES ($1, $2, $3)`
			} else {
				nArgs := 3
				query += fmt.Sprintf(", ($%d, $%d, $%d)", i*nArgs+1, i*nArgs+2, i*nArgs+3)
			}

			newOptionRepo := converter.NewQuestionOptionToPGSQL(questionId, option)
			values = append(values, newOptionRepo.QuestionId, newOptionRepo.Text, newOptionRepo.IsCorrect)
		}

		log.Printf("%s\n", query)
		log.Printf("%+v\n", values)
		_, err = r.db.Query(query, values...)
		if err != nil {
			return 0, err
		}
	}

	return questionId, nil
}

func (r *repo) Get(ctx context.Context, id int) (*model.Question, error) {
	questionRepo := &modelRepo.Question{}
	query := `SELECT * FROM question WHERE id = $1`
	row := r.db.QueryRow(query, id)
	err := row.Scan(&questionRepo.Id, &questionRepo.Text, &questionRepo.Type, &questionRepo.ReferenceAnswer, &questionRepo.CreatedAt, &questionRepo.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
	}

	questionServ := converter.QuestionFromPGSQL(questionRepo)

	questionOptionsServ := make([]*model.QuestionOption, 0)
	query = `SELECT * FROM question_option WHERE question_id = $1`
	rows, err := r.db.Query(query, id)

	for rows.Next() {
		questionOptionRepo := &modelRepo.QuestionOption{}
		err = rows.Scan(&questionOptionRepo.Id, &questionOptionRepo.QuestionId, &questionOptionRepo.Text, &questionOptionRepo.IsCorrect)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			log.Println(err)
		}
		questionOptionsServ = append(questionOptionsServ, converter.QuestionOptionFromPGSQL(questionOptionRepo))
	}

	questionServ.Options = questionOptionsServ

	return questionServ, nil
}

func (r *repo) Update(ctx context.Context, id int, updatedQuestion *model.UpdatedQuestion) error {
	return nil
}

func (r *repo) Delete(ctx context.Context, id int) error {
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
