package pg

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository"
	"github.com/Kosfedev/learn_go/internal/repository/question_option/pg/converter"
)

type repo struct {
	db db.Client
}

const (
	tableQuestionOption = "question_option"
	columnID            = "id"
	columnText          = "text"
	columnQuestionID    = "question_id"
	columnIsCorrect     = "is_correct"
)

func NewRepository(db db.Client) repository.QuestionOptionRepository {
	return &repo{db: db}
}

func (r *repo) CreateList(ctx context.Context, questionID int64, options []*model.NewQuestionOption) error {
	if len(options) == 0 {
		return nil
	}

	optionsRepo := converter.NewQuestionOptionsToPGSQL(questionID, options)

	builderInsert := sq.Insert(tableQuestionOption).
		PlaceholderFormat(sq.Dollar).
		Columns(columnQuestionID, columnText, columnIsCorrect)

	for _, optionRepo := range optionsRepo {
		builderInsert = builderInsert.Values(questionID, optionRepo.Text, optionRepo.IsCorrect)
	}

	queryRaw, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}

	query := db.Query{
		Name:     "question_option_repository.add_options",
		QueryRaw: queryRaw,
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) DeleteList(ctx context.Context, ids []int64) error {
	if len(ids) == 0 {
		return nil
	}

	builderDelete := sq.Delete(tableQuestionOption).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{columnID: ids})

	queryRaw, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}

	query := db.Query{
		Name:     "question_option_repository.delete_options",
		QueryRaw: queryRaw,
	}
	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
