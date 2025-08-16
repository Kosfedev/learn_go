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
	"github.com/Kosfedev/learn_go/internal/repository/set/pg/converter"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/set/pg/model"
)

const (
	tableSet        = "set"
	columnID        = "id"
	columnName      = "name"
	columnCreatedAt = "created_at"
	columnUpdatedAt = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.SetRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(ctx context.Context, newSet *model.NewSet) (int64, error) {
	newSetRepo := converter.NewSetToPGSQL(newSet)
	builderInsert := sq.Insert(tableSet).
		PlaceholderFormat(sq.Dollar).
		Columns(columnName).
		Values(newSetRepo.Name).
		Suffix(fmt.Sprintf("RETURNING %v", columnID))

	queryRaw, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	query := db.Query{
		Name:     "set_repository.create",
		QueryRaw: queryRaw,
	}

	var setID int64
	err = r.db.DB().ScanOne(ctx, &setID, query, args...)
	if err != nil {
		return 0, err
	}

	return setID, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.Set, error) {
	setRepo := &modelRepo.Set{}
	builderSelect := sq.Select(columnID, columnName, columnCreatedAt, columnUpdatedAt).
		From(tableSet).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{columnID: id})

	queryRaw, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}

	query := db.Query{
		Name:     "set_repository.get",
		QueryRaw: queryRaw,
	}

	err = r.db.DB().ScanOne(ctx, setRepo, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	set := converter.SetFromPGSQL(setRepo)

	return set, nil
}

func (r *repo) Update(ctx context.Context, id int64, updatedSet *model.UpdatedSet) error {
	updatedSetRepo := converter.UpdatedSetToPGSQL(updatedSet)
	builderUpdate := sq.Update(tableSet).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{columnID: id})

	if updatedSetRepo.Name.Valid {
		builderUpdate = builderUpdate.Set(columnName, updatedSetRepo.Name)
	}

	queryRaw, args, err := builderUpdate.ToSql()
	if err != nil {
		return err
	}

	query := db.Query{
		Name:     "set_repository.update",
		QueryRaw: queryRaw,
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builderDelete := sq.Delete(tableSet).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{columnID: id})

	queryRaw, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}

	query := db.Query{
		Name:     "set_repository.delete",
		QueryRaw: queryRaw,
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

// TODO: подчистить
func (r *repo) ListSearch(ctx context.Context, listSearchOptions *model.SetListSearchOptions) (*model.SetListWithTotal, error) {
	if listSearchOptions == nil {
		return nil, errors.New("listSearchOptions cannot be nil")
	}
	// TODO:вынести
	sortOrder := "ASC"
	if listSearchOptions.Sort.SortOrder == 1 {
		sortOrder = "DESC"
	}

	setListRepo := &modelRepo.SetListWithTotal{}

	nameFilter := "%" + listSearchOptions.Filters.Name + "%"
	subquery := `SELECT * FROM set`
	if len(listSearchOptions.Filters.Name) > 0 {
		subquery = fmt.Sprintf("%s WHERE name ILIKE '%s'", subquery, nameFilter)
	}

	subquery = fmt.Sprintf(
		"%s ORDER BY %s %s OFFSET %d LIMIT %d",
		subquery,
		listSearchOptions.Sort.SortBy,
		sortOrder,
		uint64(listSearchOptions.Pagination.Offset),
		uint64(listSearchOptions.Pagination.Limit),
	)

	queryRaw := fmt.Sprintf("SELECT (SELECT json_agg(set) FROM (%s) AS set) AS sets, count(*) as total FROM set WHERE name ILIKE '%s'", subquery, nameFilter)

	query := db.Query{
		Name:     "set_repository.list_search",
		QueryRaw: queryRaw,
	}

	err := r.db.DB().ScanOne(ctx, setListRepo, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	setList := &model.SetListWithTotal{
		Sets:  converter.SetsFromPGSQL(setListRepo.Sets),
		Total: setListRepo.Total,
	}

	return setList, nil
}
