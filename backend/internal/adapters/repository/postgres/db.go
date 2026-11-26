package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"satpolpp-sulsel-api/internal/config"
)

type DB struct {
	Pool *pgxpool.Pool
	Cfg  *config.Config
}

// NewDB initializes a pgxpool configured for PgBouncer transaction pooling (SimpleProtocol).
func NewDB(ctx context.Context, cfg *config.Config) (*DB, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSSLMode,
	)

	poolCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("invalid postgres dsn: %w", err)
	}

	// Required when connecting through PgBouncer in pool_mode=transaction
	poolCfg.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol
	poolCfg.MaxConns = 25
	poolCfg.MinConns = 2
	poolCfg.MaxConnLifetime = 30 * time.Minute
	poolCfg.MaxConnIdleTime = 5 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create pgx pool: %w", err)
	}

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := pool.Ping(pingCtx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("postgres ping failed: %w", err)
	}

	return &DB{Pool: pool, Cfg: cfg}, nil
}

// WithTx executes fn inside an ACID PostgreSQL transaction, rolling back on error or panic.
func (d *DB) WithTx(ctx context.Context, fn func(tx pgx.Tx) error) error {
	tx, err := d.Pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(ctx)
			panic(p)
		}
	}()

	if err := fn(tx); err != nil {
		_ = tx.Rollback(ctx)
		return err
	}
	return tx.Commit(ctx)
}

// IsTokenRevoked checks whether a JWT SHA-256 hash exists in revoked_tokens.
func (d *DB) IsTokenRevoked(ctx context.Context, tokenHash string) (bool, error) {
	var exists bool
	err := d.Pool.QueryRow(ctx,
		`SELECT EXISTS(SELECT 1 FROM revoked_tokens WHERE token_hash = $1 AND expires_at > NOW())`,
		tokenHash,
	).Scan(&exists)
	return exists, err
}

// RevokeToken inserts a JWT hash into revoked_tokens and prunes expired rows.
func (d *DB) RevokeToken(ctx context.Context, tokenHash string, expiresAt time.Time) error {
	_, err := d.Pool.Exec(ctx,
		`INSERT INTO revoked_tokens (token_hash, expires_at) VALUES ($1, $2) ON CONFLICT (token_hash) DO NOTHING`,
		tokenHash, expiresAt,
	)
	if err != nil {
		return err
	}
	_, _ = d.Pool.Exec(ctx, `DELETE FROM revoked_tokens WHERE expires_at <= NOW()`)
	return nil
}
