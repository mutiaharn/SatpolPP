package postgres

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"satpolpp-sulsel-api/internal/core/domain"
	"satpolpp-sulsel-api/internal/utils"
)

type TimelineEntry struct {
	Status    string    `json:"status"`
	Note      string    `json:"note"`
	Actor     string    `json:"actor"`
	Timestamp time.Time `json:"timestamp"`
}

type ComplaintRecord struct {
	domain.Complaint
	ReporterNIKEnc   string          `json:"-"`
	ReporterPhoneEnc string          `json:"-"`
	ReporterNIKMask  string          `json:"reporter_nik_masked"`
	Timeline         []TimelineEntry `json:"timeline"`
	Photos           []string        `json:"photos"`
}

// CreateComplaintTx atomically persists a complaint and queues its confirmation email in email_outbox.
func (d *DB) CreateComplaintTx(ctx context.Context, c *ComplaintRecord, photos []string, outboxSubject, outboxText, outboxHTML string) error {
	photosJSON, _ := json.Marshal(photos)
	timeline := []TimelineEntry{{
		Status:    string(domain.StatusPending),
		Note:      "Laporan diterima oleh sistem SIAP TERTIB dan menunggu verifikasi petugas.",
		Actor:     "Sistem SIAP TERTIB",
		Timestamp: time.Now(),
	}}
	timelineJSON, _ := json.Marshal(timeline)
	c.Timeline = timeline
	c.Photos = photos

	for attempt := 0; attempt < 3; attempt++ {
		if c.TicketCode == "" || attempt > 0 {
			c.TicketCode = utils.GenerateTicketCode()
		}

		err := d.WithTx(ctx, func(tx pgx.Tx) error {
			err := tx.QueryRow(ctx, `
				INSERT INTO complaints (
					ticket_code, is_anonymous, reporter_nik_enc, reporter_nik_masked,
					reporter_name, reporter_phone_enc, reporter_email, violation_type,
					regency_id, district_name, village_name, address_detail,
					latitude, longitude, description, incident_date,
					evidence_photos, status, timeline, updated_at
				) VALUES (
					$1, $2, $3, $4,
					$5, $6, $7, $8,
					$9, $10, $11, $12,
					$13, $14, $15, $16,
					$17::jsonb, $18::"ComplaintStatus", $19::jsonb, NOW()
				) RETURNING id, created_at, updated_at
			`,
				c.TicketCode, c.IsAnonymous, c.ReporterNIKEnc, c.ReporterNIKMask,
				c.ReporterName, c.ReporterPhoneEnc, c.ReporterEmail, c.ViolationType,
				c.RegencyID, c.DistrictName, c.VillageName, c.AddressDetail,
				c.Latitude, c.Longitude, c.Description, c.IncidentDate,
				string(photosJSON), string(domain.StatusPending), string(timelineJSON),
			).Scan(&c.ID, &c.CreatedAt, &c.UpdatedAt)
			if err != nil {
				return err
			}

			if c.ReporterEmail != "" {
				_, err = tx.Exec(ctx, `
					INSERT INTO email_outbox (complaint_id, to_email, subject, body_text, body_html, status)
					VALUES ($1, $2, $3, $4, $5, 'PENDING'::"OutboxStatus")
				`, c.ID, c.ReporterEmail, outboxSubject, outboxText, outboxHTML)
				if err != nil {
					return err
				}
			}
			return nil
		})

		if err == nil {
			return nil
		}
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			continue // Retry on rare ticket_code collision
		}
		return err
	}
	return errors.New("gagal menghasilkan kode tiket unik setelah 3 percobaan")
}

