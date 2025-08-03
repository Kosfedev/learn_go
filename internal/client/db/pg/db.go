package pg

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/Kosfedev/learn_go/internal/client/db"
)

var _ db.DB = (*pg)(nil)

type pg struct {
	dbc *pgxpool.Pool
}

func newDB(dbc *pgxpool.Pool) *pg { return &pg{dbc: dbc} }

func (p *pg) ScanOne(ctx context.Context, dest interface{}, query db.Query, args ...interface{}) error {
	rows, err := p.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanOne(dest, rows)
}

func (p *pg) ScanAll(ctx context.Context, dest interface{}, query db.Query, args ...interface{}) error {
	rows, err := p.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanAll(dest, rows)
}

func (p *pg) ExecContext(ctx context.Context, query db.Query, args ...interface{}) (pgconn.CommandTag, error) {
	return p.dbc.Exec(ctx, query.QueryRaw, args...)
}

func (p *pg) QueryContext(ctx context.Context, query db.Query, args ...interface{}) (pgx.Rows, error) {
	return p.dbc.Query(ctx, query.QueryRaw, args...)
}

func (p *pg) QueryRowContext(ctx context.Context, query db.Query, args ...interface{}) pgx.Row {
	return p.dbc.QueryRow(ctx, query.QueryRaw, args...)
}

func (p *pg) Ping(ctx context.Context) error {
	return p.dbc.Ping(ctx)
}

func (p *pg) Close() {
	p.dbc.Close()
}
