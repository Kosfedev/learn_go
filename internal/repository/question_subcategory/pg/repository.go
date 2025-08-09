package pg

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/repository"
)

const (
	tableQuestionSubcategory = "question_subcategory"
	columnQuestionID         = "question_id"
	columnSubcategoryID      = "subcategory_id"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.QuestionSubcategoryRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) AddSubcategoriesToQuestion(ctx context.Context, questionID int64, subcategoryIDs []int64) error {
	if len(subcategoryIDs) == 0 {
		return nil
	}

	builderInsert := sq.Insert(tableQuestionSubcategory).
		PlaceholderFormat(sq.Dollar).
		Columns(columnQuestionID, columnSubcategoryID)

	for _, subcategoryID := range subcategoryIDs {
		builderInsert = builderInsert.Values(questionID, subcategoryID)
	}

	queryRaw, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}

	query := db.Query{
		Name:     "question_subcategory_repository.add_subcategories_to_question",
		QueryRaw: queryRaw,
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) RemoveSubcategoriesFromQuestion(ctx context.Context, questionID int64, subcategoryIDs []int64) error {
	if len(subcategoryIDs) == 0 {
		return nil
	}

	builderDelete := sq.Delete(tableQuestionSubcategory).
		PlaceholderFormat(sq.Dollar).
		Where(sq.And{sq.Eq{columnQuestionID: questionID}, sq.Eq{columnSubcategoryID: subcategoryIDs}})

	queryRaw, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}

	query := db.Query{
		Name:     "question_subcategory_repository.remove_subcategories_from_question",
		QueryRaw: queryRaw,
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
