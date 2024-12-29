package dbpool

import (
	"context"
	"fmt"
	"galihwicaksono90/musikmarching-be/pkg/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

func NewDBPool(cfg config.Config) (*pgxpool.Pool, error) {
	// Construct connection string
	// connString := fmt.Sprintf(
	//
	// 	"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	// 	cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	// )
	connString := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	fmt.Println(connString)

	// Create a connection pool configuration
	poolConfig, err := pgxpool.ParseConfig(connString)
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
