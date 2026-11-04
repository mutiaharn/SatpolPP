package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"satpolpp-sulsel-api/internal/adapters/middleware"
	"satpolpp-sulsel-api/internal/adapters/storage"
	"satpolpp-sulsel-api/internal/config"
	"satpolpp-sulsel-api/internal/core/domain"
	"satpolpp-sulsel-api/internal/utils"
)

// In-Memory Data Store with Seeder for Instant Development
type Store struct {
	mu          sync.RWMutex
	Users       map[string]*domain.User
	Articles    map[string]*domain.Article
	Documents   map[uint]*domain.PublicDocument
	Complaints  map[string]*domain.Complaint
	Settings    map[string]string
	Regencies   []domain.Regency
	AuditLogs   []domain.AuditLog
}

var store = &Store{
	Users:      make(map[string]*domain.User),
	Articles:   make(map[string]*domain.Article),
	Documents:  make(map[uint]*domain.PublicDocument),
	Complaints: make(map[string]*domain.Complaint),
	Settings:   make(map[string]string),
	Regencies:  make([]domain.Regency, 0),
	AuditLogs:  make([]domain.AuditLog, 0),
}

func initSeedData() {
	// 1. Seed 24 Kabupaten/Kota di Sulawesi Selatan
	regencyNames := []string{
		"Kota Makassar", "Kota Parepare", "Kota Palopo", "Kabupaten Gowa", "Kabupaten Maros",
		"Kabupaten Pangkajene dan Kepulauan", "Kabupaten Barru", "Kabupaten Bone", "Kabupaten Soppeng",
		"Kabupaten Wajo", "Kabupaten Sinjai", "Kabupaten Bulukumba", "Kabupaten Bantaeng",
		"Kabupaten Jeneponto", "Kabupaten Takalar", "Kabupaten Kepulauan Selayar", "Kabupaten Enrekang",
		"Kabupaten Pinrang", "Kabupaten Sidenreng Rappang", "Kabupaten Tana Toraja", "Kabupaten Toraja Utara",
		"Kabupaten Luwu", "Kabupaten Luwu Utara", "Kabupaten Luwu Timur",
	}
	for i, name := range regencyNames {
		store.Regencies = append(store.Regencies, domain.Regency{
			ID:   fmt.Sprintf("73%02d", i+1),
			Name: name,
		})
	}

	// 2. Seed Default Accounts (3-Tier RBAC)
	store.Users["superadmin"] = &domain.User{
		ID:        1,
		Name:      "Tim Pengembang Diskominfo-SP Sulsel",
		Username:  "superadmin",
		Email:     "diskominfo@sulselprov.go.id",
		Password:  "sulselprov2026!",
		Role:      domain.RoleSuperadmin,
		Bidang:    "Aplikasi & Informatika (Diskominfo)",
		IsActive:  true,
		CreatedAt: time.Now(),
	}
	store.Users["kasatpol"] = &domain.User{
		ID:        2,
		Name:      "Kasatpol PP Prov. Sulsel",
		Username:  "kasatpol",
		Email:     "satpol@sulselprov.go.id",
		Password:  "satpolpp2026!",
		Role:      domain.RoleAdmin,
		Bidang:    "Pimpinan",
		IsActive:  true,
		CreatedAt: time.Now(),
	}
	store.Users["stafftrantibum"] = &domain.User{
		ID:        3,
		Name:      "Staff Operasional Bidang Trantibum",
		Username:  "stafftrantibum",
		Email:     "trantibum@sulselprov.go.id",
		Password:  "trantibum2026!",
		Role:      domain.RoleStaff,
		Bidang:    "Bidang Ketenteraman dan Ketertiban Umum",
		IsActive:  true,
		CreatedAt: time.Now(),
	}

	// 3. Seed Anti-Hardcoded CMS Settings
	store.Settings["app_name"] = "Satpol PP & SIAP TERTIB Provinsi Sulawesi Selatan"
	store.Settings["agency_tagline"] = "Praja Wibawa: Mengayomi, Menertibkan, Melindungi Masyarakat Sulawesi Selatan"
	store.Settings["office_address"] = "Jl. Urip Sumoharjo No. 269, Kota Makassar, Sulawesi Selatan"
	store.Settings["call_center"] = "0411-453105"
	store.Settings["emergency_hotline"] = "112 / 0811-4141-112"
	store.Settings["email_contact"] = "satpol@sulselprov.go.id"
	store.Settings["kasatpol_name"] = "Andi Arwin Azis S.STP., M.M"
	store.Settings["kasatpol_quote"] = "Satpol PP Sulawesi Selatan senantiasa mengedepankan pendekatan humanis, persuasif, dan profesional dalam menegakkan Peraturan Daerah dan menjaga ketertiban umum demi kesejahteraan masyarakat."

	// 4. Seed Portal Articles
	store.Articles["perbaikan-rekonstruksi-ruas-paleteang-malimpung"] = &domain.Article{
		ID:          1,
		Title:       "Perbaikan Rekonstruksi Ruas Paleteang-Malimpung, Satpol PP Amankan Aset Wilayah",
		Slug:        "perbaikan-rekonstruksi-ruas-paleteang-malimpung",
		Summary:     "Satpol PP Prov Sulsel bersinergi dalam pengamanan aset pembangunan rekonstruksi ruas jalan provinsi di Pinrang.",
		Content:     "<p>Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan melaksanakan pendampingan dan penertiban pengamanan aset daerah...</p>",
		Thumbnail:   "assets/images/news/news-1.jpg",
		Category:    "Berita",
		AuthorName:  "Humas Satpol PP",
		ViewCount:   142,
		IsPublished: true,
		PublishedAt: time.Now().Add(-24 * time.Hour),
		CreatedAt:   time.Now().Add(-24 * time.Hour),
	}
	store.Articles["apel-persiapan-patroli-pasca-penertiban-aset"] = &domain.Article{
		ID:          2,
		Title:       "Apel Kesiapsiagaan Patroli Ketertiban Umum dan Penegakan Perda se-Sulawesi Selatan",
		Slug:        "apel-persiapan-patroli-pasca-penertiban-aset",
		Summary:     "Apel gabungan seluruh pleton operasional Satpol PP Prov Sulsel dalam rangka meningkatkan kesiapsiagaan patroli trantibum.",
		Content:     "<p>Guna menjaga ketenteraman dan ketertiban umum di wilayah Sulawesi Selatan, Satpol PP menyelenggarakan apel siaga...</p>",
		Thumbnail:   "assets/images/news/news-2.jpg",
		Category:    "Kegiatan",
		AuthorName:  "Sekretariat",
		ViewCount:   89,
		IsPublished: true,
		PublishedAt: time.Now().Add(-48 * time.Hour),
		CreatedAt:   time.Now().Add(-48 * time.Hour),
	}

	// 5. Seed Public Documents
	store.Documents[1] = &domain.PublicDocument{
		ID:          1,
		Title:       "Rencana Kerja (Renja) Satpol PP Prov. Sulsel Tahun 2026",
		DocNumber:   "050/112/SATPOL/2026",
		Category:    "renja",
		Year:        2026,
		FileURL:     "/docs/renja_2026.pdf",
		FileSize:    "2.4 MB",
		CreatedAt:   time.Now(),
	}
	store.Documents[2] = &domain.PublicDocument{
		ID:          2,
		Title:       "Laporan Kinerja Instansi Pemerintah (LAKIP) Tahun 2025",
		DocNumber:   "050/450/LAKIP/2025",
		Category:    "lakip",
		Year:        2025,
		FileURL:     "/docs/lakip_2025.pdf",
		FileSize:    "4.1 MB",
		CreatedAt:   time.Now(),
	}
	store.Documents[3] = &domain.PublicDocument{
		ID:          3,
		Title:       "Perda Prov. Sulsel No. 2 Tahun 2021 tentang Penyelenggaraan Trantibum",
		DocNumber:   "Perda No. 2/2021",
		Category:    "perda",
		Year:        2021,
		FileURL:     "/docs/perda_trantibum.pdf",
		FileSize:    "1.8 MB",
		CreatedAt:   time.Now(),
	}

	// 6. Seed Sample Complaint (SIAP TERTIB)
	sampleTicket := "TRB-202611-SAMPLE"
	store.Complaints[sampleTicket] = &domain.Complaint{
		ID:             1,
		TicketCode:     sampleTicket,
		IsAnonymous:    false,
		ReporterNIK:    "7371101234567890",
		ReporterName:   "Ahmad Fauzan",
		ReporterPhone:  "081234567890",
		ReporterEmail:  "ahmad.fauzan@gmail.com",
		ViolationType:  "Pelanggaran Ketertiban Umum (PKL Liar Menutup Trotoar)",
		RegencyID:      "7301",
		RegencyName:    "Kota Makassar",
		DistrictName:   "Panakkukang",
		VillageName:    "Panaikang",
		AddressDetail:  "Depan Halte Bus Kampus UMI, Jl. Urip Sumoharjo",
		Latitude:       -5.1432,
		Longitude:      119.4521,
		Description:    "Terdapat lapak liar yang memakan seluruh badan trotoar pejalan kaki dan menyebabkan kemacetan di jam pulang kerja.",
		IncidentDate:   time.Now().Add(-12 * time.Hour),
		Status:         domain.StatusInProgress,
		DisposedTo:     "Bidang Ketenteraman dan Ketertiban Umum",
		ActionNotes:    "Tim Patroli Regu 2 Satpol PP telah meninjau lokasi dan memberikan surat peringatan pertama kepada pemilik lapak.",
		CreatedAt:      time.Now().Add(-12 * time.Hour),
	}
}

