package pgxPractice

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectPg(config *Config) *pgxpool.Pool {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Postgres.Host, config.Postgres.Port, config.Postgres.User, config.Postgres.Password, config.Postgres.DbName)

	cfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		panic(err)
	}

	cfg.MaxConns = 20
	cfg.MinConns = 4
	cfg.MaxConnLifetime = time.Hour
	cfg.MaxConnIdleTime = time.Minute * 20
	cfg.HealthCheckPeriod = time.Minute * 2

	poolConn, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		panic(err)
	}

	return poolConn
}
