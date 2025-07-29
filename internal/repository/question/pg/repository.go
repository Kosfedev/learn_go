package pg

import (
	"context"
	"database/sql"
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

func (r *repo) Create(ctx context.Context, question *model.NewQuestion) (int, error) {
	return 0, nil
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
	return nil
}

func (r *repo) AddOptions(ctx context.Context, questionId int, options []*model.NewQuestionOption) error {
	return nil
}

func (r *repo) DeleteOptions(ctx context.Context, ids int) error {
	return nil
}
