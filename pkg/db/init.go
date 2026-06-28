package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func configurate(cfg *pgxpool.Config) {
	cfg.MaxConns = 10
	cfg.MinConns = 2
	cfg.MaxConnLifetime = 1 * time.Hour
	cfg.MaxConnIdleTime = 30 * time.Minute
}

func pageTable(ctx context.Context, pool *pgxpool.Pool) error {
	createQuery := `
	CREATE TABLE IF NOT EXISTS book_pages (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		book_id UUID NOT NULL REFERENCES books(id) ON DELETE CASCADE,
		number INT NOT NULL,
		content TEXT NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
	
		UNIQUE(book_id, number)
	);`

	_, err := pool.Exec(ctx, createQuery)
	return err
}

func bookTable(ctx context.Context, pool *pgxpool.Pool) error {
	createQuery := `
	CREATE TABLE IF NOT EXISTS books (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		title VARCHAR(100) NOT NULL,
		description VARCHAR(1000) NOT NULL,
		created_year INT NOT NULL,
		total_pages INT NOT NULL DEFAULT 0,
		genre VARCHAR(100) NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
	
		UNIQUE(title)
	);`

	_, err := pool.Exec(ctx, createQuery)
	return err
}

func InitDB(url string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database url: %w", err)
	}

	configurate(cfg)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to connection to database: %w", err)
	}

	if err := pageTable(ctx, pool); err != nil {
		return nil, fmt.Errorf("failed to create books page table: %w", err)
	}

	if err := bookTable(ctx, pool); err != nil {
		return nil, fmt.Errorf("failed to create books table: %w", err)
	}

	return pool, nil
}
