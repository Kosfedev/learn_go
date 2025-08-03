package pg

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/Kosfedev/learn_go/internal/client/db"
)

var _ db.Client = (*Client)(nil)

type Client struct {
	masterDBC db.DB
}

func New(ctx context.Context, dsn string) (*Client, error) {
	dbc, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return &Client{
		masterDBC: newDB(dbc),
	}, nil
}

func (c *Client) DB() db.DB {
	return c.masterDBC
}

func (c *Client) Closer() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}

	return nil
}
