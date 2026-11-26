package utils

import (
	"fmt"
	"html"
	"log/slog"
	"net/smtp"
	"strings"
	"time"

	"satpolpp-sulsel-api/internal/config"
)

// EmailNotice captures notification outcome for SIAP TERTIB reporter updates
type EmailNotice struct {
	Sent    bool   `json:"sent"`
	Mode    string `json:"mode"` // queued, smtp, simulated, skipped
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

// BuildTicketReceiptEmail generates both plaintext and clean HTML receipt bodies for email_outbox.
func BuildTicketReceiptEmail(siteURL, ticket, statusLabel, regencyName, disposedTo, actionNotes string) (subject, bodyText, bodyHTML string) {
	subject = fmt.Sprintf("[SIAP TERTIB] Pembaruan Tiket %s: %s", ticket, statusLabel)
	if statusLabel == "" || statusLabel == "MENUNGGU_VERIFIKASI" {
		subject = fmt.Sprintf("[SIAP TERTIB] Laporan Diterima - Tiket %s", ticket)
	}

	trackURL := fmt.Sprintf("%s/lacak?kode=%s", strings.TrimRight(siteURL, "/"), ticket)
	timestamp := time.Now().Format("02 Jan 2006 15:04 WITA")

	textLines := []string{
		"Yth. Pelapor SIAP TERTIB - Satpol PP Provinsi Sulawesi Selatan,",
		"",
		fmt.Sprintf("Nomor Tiket : %s", ticket),
		fmt.Sprintf("Status      : %s", statusLabel),
		fmt.Sprintf("Waktu       : %s", timestamp),
	}
	if regencyName != "" {
		textLines = append(textLines, fmt.Sprintf("Wilayah     : %s", regencyName))
	}
	if disposedTo != "" {
		textLines = append(textLines, fmt.Sprintf("Disposisi   : %s", disposedTo))
	}
	if actionNotes != "" {
		textLines = append(textLines, fmt.Sprintf("Catatan     : %s", actionNotes))
	}
	textLines = append(textLines, "", "Lacak perkembangan aduan Anda di:", trackURL, "", "Email ini dikirim otomatis (no-reply). Salam Praja Wibawa.")
	bodyText = strings.Join(textLines, "\r\n")

	esc := html.EscapeString
	extraRows := ""
	if regencyName != "" {
		extraRows += fmt.Sprintf(`<tr><td style="padding:6px 0;color:#64748b;font-size:13px;">Wilayah</td><td style="padding:6px 0;color:#0f172a;font-size:13px;font-weight:600;">%s</td></tr>`, esc(regencyName))
	}
	if disposedTo != "" {
		extraRows += fmt.Sprintf(`<tr><td style="padding:6px 0;color:#64748b;font-size:13px;">Disposisi</td><td style="padding:6px 0;color:#0f172a;font-size:13px;font-weight:600;">%s</td></tr>`, esc(disposedTo))
	}
	if actionNotes != "" {
		extraRows += fmt.Sprintf(`<tr><td style="padding:6px 0;color:#64748b;font-size:13px;">Catatan</td><td style="padding:6px 0;color:#0f172a;font-size:13px;">%s</td></tr>`, esc(actionNotes))
	}

	bodyHTML = fmt.Sprintf(`<!DOCTYPE html>
<html lang="id"><body style="margin:0;padding:24px;background:#f8fafc;font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,sans-serif;">
  <div style="max-width:560px;margin:0 auto;background:#ffffff;border:1px solid #e2e8f0;border-radius:16px;overflow:hidden;">
    <div style="background:#4A2E20;padding:20px 24px;color:#ffffff;">
      <div style="font-size:12px;letter-spacing:0.08em;text-transform:uppercase;opacity:0.85;">SIAP TERTIB • Satpol PP Prov. Sulsel</div>
      <h2 style="margin:6px 0 0;font-size:18px;font-weight:700;">Resi Tiket Pengaduan Masyarakat</h2>
    </div>
    <div style="padding:24px;">
      <p style="margin:0 0 16px;color:#334155;font-size:14px;line-height:1.6;">Laporan Anda telah tercatat dalam sistem kami. Simpan kode tiket berikut untuk memantau tindak lanjut petugas:</p>
      <div style="background:#fff7ed;border:1px solid #fed7aa;border-radius:12px;padding:14px 16px;margin-bottom:18px;">
        <div style="font-size:11px;color:#9a3412;text-transform:uppercase;letter-spacing:0.06em;">Nomor Tiket</div>
        <div style="font-size:20px;font-weight:800;color:#7c2d12;letter-spacing:0.04em;margin-top:2px;">%s</div>
      </div>
      <table style="width:100%%;border-collapse:collapse;">
        <tr><td style="padding:6px 0;color:#64748b;font-size:13px;width:110px;">Status</td><td style="padding:6px 0;color:#0f172a;font-size:13px;font-weight:700;">%s</td></tr>
        <tr><td style="padding:6px 0;color:#64748b;font-size:13px;">Waktu</td><td style="padding:6px 0;color:#0f172a;font-size:13px;">%s</td></tr>
        %s
      </table>
      <div style="margin-top:22px;">
        <a href="%s" style="display:inline-block;background:#F08300;color:#ffffff;text-decoration:none;font-weight:600;font-size:14px;padding:11px 20px;border-radius:10px;">Lacak Progres Tiket</a>
      </div>
    </div>
    <div style="padding:14px 24px;background:#f8fafc;border-top:1px solid #e2e8f0;font-size:12px;color:#64748b;">
      Email ini dikirim otomatis oleh sistem SIAP TERTIB (no-reply). Mohon tidak membalas email ini.
    </div>
  </div>
</body></html>`, esc(ticket), esc(statusLabel), esc(timestamp), extraRows, esc(trackURL))

	return subject, bodyText, bodyHTML
}

// SendComplaintEmail provides direct synchronous SMTP fallback if needed outside outbox.
func SendComplaintEmail(cfg *config.Config, toEmail, ticket, statusLabel, regencyName, disposedTo, actionNotes string) EmailNotice {
	to := strings.TrimSpace(toEmail)
	siteURL := "http://localhost:3000"
	if cfg != nil && cfg.PublicSiteURL != "" {
		siteURL = cfg.PublicSiteURL
	}
	subject, bodyText, _ := BuildTicketReceiptEmail(siteURL, ticket, statusLabel, regencyName, disposedTo, actionNotes)

	if to == "" {
		return EmailNotice{Sent: false, Mode: "skipped", Subject: subject, Message: "Pelapor tidak mencantumkan alamat email."}
	}
	if cfg == nil || strings.TrimSpace(cfg.SMTPHost) == "" {
		slog.Info("Simulasi notifikasi email SIAP TERTIB", "to", to, "ticket", ticket, "subject", subject)
		return EmailNotice{Sent: false, Mode: "simulated", To: to, Subject: subject, Message: "Konfigurasi SMTP_HOST belum diatur; notifikasi email disimulasikan."}
	}

	msg := strings.Join([]string{
		fmt.Sprintf("From: %s", cfg.SMTPFrom),
		fmt.Sprintf("To: %s", to),
		fmt.Sprintf("Subject: %s", subject),
		"MIME-Version: 1.0",
		"Content-Type: text/plain; charset=UTF-8",
		"",
		bodyText,
	}, "\r\n")

	addr := fmt.Sprintf("%s:%s", cfg.SMTPHost, cfg.SMTPPort)
	var auth smtp.Auth
	if cfg.SMTPUser != "" && cfg.SMTPPass != "" {
		auth = smtp.PlainAuth("", cfg.SMTPUser, cfg.SMTPPass, cfg.SMTPHost)
	}
	if err := smtp.SendMail(addr, auth, cfg.SMTPFrom, []string{to}, []byte(msg)); err != nil {
		return EmailNotice{Sent: false, Mode: "simulated", To: to, Subject: subject, Message: err.Error()}
	}
	return EmailNotice{Sent: true, Mode: "smtp", To: to, Subject: subject, Message: "Email notifikasi berhasil dikirim ke pelapor."}
}
