package db

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type Client interface {
	DB() DB
	Closer() error
}

type Query struct {
	Name     string
	QueryRaw string
}

type SQLExecer interface {
	NamedExecer
	QueryExecer
}

type NamedExecer interface {
	ScanOne(ctx context.Context, dest interface{}, query Query, args ...interface{}) error
	ScanAll(ctx context.Context, dest interface{}, query Query, args ...interface{}) error
}

type QueryExecer interface {
	ExecContext(ctx context.Context, query Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, query Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, query Query, args ...interface{}) pgx.Row
}

type Pinger interface {
	Ping(ctx context.Context) error
}

type DB interface {
	SQLExecer
	Pinger
	Close()
}
