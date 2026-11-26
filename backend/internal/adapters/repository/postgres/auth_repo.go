package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"satpolpp-sulsel-api/internal/core/domain"
)

var ErrNotFound = errors.New("record not found")

func (d *DB) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	var u domain.User
	err := d.Pool.QueryRow(ctx, `
		SELECT id, name, email, username, password, role::text, bidang, is_active, created_at, updated_at
		FROM users WHERE username = $1
	`, username).Scan(
		&u.ID, &u.Name, &u.Email, &u.Username, &u.Password, &u.Role, &u.Bidang, &u.IsActive, &u.CreatedAt, &u.UpdatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	return &u, err
}

func (d *DB) ListUsers(ctx context.Context) ([]domain.User, error) {
	rows, err := d.Pool.Query(ctx, `
		SELECT id, name, email, username, role::text, bidang, is_active, created_at, updated_at
		FROM users ORDER BY id ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var u domain.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Username, &u.Role, &u.Bidang, &u.IsActive, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

func (d *DB) CreateUser(ctx context.Context, u *domain.User) error {
	return d.Pool.QueryRow(ctx, `
		INSERT INTO users (name, email, username, password, role, bidang, is_active, updated_at)
		VALUES ($1, $2, $3, $4, $5::"Role", $6, $7, NOW())
		RETURNING id, created_at, updated_at
	`, u.Name, u.Email, u.Username, u.Password, string(u.Role), u.Bidang, u.IsActive).Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt)
}

func (d *DB) SetUserActive(ctx context.Context, id uint, isActive bool) error {
	tag, err := d.Pool.Exec(ctx, `UPDATE users SET is_active = $1, updated_at = NOW() WHERE id = $2`, isActive, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func (d *DB) RecordAuditLog(ctx context.Context, log *domain.AuditLog) error {
	_, err := d.Pool.Exec(ctx, `
		INSERT INTO audit_logs (user_id, actor_username, action, resource, details, ip_address)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, log.UserID, log.UserName, log.Action, log.Resource, log.Details, log.IPAddress)
	return err
}

func (d *DB) ListAuditLogs(ctx context.Context, limit int) ([]domain.AuditLog, error) {
	if limit <= 0 || limit > 200 {
		limit = 50
	}
	rows, err := d.Pool.Query(ctx, `
		SELECT id, user_id, actor_username, action, resource, details, ip_address, created_at
		FROM audit_logs ORDER BY created_at DESC LIMIT $1
	`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []domain.AuditLog
	for rows.Next() {
		var l domain.AuditLog
		if err := rows.Scan(&l.ID, &l.UserID, &l.UserName, &l.Action, &l.Resource, &l.Details, &l.IPAddress, &l.CreatedAt); err != nil {
			return nil, err
		}
		logs = append(logs, l)
	}
	return logs, rows.Err()
}
