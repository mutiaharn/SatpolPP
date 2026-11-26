package worker

import (
	"context"
	"fmt"
	"log/slog"
	"net/smtp"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"satpolpp-sulsel-api/internal/adapters/repository/postgres"
	"satpolpp-sulsel-api/internal/config"
)

// OutboxWorker polls email_outbox in PostgreSQL and dispatches queued emails reliably.
type OutboxWorker struct {
	db  *postgres.DB
	cfg *config.Config
}

func NewOutboxWorker(db *postgres.DB, cfg *config.Config) *OutboxWorker {
	return &OutboxWorker{db: db, cfg: cfg}
}

// Start runs the background polling loop until ctx is canceled.
func (w *OutboxWorker) Start(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	slog.Info("Email Outbox Worker aktif", "smtp_host", w.cfg.SMTPHost)

	for {
		select {
		case <-ctx.Done():
			slog.Info("Email Outbox Worker berhenti dengan aman")
			return
		case <-ticker.C:
			w.processBatch(ctx)
		}
	}
}

func (w *OutboxWorker) processBatch(ctx context.Context) {
	type job struct {
		id       int
		to       string
		subject  string
		bodyText string
		bodyHTML string
		attempts int
	}

	var jobs []job
	err := w.db.WithTx(ctx, func(tx pgx.Tx) error {
		rows, err := tx.Query(ctx, `
			SELECT id, to_email, subject, body_text, body_html, attempts
			FROM email_outbox
			WHERE status = 'PENDING'::"OutboxStatus" AND next_retry_at <= NOW()
			ORDER BY created_at ASC
			LIMIT 10
			FOR UPDATE SKIP LOCKED
		`)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var j job
			if err := rows.Scan(&j.id, &j.to, &j.subject, &j.bodyText, &j.bodyHTML, &j.attempts); err != nil {
				return err
			}
			jobs = append(jobs, j)
		}
		return rows.Err()
	})
	if err != nil {
		slog.Warn("Outbox poll error", "error", err)
		return
	}

	for _, j := range jobs {
		sendErr := w.dispatchSMTP(j.to, j.subject, j.bodyText, j.bodyHTML)
		if sendErr == nil {
			_, _ = w.db.Pool.Exec(ctx, `
				UPDATE email_outbox
				SET status = 'SENT'::"OutboxStatus", sent_at = NOW(), last_error = ''
				WHERE id = $1
			`, j.id)
			continue
		}

		nextAttempt := j.attempts + 1
		if nextAttempt >= 5 {
			_, _ = w.db.Pool.Exec(ctx, `
				UPDATE email_outbox
				SET status = 'FAILED'::"OutboxStatus", attempts = $1, last_error = $2
				WHERE id = $3
			`, nextAttempt, sendErr.Error(), j.id)
			slog.Error("Email outbox gagal permanen", "outbox_id", j.id, "to", j.to, "error", sendErr)
		} else {
			backoff := time.Duration(nextAttempt*30) * time.Second
			_, _ = w.db.Pool.Exec(ctx, `
				UPDATE email_outbox
				SET attempts = $1, last_error = $2, next_retry_at = $3
				WHERE id = $4
			`, nextAttempt, sendErr.Error(), time.Now().Add(backoff), j.id)
		}
	}
}

func (w *OutboxWorker) dispatchSMTP(to, subject, bodyText, bodyHTML string) error {
	if strings.TrimSpace(w.cfg.SMTPHost) == "" {
		slog.Info("Simulasi email outbox (SMTP_HOST kosong)", "to", to, "subject", subject)
		return nil
	}

	content := bodyText
	contentType := "text/plain; charset=UTF-8"
	if strings.TrimSpace(bodyHTML) != "" {
		content = bodyHTML
		contentType = "text/html; charset=UTF-8"
	}

	msg := strings.Join([]string{
		fmt.Sprintf("From: SIAP TERTIB Satpol PP Sulsel <%s>", w.cfg.SMTPFrom),
		fmt.Sprintf("Reply-To: %s", w.cfg.SMTPFrom),
		fmt.Sprintf("To: %s", to),
		fmt.Sprintf("Subject: %s", subject),
		"Auto-Submitted: auto-generated",
		"MIME-Version: 1.0",
		fmt.Sprintf("Content-Type: %s", contentType),
		"",
		content,
	}, "\r\n")

	addr := fmt.Sprintf("%s:%s", w.cfg.SMTPHost, w.cfg.SMTPPort)
	var auth smtp.Auth
	if w.cfg.SMTPUser != "" && w.cfg.SMTPPass != "" {
		auth = smtp.PlainAuth("", w.cfg.SMTPUser, w.cfg.SMTPPass, w.cfg.SMTPHost)
	}

	return smtp.SendMail(addr, auth, w.cfg.SMTPFrom, []string{to}, []byte(msg))
}
