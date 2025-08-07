package pg

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository"
	"github.com/Kosfedev/learn_go/internal/repository/category/pg/converter"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/category/pg/model"
)

const (
	tableCategory  = "category"
	columnID       = "id"
	columnName     = "name"
	columnDomainID = "domain_id"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.CategoryRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, category *model.NewCategory) (int64, error) {
	categoryRepo := converter.NewCategoryToPGSQL(category)
	builderInsert := sq.Insert(tableCategory).
		PlaceholderFormat(sq.Dollar).
		Columns(columnName, columnDomainID).
		Values(categoryRepo.Name, category.DomainID).
		Suffix(fmt.Sprintf("RETURNING %v", columnID))

	queryRaw, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	query := db.Query{
		Name:     "category_repository.create",
		QueryRaw: queryRaw,
	}

	var id int64
	err = r.db.DB().ScanOne(ctx, &id, query, args...)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.Category, error) {
	builderSelect := sq.Select(columnName, columnDomainID).
		From(tableCategory).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{columnID: id})

	queryRaw, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}

	query := db.Query{
		Name:     "category_repository.get",
		QueryRaw: queryRaw,
	}

	categoryRepo := &modelRepo.Category{}
	err = r.db.DB().ScanOne(ctx, categoryRepo, query, args...)
	if err != nil {
		return nil, err
	}

	category := converter.CategoryFromPGSQL(categoryRepo)

	return category, nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builderDelete := sq.Delete(tableCategory).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{columnID: id})

	queryRaw, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}

	query := db.Query{
		Name:     "category_repository.delete",
		QueryRaw: queryRaw,
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Update(ctx context.Context, id int64, category *model.UpdatedCategory) error {
	categoryRepo := converter.UpdatedCategoryToPGSQL(category)
	builderUpdate := sq.Update(tableCategory).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{columnID: id})

	if categoryRepo.Name.Valid {
		builderUpdate = builderUpdate.Set(columnName, categoryRepo.Name)
	}

	if categoryRepo.DomainID.Valid {
		builderUpdate = builderUpdate.Set(columnDomainID, categoryRepo.DomainID)
	}

	queryRaw, args, err := builderUpdate.ToSql()
	if err != nil {
		return err
	}

	query := db.Query{
		Name:     "category_repository.update",
		QueryRaw: queryRaw,
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