// JSON Helper
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg := config.LoadConfig()
	initSeedData()

	mux := http.NewServeMux()

	// 1. Health Check
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{
			"status":    "healthy",
			"app":       "Satpol PP Sulsel & SIAP TERTIB Unified API",
			"version":   "1.0.0",
			"timestamp": time.Now().Format(time.RFC3339),
		})
	})

	// 2. Auth Login (3 Roles: SUPERADMIN, ADMIN, STAFF)
	mux.HandleFunc("POST /api/v1/auth/login", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
			return
		}

		store.mu.RLock()
		user, exists := store.Users[req.Username]
		store.mu.RUnlock()

		if !exists || user.Password != req.Password {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "Username atau kata sandi tidak valid"})
			return
		}

		// In production, sign JWT token with cfg.JWTSecret
		token := fmt.Sprintf("jwt-mock-%s-%d", user.Role, user.ID)

		writeJSON(w, http.StatusOK, map[string]any{
			"message": "Login berhasil",
			"token":   token,
			"user": map[string]any{
				"id":       user.ID,
				"name":     user.Name,
				"username": user.Username,
				"role":     user.Role,
				"bidang":   user.Bidang,
			},
		})
	})

	// 3. Portal Public Settings (Anti-Hardcoded CMS)
	mux.HandleFunc("GET /api/v1/portal/settings", func(w http.ResponseWriter, r *http.Request) {
		store.mu.RLock()
		defer store.mu.RUnlock()
		writeJSON(w, http.StatusOK, map[string]any{
			"success":  true,
			"settings": store.Settings,
		})
	})

	// 4. Portal Articles (News)
	mux.HandleFunc("GET /api/v1/portal/articles", func(w http.ResponseWriter, r *http.Request) {
		store.mu.RLock()
		defer store.mu.RUnlock()
		list := make([]*domain.Article, 0, len(store.Articles))
		for _, a := range store.Articles {
			list = append(list, a)
		}
		writeJSON(w, http.StatusOK, map[string]any{
			"success": true,
			"data":    list,
		})
	})

	mux.HandleFunc("GET /api/v1/portal/articles/{slug}", func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		store.mu.RLock()
		article, exists := store.Articles[slug]
		store.mu.RUnlock()

		if !exists {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "Artikel tidak ditemukan"})
			return
		}
		article.ViewCount++
		writeJSON(w, http.StatusOK, map[string]any{
			"success": true,
			"data":    article,
		})
	})

	// 5. Portal Documents (Renja, DPA, LAKIP)
	mux.HandleFunc("GET /api/v1/portal/documents", func(w http.ResponseWriter, r *http.Request) {
		category := r.URL.Query().Get("category")
		store.mu.RLock()
		defer store.mu.RUnlock()

		list := make([]*domain.PublicDocument, 0)
		for _, doc := range store.Documents {
			if category == "" || strings.EqualFold(doc.Category, category) {
				list = append(list, doc)
			}
		}
		writeJSON(w, http.StatusOK, map[string]any{
			"success": true,
			"data":    list,
		})
	})

	// 6. Master Data Wilayah (24 Kabupaten/Kota se-Sulsel)
	mux.HandleFunc("GET /api/v1/portal/regencies", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{
			"success": true,
			"data":    store.Regencies,
		})
	})

	// 7. SIAP TERTIB: Kirim Laporan Publik
	mux.HandleFunc("POST /api/v1/siaptertib/lapor", func(w http.ResponseWriter, r *http.Request) {
		var req domain.Complaint
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Format data tidak valid"})
			return
		}

		if req.ViolationType == "" || req.AddressDetail == "" {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Jenis pelanggaran dan detail alamat wajib diisi"})
			return
		}

		ticket := utils.GenerateTicketCode()
		req.ID = uint(len(store.Complaints) + 1)
		req.TicketCode = ticket
		req.Status = domain.StatusPending
		req.CreatedAt = time.Now()
		req.UpdatedAt = time.Now()

		store.mu.Lock()
		store.Complaints[ticket] = &req
		store.mu.Unlock()

		slog.Info("Laporan aduan baru diterima", "ticket", ticket, "anonim", req.IsAnonymous, "jenis", req.ViolationType)

		writeJSON(w, http.StatusCreated, map[string]any{
			"success":     true,
			"message":     "Laporan aduan Anda berhasil diterima dan sedang diverifikasi petugas.",
			"ticket_code": ticket,
		})
	})

	// 8. SIAP TERTIB: Lacak Tiket Aduan Publik
	mux.HandleFunc("GET /api/v1/siaptertib/lacak/{ticket}", func(w http.ResponseWriter, r *http.Request) {
		ticket := r.PathValue("ticket")
		store.mu.RLock()
		comp, exists := store.Complaints[ticket]
		store.mu.RUnlock()

		if !exists {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "Nomor tiket aduan tidak ditemukan"})
			return
		}

		// Mask sensitive identity if report is marked anonymous
		displayName := comp.ReporterName
		displayNIK := comp.ReporterNIK
		displayPhone := comp.ReporterPhone
		if comp.IsAnonymous {
			displayName, displayNIK, displayPhone = utils.MaskReporterIdentity(comp.ReporterName, comp.ReporterNIK, comp.ReporterPhone)
		}

		writeJSON(w, http.StatusOK, map[string]any{
			"success": true,
			"data": map[string]any{
				"ticket_code":    comp.TicketCode,
				"status":         comp.Status,
				"is_anonymous":   comp.IsAnonymous,
				"reporter_name":  displayName,
				"reporter_nik":   displayNIK,
				"reporter_phone": displayPhone,
				"violation_type": comp.ViolationType,
				"regency_name":   comp.RegencyName,
				"district_name":  comp.DistrictName,
				"address_detail": comp.AddressDetail,
				"latitude":       comp.Latitude,
				"longitude":      comp.Longitude,
				"description":    comp.Description,
				"disposed_to":    comp.DisposedTo,
				"action_notes":   comp.ActionNotes,
				"created_at":     comp.CreatedAt,
			},
		})
	})

	// 9. Admin Satpol PP: List All Complaints
	mux.HandleFunc("GET /api/v1/admin/complaints", func(w http.ResponseWriter, r *http.Request) {
		store.mu.RLock()
		defer store.mu.RUnlock()

		list := make([]*domain.Complaint, 0, len(store.Complaints))
		for _, c := range store.Complaints {
			list = append(list, c)
		}

		writeJSON(w, http.StatusOK, map[string]any{
			"success": true,
			"total":   len(list),
			"data":    list,
		})
	})

	// 10. Admin Satpol PP: Update Complaint Status & Disposisi
	mux.HandleFunc("PUT /api/v1/admin/complaints/{ticket}/status", func(w http.ResponseWriter, r *http.Request) {
		ticket := r.PathValue("ticket")
		var req struct {
			Status      domain.ComplaintStatus `json:"status"`
			DisposedTo  string                 `json:"disposed_to"`
			ActionNotes string                 `json:"action_notes"`
			OfficerName string                 `json:"officer_name"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid body"})
			return
		}

		store.mu.Lock()
		comp, exists := store.Complaints[ticket]
		if !exists {
			store.mu.Unlock()
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "Tiket tidak ditemukan"})
			return
		}

		if req.Status != "" {
			comp.Status = req.Status
		}
		if req.DisposedTo != "" {
			comp.DisposedTo = req.DisposedTo
		}
		if req.ActionNotes != "" {
			comp.ActionNotes = req.ActionNotes
		}
		if req.OfficerName != "" {
			comp.OfficerName = req.OfficerName
		}
		comp.UpdatedAt = time.Now()
		store.mu.Unlock()

		writeJSON(w, http.StatusOK, map[string]any{
			"success": true,
			"message": "Status penanganan aduan berhasil diperbarui",
			"data":    comp,
		})
	})

	// 11. Superadmin Diskominfo: Dynamic CMS Settings Update
	mux.HandleFunc("PUT /api/v1/superadmin/settings", func(w http.ResponseWriter, r *http.Request) {
		var updates map[string]string
		if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid body"})
			return
		}

		store.mu.Lock()
		for k, v := range updates {
			store.Settings[k] = v
		}
		store.mu.Unlock()

		writeJSON(w, http.StatusOK, map[string]any{
			"success": true,
			"message": "Pengaturan CMS berhasil diperbarui",
			"data":    store.Settings,
		})
	})

	// 12. Strict Image Evidence Upload (MIME Magic-Bytes Validated & 5MB Max)
	imageStorage, err := storage.NewImageStorage("./uploads/evidence")
	if err != nil {
		slog.Error("Gagal inisialisasi storage gambar", "error", err)
	}

	mux.HandleFunc("POST /api/v1/upload/evidence", func(w http.ResponseWriter, r *http.Request) {
		// Limit body size to 6 MB (5MB file + multipart form overhead)
		r.Body = http.MaxBytesReader(w, r.Body, 6*1024*1024)
		if err := r.ParseMultipartForm(6 * 1024 * 1024); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{
				"success": false,
				"error":   "Ukuran berkas melebihi batas maksimum 5 MB",
				"code":    "FILE_TOO_LARGE",
			})
			return
		}

		file, header, err := r.FormFile("foto")
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{
				"success": false,
				"error":   "Berkas bukti foto tidak ditemukan pada form (field 'foto')",
				"code":    "MISSING_FILE",
			})
			return
		}
		defer file.Close()

		savedURL, err := imageStorage.SaveUploadedImage(header)
		if err != nil {
			status := http.StatusBadRequest
			if strings.Contains(err.Error(), "5 MB") {
				status = http.StatusRequestEntityTooLarge
			}
			writeJSON(w, status, map[string]any{
				"success": false,
				"error":   err.Error(),
				"code":    "UPLOAD_FAILED",
			})
			return
		}

		slog.Info("Bukti foto pelanggaran berhasil diunggah", "path", savedURL, "size", header.Size)

		writeJSON(w, http.StatusCreated, map[string]any{
			"success":   true,
			"message":   "Bukti foto berhasil diverifikasi dan disimpan",
			"image_url": savedURL,
		})
	})

	// Serve Static Uploads
	fs := http.FileServer(http.Dir("./uploads"))
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", fs))

	// Frontier Defense-in-Depth Middleware Chain:
	// 1. Panic Recovery (Outermost)
	// 2. HTTP Security Headers (CSP, HSTS, X-Content-Type-Options, etc.)
	// 3. IP Token-Bucket Rate Limiter (60 req/min)
	// 4. Configurable CORS (Vite + Gov Domains)
	rateLimiter := middleware.NewIPRateLimiter(60, time.Second)
	corsMiddleware := middleware.CORS([]string{
		"*", // Allow development origins
		"http://localhost:5173",
		"http://localhost:3000",
		"http://127.0.0.1:5173",
		"https://satpolpp.sulselprov.go.id",
		"https://siaptertib.sulselprov.go.id",
	})

	secureHandler := middleware.PanicRecovery(
		middleware.SecurityHeaders(
			rateLimiter.Middleware(
				corsMiddleware(mux),
			),
		),
	)

	addr := ":" + cfg.AppPort
	slog.Info("Server Backend Golang Satpol PP & SIAP TERTIB berjalan", "port", cfg.AppPort, "env", cfg.AppEnv)
	slog.Info("Health check endpoint: http://localhost:" + cfg.AppPort + "/health")
	if err := http.ListenAndServe(addr, secureHandler); err != nil {
		slog.Error("Gagal menjalankan server", "error", err)
	}
}
