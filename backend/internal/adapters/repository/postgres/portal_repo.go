package postgres

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/jackc/pgx/v5"
	"satpolpp-sulsel-api/internal/core/domain"
)

func (d *DB) ListRegencies(ctx context.Context) ([]domain.Regency, error) {
	rows, err := d.Pool.Query(ctx, `SELECT id, name FROM regencies ORDER BY id ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.Regency
	for rows.Next() {
		var r domain.Regency
		if err := rows.Scan(&r.ID, &r.Name); err != nil {
			return nil, err
		}
		list = append(list, r)
	}
	return list, rows.Err()
}

func (d *DB) GetSettingsMap(ctx context.Context) (map[string]string, error) {
	rows, err := d.Pool.Query(ctx, `SELECT key, value FROM site_settings`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make(map[string]string)
	for rows.Next() {
		var k, v string
		if err := rows.Scan(&k, &v); err != nil {
			return nil, err
		}
		res[k] = v
	}
	return res, rows.Err()
}

func (d *DB) UpsertSettings(ctx context.Context, settings map[string]string) error {
	return d.WithTx(ctx, func(tx pgx.Tx) error {
		for k, v := range settings {
			_, err := tx.Exec(ctx, `
				INSERT INTO site_settings (key, value, "group", updated_at)
				VALUES ($1, $2, 'general', NOW())
				ON CONFLICT (key) DO UPDATE SET value = EXCLUDED.value, updated_at = NOW()
			`, k, v)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (d *DB) ListArticles(ctx context.Context, onlyPublished bool) ([]domain.Article, error) {
	query := `
		SELECT a.id, a.title, a.slug, a.summary, a.content, a.thumbnail, a.category,
		       a.author_id, COALESCE(u.name, 'Humas Satpol PP'), a.view_count, a.is_published,
		       a.published_at, a.created_at, a.updated_at
		FROM articles a
		LEFT JOIN users u ON u.id = a.author_id
	`
	if onlyPublished {
		query += ` WHERE a.is_published = true`
	}
	query += ` ORDER BY a.published_at DESC`

	rows, err := d.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.Article
	for rows.Next() {
		var a domain.Article
		if err := rows.Scan(
			&a.ID, &a.Title, &a.Slug, &a.Summary, &a.Content, &a.Thumbnail, &a.Category,
			&a.AuthorID, &a.AuthorName, &a.ViewCount, &a.IsPublished,
			&a.PublishedAt, &a.CreatedAt, &a.UpdatedAt,
		); err != nil {
			return nil, err
		}
		list = append(list, a)
	}
	return list, rows.Err()
}

func (d *DB) GetArticleBySlug(ctx context.Context, slug string) (*domain.Article, error) {
	var a domain.Article
	err := d.Pool.QueryRow(ctx, `
		UPDATE articles SET view_count = view_count + 1 WHERE slug = $1
		RETURNING id, title, slug, summary, content, thumbnail, category, author_id, view_count, is_published, published_at, created_at, updated_at
	`, slug).Scan(
		&a.ID, &a.Title, &a.Slug, &a.Summary, &a.Content, &a.Thumbnail, &a.Category,
		&a.AuthorID, &a.ViewCount, &a.IsPublished, &a.PublishedAt, &a.CreatedAt, &a.UpdatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	_ = d.Pool.QueryRow(ctx, `SELECT name FROM users WHERE id = $1`, a.AuthorID).Scan(&a.AuthorName)
	return &a, nil
}

func (d *DB) CreateArticle(ctx context.Context, a *domain.Article) error {
	return d.Pool.QueryRow(ctx, `
		INSERT INTO articles (title, slug, summary, content, thumbnail, category, author_id, is_published, published_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW())
		RETURNING id, published_at, created_at, updated_at
	`, a.Title, a.Slug, a.Summary, a.Content, a.Thumbnail, a.Category, a.AuthorID, a.IsPublished).
		Scan(&a.ID, &a.PublishedAt, &a.CreatedAt, &a.UpdatedAt)
}

func (d *DB) DeleteArticle(ctx context.Context, id uint) error {
	tag, err := d.Pool.Exec(ctx, `DELETE FROM articles WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func (d *DB) ListDocuments(ctx context.Context, category string) ([]domain.PublicDocument, error) {
	query := `
		SELECT id, title, doc_number, category, year, file_url, file_size, download_count, created_at, updated_at
		FROM public_documents
	`
	var args []any
	if category != "" {
		query += ` WHERE LOWER(category) = LOWER($1)`
		args = append(args, category)
	}
	query += ` ORDER BY year DESC, id DESC`

	rows, err := d.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.PublicDocument
	for rows.Next() {
		var doc domain.PublicDocument
		if err := rows.Scan(&doc.ID, &doc.Title, &doc.DocNumber, &doc.Category, &doc.Year, &doc.FileURL, &doc.FileSize, &doc.DownloadCount, &doc.CreatedAt, &doc.UpdatedAt); err != nil {
			return nil, err
		}
		list = append(list, doc)
	}
	return list, rows.Err()
}

func (d *DB) CreateDocument(ctx context.Context, doc *domain.PublicDocument) error {
	return d.Pool.QueryRow(ctx, `
		INSERT INTO public_documents (title, doc_number, category, year, file_url, file_size, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW())
		RETURNING id, created_at, updated_at
	`, doc.Title, doc.DocNumber, doc.Category, doc.Year, doc.FileURL, doc.FileSize).
		Scan(&doc.ID, &doc.CreatedAt, &doc.UpdatedAt)
}

func (d *DB) DeleteDocument(ctx context.Context, id uint) error {
	tag, err := d.Pool.Exec(ctx, `DELETE FROM public_documents WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func (d *DB) ListGalleries(ctx context.Context) ([]domain.Gallery, error) {
	rows, err := d.Pool.Query(ctx, `
		SELECT id, title, description, image_url, event_date, created_at
		FROM galleries ORDER BY event_date DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.Gallery
	for rows.Next() {
		var g domain.Gallery
		if err := rows.Scan(&g.ID, &g.Title, &g.Description, &g.ImageURL, &g.EventDate, &g.CreatedAt); err != nil {
			return nil, err
		}
		list = append(list, g)
	}
	return list, rows.Err()
}

func (d *DB) CreateGallery(ctx context.Context, g *domain.Gallery) error {
	return d.Pool.QueryRow(ctx, `
		INSERT INTO galleries (title, description, image_url, event_date, updated_at)
		VALUES ($1, $2, $3, $4, NOW())
		RETURNING id, created_at
	`, g.Title, g.Description, g.ImageURL, g.EventDate).Scan(&g.ID, &g.CreatedAt)
}

func (d *DB) DeleteGallery(ctx context.Context, id uint) error {
	tag, err := d.Pool.Exec(ctx, `DELETE FROM galleries WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func (d *DB) ListOrganizations(ctx context.Context) ([]domain.OrganizationNode, error) {
	rows, err := d.Pool.Query(ctx, `
		SELECT id, org_code, parent_code, position_title, official_name, nip, rank, description, sub_units, photo_url, level, sort_order
		FROM organization_nodes ORDER BY level ASC, sort_order ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.OrganizationNode
	for rows.Next() {
		var o domain.OrganizationNode
		var subRaw []byte
		if err := rows.Scan(&o.ID, &o.OrgCode, &o.ParentCode, &o.PositionTitle, &o.OfficialName, &o.NIP, &o.Rank, &o.Description, &subRaw, &o.PhotoURL, &o.Level, &o.SortOrder); err != nil {
			return nil, err
		}
		_ = json.Unmarshal(subRaw, &o.SubUnits)
		list = append(list, o)
	}
	return list, rows.Err()
}

func (d *DB) UpsertOrganization(ctx context.Context, o *domain.OrganizationNode) error {
	subRaw, _ := json.Marshal(o.SubUnits)
	return d.Pool.QueryRow(ctx, `
		INSERT INTO organization_nodes (org_code, parent_code, position_title, official_name, nip, rank, description, sub_units, photo_url, level, sort_order, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8::jsonb, $9, $10, $11, NOW())
		ON CONFLICT (org_code) DO UPDATE SET
			parent_code = EXCLUDED.parent_code,
			position_title = EXCLUDED.position_title,
			official_name = EXCLUDED.official_name,
			nip = EXCLUDED.nip,
			rank = EXCLUDED.rank,
			description = EXCLUDED.description,
			sub_units = EXCLUDED.sub_units,
			photo_url = EXCLUDED.photo_url,
			level = EXCLUDED.level,
			sort_order = EXCLUDED.sort_order,
			updated_at = NOW()
		RETURNING id
	`, o.OrgCode, o.ParentCode, o.PositionTitle, o.OfficialName, o.NIP, o.Rank, o.Description, string(subRaw), o.PhotoURL, o.Level, o.SortOrder).Scan(&o.ID)
}

func (d *DB) DeleteOrganization(ctx context.Context, id uint) error {
	tag, err := d.Pool.Exec(ctx, `DELETE FROM organization_nodes WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
