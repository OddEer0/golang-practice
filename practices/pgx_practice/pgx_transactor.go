package pgxPractice

import (
	"context"
	"github.com/jackc/pgx/v5"

	"github.com/OddEer0/golang-practice/resources/sql"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionKey string

const (
	PgxTransactionKey TransactionKey = "pgx_transaction_key"
)

type PgxTransaction struct {
	Conn         *pgxpool.Pool
	TxController sql.TransactionController
}

type PgxTransactionController struct {
}

func (p PgxTransactionController) InjectTransaction(ctx context.Context, tx sql.QueryExecutor) context.Context {
	return context.WithValue(ctx, PgxTransactionKey, tx)
}

func (p PgxTransactionController) ExtractTransaction(ctx context.Context) sql.QueryExecutor {
	if tx, ok := ctx.Value(PgxTransactionKey).(sql.QueryExecutor); ok {
		return tx
	}
	return nil
}

func (p *PgxTransaction) WithinTransaction(ctx context.Context, callback func(context.Context) error) error {
	tx, err := p.Conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			panic(err)
		}
	}(tx, ctx)

	err = callback(p.TxController.InjectTransaction(ctx, tx))
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}
