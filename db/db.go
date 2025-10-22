package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Responsible for creating connection to postgres db and returning a connection pool
// or an error
func NewPostgresStorage(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	if dsn == "" {
		return nil, fmt.Errorf("DATABASE_URL not set")
	}

	// convertin the db url to a config struct
	config, err := pgxpool.ParseConfig(dsn)

	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}

	// adding optional configurations
	config.MaxConns = 10 // upto 10 simultaneous connections to db
	config.MinConns = 2
	config.MaxConnLifetime = time.Hour // automatically recycle a conn after an hour

	// creating a pool of connections with provided configs
	pool, err := pgxpool.NewWithConfig(ctx, config)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return pool, nil
}
