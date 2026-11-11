package utils

import (
	"fmt"
	"log/slog"
	"net/smtp"
	"strings"

	"satpolpp-sulsel-api/internal/config"
)

// EmailNotice captures notification outcome for SIAP TERTIB reporter updates
type EmailNotice struct {
	Sent    bool   `json:"sent"`
	Mode    string `json:"mode"` // smtp, simulated, skipped
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

// SendComplaintEmail dispatches complaint receipt or status update email using net/smtp when configured.
func SendComplaintEmail(cfg *config.Config, toEmail, ticket, statusLabel, regencyName, disposedTo, actionNotes string) EmailNotice {
	to := strings.TrimSpace(toEmail)
	subject := fmt.Sprintf("[SIAP TERTIB] Pembaruan Tiket %s: %s", ticket, statusLabel)
	if statusLabel == "" || statusLabel == "MENUNGGU_VERIFIKASI" {
		subject = fmt.Sprintf("[SIAP TERTIB] Laporan Diterima - Tiket %s", ticket)
	}

	if to == "" {
		return EmailNotice{
			Sent:    false,
			Mode:    "skipped",
			To:      "",
			Subject: subject,
			Message: "Pelapor tidak mencantumkan alamat email.",
		}
	}

	if cfg == nil || strings.TrimSpace(cfg.SMTPHost) == "" {
		slog.Info("Simulasi notifikasi email SIAP TERTIB", "to", to, "ticket", ticket, "subject", subject)
		return EmailNotice{
			Sent:    false,
			Mode:    "simulated",
			To:      to,
			Subject: subject,
			Message: "Konfigurasi SMTP_HOST belum diatur; notifikasi email disimulasikan di mode pengembangan.",
		}
	}

	bodyLines := []string{
		"Yth. Pelapor SIAP TERTIB - Satpol PP Prov. Sulawesi Selatan,",
		"",
		fmt.Sprintf("Nomor Tiket : %s", ticket),
		fmt.Sprintf("Status      : %s", statusLabel),
	}
	if regencyName != "" {
		bodyLines = append(bodyLines, fmt.Sprintf("Wilayah     : %s", regencyName))
	}
	if disposedTo != "" {
		bodyLines = append(bodyLines, fmt.Sprintf("Disposisi   : %s", disposedTo))
	}
	if actionNotes != "" {
		bodyLines = append(bodyLines, fmt.Sprintf("Catatan     : %s", actionNotes))
	}
	bodyLines = append(bodyLines, "", "Lacak perkembangan aduan Anda melalui menu Lacak Tiket pada portal resmi Satpol PP Prov. Sulsel.", "", "Salam Praja Wibawa.")

	msg := strings.Join([]string{
		fmt.Sprintf("From: %s", cfg.SMTPFrom),
		fmt.Sprintf("To: %s", to),
		fmt.Sprintf("Subject: %s", subject),
		"MIME-Version: 1.0",
		"Content-Type: text/plain; charset=UTF-8",
		"",
		strings.Join(bodyLines, "\r\n"),
	}, "\r\n")

	addr := fmt.Sprintf("%s:%s", cfg.SMTPHost, cfg.SMTPPort)
	var auth smtp.Auth
	if cfg.SMTPUser != "" && cfg.SMTPPass != "" {
		auth = smtp.PlainAuth("", cfg.SMTPUser, cfg.SMTPPass, cfg.SMTPHost)
	}

	if err := smtp.SendMail(addr, auth, cfg.SMTPFrom, []string{to}, []byte(msg)); err != nil {
		slog.Warn("Gagal mengirim email SMTP", "error", err, "to", to, "ticket", ticket)
		return EmailNotice{
			Sent:    false,
			Mode:    "simulated",
			To:      to,
			Subject: subject,
			Message: fmt.Sprintf("Pengiriman SMTP gagal (%v); dicatat sebagai simulasi.", err),
		}
	}

	return EmailNotice{
		Sent:    true,
		Mode:    "smtp",
		To:      to,
		Subject: subject,
		Message: "Email notifikasi berhasil dikirim ke pelapor.",
	}
}
