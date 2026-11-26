package postgres

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"satpolpp-sulsel-api/internal/contracts"
	"satpolpp-sulsel-api/internal/core/domain"
	"satpolpp-sulsel-api/internal/utils"
)

// SeedIfEmpty populates initial master data idempotently on first boot.
func (d *DB) SeedIfEmpty(ctx context.Context) error {
	// 1. Always ensure 24 Regencies exist (idempotent)
	for _, r := range contracts.Shared.Regencies {
		_, err := d.Pool.Exec(ctx,
			`INSERT INTO regencies (id, name) VALUES ($1, $2) ON CONFLICT (id) DO NOTHING`,
			r.ID, r.Name,
		)
		if err != nil {
			return err
		}
	}

	var userCount int
	if err := d.Pool.QueryRow(ctx, `SELECT COUNT(*) FROM users`).Scan(&userCount); err != nil {
		return err
	}
	if userCount > 0 {
		return nil // Already seeded
	}

	slog.Info("Menjalankan seeder awal database PostgreSQL...")

	return d.WithTx(ctx, func(tx pgx.Tx) error {
		// 2. Seed 3 Default Accounts with bcrypt hash
		accounts := []struct {
			name, username, email, pass string
			role                        domain.Role
			bidang                      string
		}{
			{"Tim Pengembang Diskominfo-SP Sulsel", "superadmin", "diskominfo@sulselprov.go.id", "sulselprov2026!", domain.RoleSuperadmin, "Aplikasi & Informatika (Diskominfo)"},
			{"Kasatpol PP Prov. Sulsel", "kasatpol", "satpol@sulselprov.go.id", "satpolpp2026!", domain.RoleAdmin, "Pimpinan"},
			{"Staff Operasional Bidang Trantibum", "stafftrantibum", "trantibum@sulselprov.go.id", "trantibum2026!", domain.RoleStaff, "Bidang Ketenteraman dan Ketertiban Umum"},
		}

		var superadminID uint = 1
		for i, acc := range accounts {
			hashed, err := utils.HashPassword(acc.pass)
			if err != nil {
				return err
			}
			var uid uint
			err = tx.QueryRow(ctx, `
				INSERT INTO users (name, email, username, password, role, bidang, is_active, updated_at)
				VALUES ($1, $2, $3, $4, $5::"Role", $6, true, NOW())
				ON CONFLICT (username) DO UPDATE SET name = EXCLUDED.name
				RETURNING id
			`, acc.name, acc.email, acc.username, hashed, string(acc.role), acc.bidang).Scan(&uid)
			if err != nil {
				return err
			}
			if i == 0 {
				superadminID = uid
			}
		}

		// 3. Seed Site Settings
		settings := map[string]string{
			"app_name":          "Satpol PP & SIAP TERTIB Provinsi Sulawesi Selatan",
			"agency_tagline":    "Praja Wibawa: Mengayomi, Menertibkan, Melindungi Masyarakat Sulawesi Selatan",
			"office_address":    "Jl. Urip Sumoharjo No. 269, Kota Makassar, Sulawesi Selatan",
			"call_center":       "0411-453105",
			"emergency_hotline": "112 / 0811-4141-112",
			"email_contact":     "satpol@sulselprov.go.id",
			"kasatpol_name":     "Andi Arwin Azis S.STP., M.M",
			"kasatpol_quote":    "Satpol PP Sulawesi Selatan senantiasa mengedepankan pendekatan humanis, persuasif, dan profesional dalam menegakkan Peraturan Daerah dan menjaga ketertiban umum demi kesejahteraan masyarakat.",
		}
		for k, v := range settings {
			if _, err := tx.Exec(ctx, `
				INSERT INTO site_settings (key, value, "group", updated_at)
				VALUES ($1, $2, 'general', NOW()) ON CONFLICT (key) DO NOTHING
			`, k, v); err != nil {
				return err
			}
		}

		// 4. Seed Sample Articles
		articles := []struct {
			title, slug, summary, content, thumb, cat string
			views                                     int
		}{
			{
				"Perbaikan Rekonstruksi Ruas Paleteang-Malimpung, Satpol PP Amankan Aset Wilayah",
				"perbaikan-rekonstruksi-ruas-paleteang-malimpung",
				"Satpol PP Prov Sulsel bersinergi dalam pengamanan aset pembangunan rekonstruksi ruas jalan provinsi di Pinrang.",
				"<p>Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan melaksanakan pendampingan dan penertiban pengamanan aset daerah pada proyek rekonstruksi ruas jalan Paleteang-Malimpung.</p>",
				"/uploads/foto_berita/1768182343WhatsApp Image 2026-01-12 at 08.08.43.jpeg",
				"Kegiatan",
				320,
			},
			{
				"Rapat Internal Danki dan Danton Operasional Satpol PP Sulsel",
				"rapat-internal-danki-dan-danton-ops",
				"Konsolidasi jajaran pimpinan kompi dan pleton operasional guna meningkatkan efektivitas patroli penegakan Perda.",
				"<p>Kepala Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan memimpin rapat koordinasi berkala bersama seluruh Danki dan Danton operasional.</p>",
				"/uploads/foto_berita/1758537465AF0666BD-2416-4D0A-9101-E1BBB12A8FA9.png",
				"Kegiatan",
				185,
			},
			{
				"Apel Kesiapsiagaan Patroli Ketertiban Umum Pasca Penertiban Aset Pemprov",
				"apel-persiapan-patroli-pasca-penertiban-aset",
				"Pemeriksaan kesiapan personel dan armada patroli Satpol PP dalam mengawal aset tanah dan bangunan milik Pemprov Sulsel.",
				"<p>Guna menjaga ketenteraman dan ketertiban umum pasca dilaksanakannya penertiban aset daerah di beberapa titik strategis, Satpol PP Sulsel menggelar apel siaga.</p>",
				"/uploads/foto_berita/175853755659BEED7D-9D9F-4A4E-A69E-DCC8EE0A2A50.png",
				"Penertiban Aset",
				245,
			},
		}
		for _, a := range articles {
			if _, err := tx.Exec(ctx, `
				INSERT INTO articles (title, slug, summary, content, thumbnail, category, author_id, view_count, is_published, published_at, updated_at)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, true, NOW(), NOW())
				ON CONFLICT (slug) DO NOTHING
			`, a.title, a.slug, a.summary, a.content, a.thumb, a.cat, superadminID, a.views); err != nil {
				return err
			}
		}

		// 5. Seed Public Documents
		docs := []struct {
			title, num, cat, url, size string
			year, dl                   int
		}{
			{"Rencana Kerja (RENJA) Satpol PP Prov. Sulsel Tahun 2026", "001/RENJA/SATPOLPP/2026", "renja", "/uploads/dokumen/renja_2026.pdf", "2.4 MB", 2026, 412},
			{"Peraturan Daerah Prov. Sulsel No. 2 Tahun 2021 tentang Trantibum", "PERDA-02/2021", "perda", "/uploads/dokumen/perda_trantibum_2021.pdf", "1.8 MB", 2021, 1289},
			{"Laporan Kinerja Instansi Pemerintah (LAKIP) Tahun 2025", "004/LAKIP/SATPOLPP/2025", "lakip", "/uploads/dokumen/lakip_2025.pdf", "3.1 MB", 2025, 290},
		}
		for _, doc := range docs {
			if _, err := tx.Exec(ctx, `
				INSERT INTO public_documents (title, doc_number, category, year, file_url, file_size, download_count, updated_at)
				VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())
			`, doc.title, doc.num, doc.cat, doc.year, doc.url, doc.size, doc.dl); err != nil {
				return err
			}
		}

		// 6. Seed SOTK Organization Nodes
		subSekretariat, _ := json.Marshal([]string{"Sub Bagian Program & Pelaporan", "Sub Bagian Keuangan", "Sub Bagian Umum & Kepegawaian"})
		subTrantibum, _ := json.Marshal([]string{"Seksi Operasi dan Pengendalian", "Seksi Kerjasama"})
		subPerda, _ := json.Marshal([]string{"Seksi Penyelidikan dan Penyidikan", "Seksi Pembinaan dan Pengawasan"})

		orgs := []struct {
			code, parent, title, name, nip, rank, desc string
			subs                                       []byte
			level, order                               int
		}{
			{"KASATPOL", "", "Kepala Satuan Polisi Pamong Praja", "Andi Arwin Azis S.STP., M.M", "19780412 199810 1 001", "Pembina Utama Muda (IV/c)", "Memimpin pelaksanaan urusan pemerintahan bidang ketenteraman dan ketertiban umum serta penegakan Perda.", []byte("[]"), 1, 1},
			{"SEKRETARIAT", "KASATPOL", "Sekretaris Satpol PP", "Drs. H. Muh. Idris, M.Si", "19720817 199303 1 004", "Pembina Tk. I (IV/b)", "Mengelola administrasi kepegawaian, keuangan, perencanaan, dan kehumasan.", subSekretariat, 2, 1},
			{"BID-TRANTIBUM", "KASATPOL", "Kepala Bidang Ketenteraman & Ketertiban Umum", "Andi Bau Sawa, S.Sos., M.M", "19750520 200003 1 002", "Pembina (IV/a)", "Pengendalian operasi patroli dan penertiban ketenteraman umum.", subTrantibum, 2, 2},
			{"BID-PERDA", "KASATPOL", "Kepala Bidang Penegakan Perda & Perkada", "Irwan Syah, S.H., M.H", "19800115 200502 1 003", "Pembina (IV/a)", "Penyelidikan, penyidikan PPNS, dan pengawasan kepatuhan Perda.", subPerda, 2, 3},
		}
		for _, o := range orgs {
			var pCode *string
			if o.parent != "" {
				p := o.parent
				pCode = &p
			}
			if _, err := tx.Exec(ctx, `
				INSERT INTO organization_nodes (org_code, parent_code, position_title, official_name, nip, rank, description, sub_units, level, sort_order, updated_at)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8::jsonb, $9, $10, NOW())
				ON CONFLICT (org_code) DO NOTHING
			`, o.code, pCode, o.title, o.name, o.nip, o.rank, o.desc, string(o.subs), o.level, o.order); err != nil {
				return err
			}
		}

		return nil
	})
}
