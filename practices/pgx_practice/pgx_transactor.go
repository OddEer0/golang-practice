package pgxPractice

import (
	"context"
	"log/slog"

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
	//defer func() {
	//	err := tx.Rollback(ctx)
	//	if err != nil {
	//		slog.Error("rollback error", slog.Any("cause", err))
	//	}
	//}()

	err = callback(p.TxController.InjectTransaction(ctx, tx))
	if err != nil {
		errLoc := tx.Rollback(ctx)
		if errLoc != nil {
			slog.Error("rollback error", slog.Any("cause", err))
		}
		slog.Info("rollback transaction")
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		err := tx.Rollback(ctx)
		if err != nil {
			slog.Error("rollback error", slog.Any("cause", err))
		}
		slog.Info("rollback transaction")
		return err
	}
	slog.Info("commit transaction")
	return nil
}
