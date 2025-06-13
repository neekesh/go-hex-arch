package storage

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxUnitOfWork struct {
	db *pgxpool.Pool
}

func NewPgxUnitOfWork(db *pgxpool.Pool) *PgxUnitOfWork {
	return &PgxUnitOfWork{db: db}
}

type txKey struct{}

func (u *PgxUnitOfWork) Do(ctx context.Context, fn func(txCtx context.Context) error) error {
	tx, err := u.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(ctx)
			panic(p)
		}
	}()

	txCtx := context.WithValue(ctx, txKey{}, tx)

	if err := fn(txCtx); err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}

func GetTx(ctx context.Context) pgx.Tx {
	return ctx.Value(txKey{}).(pgx.Tx)
}