// GetComplaintByTicket fetches a complaint joined with regency and officer metadata.
func (d *DB) GetComplaintByTicket(ctx context.Context, ticketCode string) (*ComplaintRecord, error) {
	var rec ComplaintRecord
	var photosRaw, timelineRaw []byte

	err := d.Pool.QueryRow(ctx, `
		SELECT c.id, c.ticket_code, c.is_anonymous, c.reporter_nik_enc, c.reporter_nik_masked,
		       c.reporter_name, c.reporter_phone_enc, c.reporter_email, c.violation_type,
		       c.regency_id, COALESCE(r.name, ''), c.district_name, c.village_name, c.address_detail,
		       c.latitude, c.longitude, c.description, c.incident_date,
		       c.evidence_photos, c.status::text, c.disposed_to, c.action_notes, c.action_proof,
		       c.timeline, c.officer_id, COALESCE(u.name, ''), c.completed_at, c.created_at, c.updated_at
		FROM complaints c
		LEFT JOIN regencies r ON r.id = c.regency_id
		LEFT JOIN users u ON u.id = c.officer_id
		WHERE c.ticket_code = $1
	`, ticketCode).Scan(
		&rec.ID, &rec.TicketCode, &rec.IsAnonymous, &rec.ReporterNIKEnc, &rec.ReporterNIKMask,
		&rec.ReporterName, &rec.ReporterPhoneEnc, &rec.ReporterEmail, &rec.ViolationType,
		&rec.RegencyID, &rec.RegencyName, &rec.DistrictName, &rec.VillageName, &rec.AddressDetail,
		&rec.Latitude, &rec.Longitude, &rec.Description, &rec.IncidentDate,
		&photosRaw, &rec.Status, &rec.DisposedTo, &rec.ActionNotes, &rec.ActionProof,
		&timelineRaw, &rec.OfficerID, &rec.OfficerName, &rec.CompletedAt, &rec.CreatedAt, &rec.UpdatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	_ = json.Unmarshal(photosRaw, &rec.Photos)
	_ = json.Unmarshal(timelineRaw, &rec.Timeline)
	if len(rec.Photos) > 0 {
		rec.EvidenceImageURL = rec.Photos[0]
	}
	return &rec, nil
}

// ListComplaints returns complaints ordered newest first, optionally filtered by status.
func (d *DB) ListComplaints(ctx context.Context, statusFilter string) ([]ComplaintRecord, error) {
	query := `
		SELECT c.id, c.ticket_code, c.is_anonymous, c.reporter_nik_enc, c.reporter_nik_masked,
		       c.reporter_name, c.reporter_phone_enc, c.reporter_email, c.violation_type,
		       c.regency_id, COALESCE(r.name, ''), c.district_name, c.village_name, c.address_detail,
		       c.latitude, c.longitude, c.description, c.incident_date,
		       c.evidence_photos, c.status::text, c.disposed_to, c.action_notes, c.action_proof,
		       c.timeline, c.officer_id, COALESCE(u.name, ''), c.completed_at, c.created_at, c.updated_at
		FROM complaints c
		LEFT JOIN regencies r ON r.id = c.regency_id
		LEFT JOIN users u ON u.id = c.officer_id
	`
	var args []any
	if statusFilter != "" {
		query += ` WHERE c.status::text = $1`
		args = append(args, statusFilter)
	}
	query += ` ORDER BY c.created_at DESC`

	rows, err := d.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []ComplaintRecord
	for rows.Next() {
		var rec ComplaintRecord
		var photosRaw, timelineRaw []byte
		if err := rows.Scan(
			&rec.ID, &rec.TicketCode, &rec.IsAnonymous, &rec.ReporterNIKEnc, &rec.ReporterNIKMask,
			&rec.ReporterName, &rec.ReporterPhoneEnc, &rec.ReporterEmail, &rec.ViolationType,
			&rec.RegencyID, &rec.RegencyName, &rec.DistrictName, &rec.VillageName, &rec.AddressDetail,
			&rec.Latitude, &rec.Longitude, &rec.Description, &rec.IncidentDate,
			&photosRaw, &rec.Status, &rec.DisposedTo, &rec.ActionNotes, &rec.ActionProof,
			&timelineRaw, &rec.OfficerID, &rec.OfficerName, &rec.CompletedAt, &rec.CreatedAt, &rec.UpdatedAt,
		); err != nil {
			return nil, err
		}
		_ = json.Unmarshal(photosRaw, &rec.Photos)
		_ = json.Unmarshal(timelineRaw, &rec.Timeline)
		if len(rec.Photos) > 0 {
			rec.EvidenceImageURL = rec.Photos[0]
		}
		list = append(list, rec)
	}
	return list, rows.Err()
}

// UpdateComplaintStatusTx uses SELECT ... FOR UPDATE to prevent concurrent status overwrite races.
func (d *DB) UpdateComplaintStatusTx(
	ctx context.Context,
	ticketCode string,
	newStatus domain.ComplaintStatus,
	disposedTo, actionNotes, actionProof string,
	officerID uint,
	officerUsername, officerName, ipAddr string,
	outboxSubject, outboxText, outboxHTML string,
) (*ComplaintRecord, error) {
	var updated *ComplaintRecord

	err := d.WithTx(ctx, func(tx pgx.Tx) error {
		var id uint
		var email string
		var timelineRaw []byte
		err := tx.QueryRow(ctx, `
			SELECT id, reporter_email, timeline
			FROM complaints
			WHERE ticket_code = $1
			FOR UPDATE
		`, ticketCode).Scan(&id, &email, &timelineRaw)
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrNotFound
		}
		if err != nil {
			return err
		}

		var timeline []TimelineEntry
		_ = json.Unmarshal(timelineRaw, &timeline)
		timeline = append(timeline, TimelineEntry{
			Status:    string(newStatus),
			Note:      actionNotes,
			Actor:     officerName,
			Timestamp: time.Now(),
		})
		newTimelineJSON, _ := json.Marshal(timeline)

		var completedAt *time.Time
		if newStatus == domain.StatusCompleted || newStatus == domain.StatusRejected {
			now := time.Now()
			completedAt = &now
		}

		_, err = tx.Exec(ctx, `
			UPDATE complaints
			SET status = $1::"ComplaintStatus",
			    disposed_to = CASE WHEN $2 <> '' THEN $2 ELSE disposed_to END,
			    action_notes = CASE WHEN $3 <> '' THEN $3 ELSE action_notes END,
			    action_proof = CASE WHEN $4 <> '' THEN $4 ELSE action_proof END,
			    timeline = $5::jsonb,
			    officer_id = $6,
			    completed_at = COALESCE($7, completed_at),
			    updated_at = NOW()
			WHERE id = $8
		`, string(newStatus), disposedTo, actionNotes, actionProof, string(newTimelineJSON), officerID, completedAt, id)
		if err != nil {
			return err
		}

		_, err = tx.Exec(ctx, `
			INSERT INTO audit_logs (user_id, actor_username, action, resource, details, ip_address)
			VALUES ($1, $2, 'UPDATE_COMPLAINT_STATUS', $3, $4, $5)
		`, officerID, officerUsername, ticketCode, fmt.Sprintf("Status diubah ke %s (%s)", newStatus, disposedTo), ipAddr)
		if err != nil {
			return err
		}

		if email != "" {
			_, err = tx.Exec(ctx, `
				INSERT INTO email_outbox (complaint_id, to_email, subject, body_text, body_html, status)
				VALUES ($1, $2, $3, $4, $5, 'PENDING'::"OutboxStatus")
			`, id, email, outboxSubject, outboxText, outboxHTML)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	updated, err = d.GetComplaintByTicket(ctx, ticketCode)
	return updated, err
}
