package migrations

import (
	"context"
	"embed"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed migrations
var migrationFS embed.FS

func Up(ctx context.Context, pool *pgxpool.Pool) (err error) {
	conn := stdlib.OpenDBFromPool(pool)

	defer func() {
		err = errors.Join(err, conn.Close())
	}()

	err = goose.SetDialect("postgres")
	if err != nil {
		return
	}

	goose.SetBaseFS(migrationFS)

	return goose.UpContext(ctx, conn, "migrations")
}
