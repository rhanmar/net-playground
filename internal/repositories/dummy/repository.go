package dummy

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type db interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type Repository struct {
	db db
}

func NewRepository(db db) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(ctx context.Context, data string) error {
	sql := "INSERT INTO dummy (data, created_at) VALUES ($1, $2)"
	_, err := r.db.Exec(
		ctx,
		sql,
		data,
		time.Now(),
	)
	if err != nil {
		return errors.Wrap(err, "db.Exec")
	}
	return nil
}
