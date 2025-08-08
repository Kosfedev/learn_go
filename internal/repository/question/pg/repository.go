package pg

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository"
	"github.com/Kosfedev/learn_go/internal/repository/question/pg/converter"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/question/pg/model"
)

const (
	tableQuestion         = "question"
	tableQuestionOption   = "question_option"
	columnID              = "id"
	columnText            = "text"
	columnType            = "type"
	columnReferenceAnswer = "reference_answer"
	columnCreatedAt       = "created_at"
	columnUpdatedAt       = "updated_at"
	columnQuestionID      = "question_id"
	columnIsCorrect       = "is_correct"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.QuestionRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(ctx context.Context, newQuestion *model.NewQuestion) (int64, error) {
	newQuestionRepo := converter.NewQuestionToPGSQL(newQuestion)
	builderInsert := sq.Insert(tableQuestion).
		PlaceholderFormat(sq.Dollar).
		Columns(columnText, columnType, columnReferenceAnswer).
		Values(newQuestionRepo.Text, newQuestionRepo.Type, newQuestionRepo.ReferenceAnswer).
		Suffix(fmt.Sprintf("RETURNING %v", columnID))

	queryRaw, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	query := db.Query{
		Name:     "question_repository.create",
		QueryRaw: queryRaw,
	}

	var questionID int64
	err = r.db.DB().ScanOne(ctx, &questionID, query, args...)
	if err != nil {
		return 0, err
	}

	// TODO: нужна транзакция
	err = r.AddOptions(ctx, questionID, newQuestion.Options)
	if err != nil {
		return 0, err
	}

	return questionID, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.Question, error) {
	questionRepo := &modelRepo.Question{}
	builderSelect := sq.Select(columnID, columnText, columnType, columnReferenceAnswer, columnCreatedAt, columnUpdatedAt).
		From(tableQuestion).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{columnID: id})

	queryRaw, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}

	query := db.Query{
		Name:     "question_repository.get",
		QueryRaw: queryRaw,
	}

	err = r.db.DB().ScanOne(ctx, questionRepo, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	questionServ := converter.QuestionFromPGSQL(questionRepo)

	return questionServ, nil
}

func (r *repo) Update(ctx context.Context, id int64, updatedQuestion *model.UpdatedQuestion) error {
	updatedQuestionRepo := converter.UpdatedQuestionToPGSQL(updatedQuestion)
	builderUpdate := sq.Update(tableQuestion).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{columnID: id})

	if updatedQuestionRepo.Text.Valid {
		builderUpdate = builderUpdate.Set(columnText, updatedQuestionRepo.Text)
	}

	if updatedQuestionRepo.Type.Valid {
		builderUpdate = builderUpdate.Set(columnType, updatedQuestionRepo.Type)
	}

	if updatedQuestionRepo.ReferenceAnswer.Valid {
		builderUpdate = builderUpdate.Set(columnReferenceAnswer, updatedQuestionRepo.ReferenceAnswer)
	}

	queryRaw, args, err := builderUpdate.ToSql()
	if err != nil {
		return err
	}

	query := db.Query{
		Name:     "question_repository.update",
		QueryRaw: queryRaw,
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builderDelete := sq.Delete(tableQuestion).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{columnID: id})

	queryRaw, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}

	query := db.Query{
		Name:     "question_repository.delete",
		QueryRaw: queryRaw,
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) ListOptionsByQuestionID(ctx context.Context, questionID int64) ([]*model.QuestionOption, error) {
	questionOptionsRepo := make([]*modelRepo.QuestionOption, 0)
	builderSelect := sq.Select(columnID, columnQuestionID, columnText, columnIsCorrect).
		From(tableQuestionOption).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{columnQuestionID: questionID})

	queryRaw, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}

	query := db.Query{
		Name:     "question_repository.list_options",
		QueryRaw: queryRaw,
	}

	err = r.db.DB().ScanAll(ctx, &questionOptionsRepo, query, args...)
	if err != nil {
		return nil, err
	}

	questionOptions := converter.QuestionOptionsFromPGSQL(questionOptionsRepo)

	return questionOptions, nil
}

func (r *repo) AddOptions(ctx context.Context, questionID int64, options []*model.NewQuestionOption) error {
	if len(options) == 0 {
		return nil
	}

	builderInsert := sq.Insert(tableQuestionOption).
		PlaceholderFormat(sq.Dollar).
		Columns(columnQuestionID, columnText, columnIsCorrect)

	for _, option := range options {
		builderInsert = builderInsert.Values(questionID, option.Text, option.IsCorrect)
	}

	queryRaw, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}

	query := db.Query{
		Name:     "question_repository.add_options",
		QueryRaw: queryRaw,
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) DeleteOptions(ctx context.Context, ids []int64) error {
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
		Name:     "question_repository.delete_options",
		QueryRaw: queryRaw,
	}
	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
