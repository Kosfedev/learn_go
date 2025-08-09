package pg

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/repository"
)

const (
	tableQuestionSet = "question_set"
	columnQuestionID = "question_id"
	columnSetID      = "set_id"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.QuestionSetRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) AddQuestionToSets(ctx context.Context, questionID int64, setIDs []int64) error {
	if len(setIDs) == 0 {
		return nil
	}

	builderInsert := sq.Insert(tableQuestionSet).
		PlaceholderFormat(sq.Dollar).
		Columns(columnQuestionID, columnSetID)

	for _, setID := range setIDs {
		builderInsert = builderInsert.Values(questionID, setID)
	}

	queryRaw, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}

	query := db.Query{
		Name:     "question_set_repository.add_question_to_sets",
		QueryRaw: queryRaw,
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) AddQuestionsToSet(ctx context.Context, setID int64, questionIDs []int64) error {
	if len(questionIDs) == 0 {
		return nil
	}

	builderInsert := sq.Insert(tableQuestionSet).
		PlaceholderFormat(sq.Dollar).
		Columns(columnQuestionID, columnSetID)

	for _, questionID := range questionIDs {
		builderInsert = builderInsert.Values(questionID, setID)
	}

	queryRaw, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}

	query := db.Query{
		Name:     "question_set_repository.add_questions_to_set",
		QueryRaw: queryRaw,
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) RemoveQuestionFromSets(ctx context.Context, questionID int64, setIDs []int64) error {
	if len(setIDs) == 0 {
		return nil
	}

	builderDelete := sq.Delete(tableQuestionSet).
		PlaceholderFormat(sq.Dollar).
		Where(sq.And{sq.Eq{columnQuestionID: questionID}, sq.Eq{columnSetID: setIDs}})

	queryRaw, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}

	query := db.Query{
		Name:     "question_set_repository.remove_question_from_sets",
		QueryRaw: queryRaw,
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) RemoveQuestionsFromSet(ctx context.Context, setID int64, questionIDs []int64) error {
	if len(questionIDs) == 0 {
		return nil
	}

	builderDelete := sq.Delete(tableQuestionSet).
		PlaceholderFormat(sq.Dollar).
		Where(sq.And{sq.Eq{columnSetID: setID}, sq.Eq{columnQuestionID: questionIDs}})

	queryRaw, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}

	query := db.Query{
		Name:     "question_set_repository.remove_questions_from_set",
		QueryRaw: queryRaw,
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
