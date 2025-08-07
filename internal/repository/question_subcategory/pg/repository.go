package pg

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/repository"
	"github.com/Kosfedev/learn_go/internal/repository/question_subcategory/pg/converter"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/question_subcategory/pg/model"
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

func (r *repo) ListSubcategoryIDsByQuestionID(ctx context.Context, questionID int64) ([]int64, error) {
	builderSelect := sq.Select(columnSubcategoryID).
		From(tableQuestionSubcategory).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{columnQuestionID: questionID})

	queryRaw, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}

	query := db.Query{
		Name:     "question_subcategory_repository.list_subcategory_ids_by_question_id",
		QueryRaw: queryRaw,
	}

	questionSubcategories := make([]*modelRepo.QuestionSubcategory, 0)
	err = r.db.DB().ScanAll(ctx, &questionSubcategories, query, args...)
	if err != nil {
		return nil, err
	}

	subcategoryIDs := converter.QuestionSubcategoriesIDsFromPGSQL(questionSubcategories)

	return subcategoryIDs, nil
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
