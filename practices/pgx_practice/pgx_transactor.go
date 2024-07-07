package pgxPractice

import (
	"context"

	"github.com/OddEer0/golang-practice/resources/sql"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionKey string

const (
	PgxTransactionKey TransactionKey = "pgx_transaction_key"
)

type pgxTransaction struct {
	conn *pgxpool.Pool
}

func InjectTransaction(ctx context.Context, tx sql.QueryExecutor) context.Context {
	return context.WithValue(ctx, PgxTransactionKey, tx)
}

func ExtractTransaction(ctx context.Context) sql.QueryExecutor {
	if tx, ok := ctx.Value(PgxTransactionKey).(sql.QueryExecutor); ok {
		return tx
	}
	return nil
}

func (p *pgxTransaction) WithinTransaction(ctx context.Context, callback func(context.Context) error) error {
	tx, err := p.conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	err = callback(InjectTransaction(ctx, tx))
	if err != nil {
		return err
	}

	tx.Commit(ctx)
	return nil
}
