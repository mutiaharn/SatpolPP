package ports

import (
	"context"
	"satpolpp-sulsel-api/internal/core/domain"
)

// UserRepository interface
type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, id uint) (*domain.User, error)
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	List(ctx context.Context, limit, offset int) ([]domain.User, int64, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id uint) error
}

// ArticleRepository interface
type ArticleRepository interface {
	Create(ctx context.Context, article *domain.Article) error
	GetByID(ctx context.Context, id uint) (*domain.Article, error)
	GetBySlug(ctx context.Context, slug string) (*domain.Article, error)
	List(ctx context.Context, category string, search string, limit, offset int) ([]domain.Article, int64, error)
	Update(ctx context.Context, article *domain.Article) error
	Delete(ctx context.Context, id uint) error
	IncrementView(ctx context.Context, id uint) error
}

// DocumentRepository interface
type DocumentRepository interface {
	Create(ctx context.Context, doc *domain.PublicDocument) error
	GetByID(ctx context.Context, id uint) (*domain.PublicDocument, error)
	List(ctx context.Context, category string, year int, search string, limit, offset int) ([]domain.PublicDocument, int64, error)
	Update(ctx context.Context, doc *domain.PublicDocument) error
	Delete(ctx context.Context, id uint) error
}

// ComplaintRepository interface
type ComplaintRepository interface {
	Create(ctx context.Context, complaint *domain.Complaint) error
	GetByID(ctx context.Context, id uint) (*domain.Complaint, error)
	GetByTicket(ctx context.Context, ticketCode string) (*domain.Complaint, error)
	List(ctx context.Context, status domain.ComplaintStatus, regencyID string, limit, offset int) ([]domain.Complaint, int64, error)
	Update(ctx context.Context, complaint *domain.Complaint) error
	GetStats(ctx context.Context) (map[string]int64, error)
}

// SettingRepository interface for Anti-Hardcoded CMS
type SettingRepository interface {
	Get(ctx context.Context, key string) (string, error)
	GetAll(ctx context.Context) (map[string]string, error)
	Set(ctx context.Context, key, value, group string) error
	SetMultiple(ctx context.Context, settings map[string]string) error
}

// AuditLogRepository interface
type AuditLogRepository interface {
	Create(ctx context.Context, log *domain.AuditLog) error
	List(ctx context.Context, limit, offset int) ([]domain.AuditLog, int64, error)
}
