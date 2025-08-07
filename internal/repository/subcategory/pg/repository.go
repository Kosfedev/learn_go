package pg

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository"
	"github.com/Kosfedev/learn_go/internal/repository/subcategory/pg/converter"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/subcategory/pg/model"
)

const (
	tableSubcategory = "subcategory"
	columnID         = "id"
	columnName       = "name"
	columnCategoryID = "category_id"
	columnCreatedAt  = "created_at"
	columnUpdatedAt  = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.SubcategoryRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, subcategory *model.NewSubcategory) (int64, error) {
	subcategoryRepo := converter.NewSubcategoryToPGSQL(subcategory)
	builderInsert := sq.Insert(tableSubcategory).
		PlaceholderFormat(sq.Dollar).
		Columns(columnName, columnCategoryID).
		Values(subcategoryRepo.Name, subcategoryRepo.CategoryID).
		Suffix(fmt.Sprintf("RETURNING %v", columnID))

	queryRaw, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	query := db.Query{
		Name:     "subcategory_repository.create",
		QueryRaw: queryRaw,
	}

	var id int64
	err = r.db.DB().ScanOne(ctx, &id, query, args...)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.Subcategory, error) {
	builderSelect := sq.Select(columnID, columnName, columnCategoryID, columnCreatedAt, columnUpdatedAt).
		From(tableSubcategory).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{columnID: id})

	queryRaw, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}

	query := db.Query{
		Name:     "subcategory_repository.get",
		QueryRaw: queryRaw,
	}

	subcategoryRepo := &modelRepo.Subcategory{}
	err = r.db.DB().ScanOne(ctx, subcategoryRepo, query, args...)
	if err != nil {
		return nil, err
	}

	subcategory := converter.SubcategoryFromPGSQL(subcategoryRepo)

	return subcategory, nil
}

func (r *repo) Update(ctx context.Context, id int64, subcategory *model.UpdatedSubcategory) error {
	subcategoryRepo := converter.UpdatedSubcategoryToPGSQL(subcategory)
	builderUpdate := sq.Update(tableSubcategory).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{columnID: id})

	if subcategoryRepo.Name.Valid {
		builderUpdate = builderUpdate.Set(columnName, subcategoryRepo.Name)
	}

	if subcategoryRepo.CategoryID.Valid {
		builderUpdate = builderUpdate.Set(columnCategoryID, subcategoryRepo.CategoryID)
	}

	queryRaw, args, err := builderUpdate.ToSql()
	if err != nil {
		return err
	}

	query := db.Query{
		Name:     "subcategory_repository.update",
		QueryRaw: queryRaw,
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builderDelete := sq.Delete(tableSubcategory).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{columnID: id})

	queryRaw, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}

	query := db.Query{
		Name:     "subcategory_repository.delete",
		QueryRaw: queryRaw,
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) ListByIDs(ctx context.Context, ids []int64) ([]*model.Subcategory, error) {
	builderSelect := sq.Select(columnID, columnName, columnCategoryID, columnCreatedAt, columnUpdatedAt).
		From(tableSubcategory).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{columnID: ids})

	queryRaw, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}

	query := db.Query{
		Name:     "subcategory_repository.list_by_ids",
		QueryRaw: queryRaw,
	}

	subcategoriesRepo := make([]*modelRepo.Subcategory, len(ids))
	err = r.db.DB().ScanAll(ctx, &subcategoriesRepo, query, args...)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n", *subcategoriesRepo[1])
	subcategory := converter.SubcategoriesFromPGSQL(subcategoriesRepo)

	return subcategory, nil
}
