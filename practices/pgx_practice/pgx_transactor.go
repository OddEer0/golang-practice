package pgxPractice

import (
	"context"

	resourcePort "github.com/OddEer0/golang-practice/resources/resource_port"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionKey string

const (
	PgxTransactionKey TransactionKey = "pgx_transaction_key"
)

type pgxTransaction struct {
	conn *pgxpool.Pool
}

func InjectTransaction(ctx context.Context, tx QueryExecutor) context.Context {
	return context.WithValue(ctx, PgxTransactionKey, tx)
}

func ExtractTransaction(ctx context.Context) QueryExecutor {
	if tx, ok := ctx.Value(PgxTransactionKey).(QueryExecutor); ok {
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

func NewPgxTransaction() resourcePort.Transactor {
	return &pgxTransaction{}
}
