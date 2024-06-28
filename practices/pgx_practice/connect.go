package pgxpractice

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectPg(connstr string) *pgxpool.Pool {
	cfg, err := pgxpool.ParseConfig(connstr)
	if err != nil {
		panic(err)
	}

	cfg.MaxConns = 20
	cfg.MinConns = 4
	cfg.MaxConnLifetime = time.Hour
	cfg.MaxConnIdleTime = time.Minute * 20
	cfg.HealthCheckPeriod = time.Minute * 2

	poolconn, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		panic(err)
	}

	return poolconn
}
