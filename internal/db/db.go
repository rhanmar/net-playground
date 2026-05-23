package db

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pkg/errors"
	"github.com/pressly/goose/v3"
)

type config interface {
	GetPostgresDSN() string
	GetMigrationsDir() string
}

type DB struct {
	*pgxpool.Pool
}

func NewDB(ctx context.Context, config config) (*DB, error) {
	pool, err := pgxpool.New(ctx, config.GetPostgresDSN())
	if err != nil {
		return nil, errors.Wrap(err, "pgxpool.New")
	}
	err = pool.Ping(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "pool.Ping")
	}

	// необходимо сделать такую конвертацию, так как goose не умеет работать с pgx.Pool
	sqlDB := stdlib.OpenDBFromPool(pool)
	defer func() {
		err = sqlDB.Close()
		if err != nil {
			slog.Error(fmt.Sprintf("stdlib.OpenDBFromPool: %v", err))
		}
	}()
	err = sqlDB.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "sqlDB.Ping")
	}
	err = goose.Up(sqlDB, config.GetMigrationsDir())
	if err != nil {
		return nil, errors.Wrap(err, "goose.Up")
	}
	return &DB{pool}, nil
}
