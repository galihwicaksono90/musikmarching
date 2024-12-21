package dbpool

import (
  "context"
  "time"
	"fmt"
	"galihwicaksono90/musikmarching-be/pkg/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDBPool(cfg config.Config) (*pgxpool.Pool, error) {
	// Construct connection string
	// connString := fmt.Sprintf(
	// 	"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	// 	cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	// )

	// Create a connection pool configuration
	poolConfig, err := pgxpool.ParseConfig(cfg.DB_SOURCE)
	if err != nil {
		return nil, fmt.Errorf("error parsing pool config: %v", err)
	}

	// Set pool configuration options
	poolConfig.MaxConns = 10                   // Maximum number of connections in the pool
	poolConfig.MinConns = 2                    // Minimum number of connections in the pool
	poolConfig.MaxConnLifetime = 1 * time.Hour // Maximum lifetime of a connection

	// Create the connection pool
	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("error creating connection pool: %v", err)
	}

	// Test the connection
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	return pool, nil
}
