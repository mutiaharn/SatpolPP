package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"sort"
	"strconv"
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
	mu            sync.RWMutex
	Users         map[string]*domain.User
	Articles      map[string]*domain.Article
	Documents     map[uint]*domain.PublicDocument
	Complaints    map[string]*domain.Complaint
	Organizations map[uint]*domain.OrganizationNode
	Settings      map[string]string
	Regencies     []domain.Regency
	AuditLogs     []domain.AuditLog
	NextUserID    uint
	NextOrgID     uint
}

var store = &Store{
	Users:         make(map[string]*domain.User),
	Articles:      make(map[string]*domain.Article),
	Documents:     make(map[uint]*domain.PublicDocument),
	Complaints:    make(map[string]*domain.Complaint),
	Organizations: make(map[uint]*domain.OrganizationNode),
	Settings:      make(map[string]string),
	Regencies:     make([]domain.Regency, 0),
	AuditLogs:     make([]domain.AuditLog, 0),
	NextUserID:    4,
	NextOrgID:     6,
}

func strPtr(s string) *string {
	return &s
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
		ID:        1,
		Title:     "Rencana Kerja (Renja) Satpol PP Prov. Sulsel Tahun 2026",
		DocNumber: "050/112/SATPOL/2026",
		Category:  "renja",
		Year:      2026,
		FileURL:   "/docs/renja_2026.pdf",
		FileSize:  "2.4 MB",
		CreatedAt: time.Now(),
	}
	store.Documents[2] = &domain.PublicDocument{
		ID:        2,
		Title:     "Laporan Kinerja Instansi Pemerintah (LAKIP) Tahun 2025",
		DocNumber: "050/450/LAKIP/2025",
		Category:  "lakip",
		Year:      2025,
		FileURL:   "/docs/lakip_2025.pdf",
		FileSize:  "4.1 MB",
		CreatedAt: time.Now(),
	}
	store.Documents[3] = &domain.PublicDocument{
		ID:        3,
		Title:     "Perda Prov. Sulsel No. 2 Tahun 2021 tentang Penyelenggaraan Trantibum",
		DocNumber: "Perda No. 2/2021",
		Category:  "perda",
		Year:      2021,
		FileURL:   "/docs/perda_trantibum.pdf",
		FileSize:  "1.8 MB",
		CreatedAt: time.Now(),
	}

	// 6. Seed Sample Complaint (SIAP TERTIB)
	sampleTicket := "TRB-202611-SAMPLE"
	store.Complaints[sampleTicket] = &domain.Complaint{
		ID:            1,
		TicketCode:    sampleTicket,
		IsAnonymous:   false,
		ReporterNIK:   "7371101234567890",
		ReporterName:  "Ahmad Fauzan",
		ReporterPhone: "081234567890",
		ReporterEmail: "ahmad.fauzan@gmail.com",
		ViolationType: "Pelanggaran Ketertiban Umum (PKL Liar Menutup Trotoar)",
		RegencyID:     "7301",
		RegencyName:   "Kota Makassar",
		DistrictName:  "Panakkukang",
		VillageName:   "Panaikang",
		AddressDetail: "Depan Halte Bus Kampus UMI, Jl. Urip Sumoharjo",
		Latitude:      -5.1432,
		Longitude:     119.4521,
		Description:   "Terdapat lapak liar yang memakan seluruh badan trotoar pejalan kaki dan menyebabkan kemacetan di jam pulang kerja.",
		IncidentDate:  time.Now().Add(-12 * time.Hour),
		Status:        domain.StatusInProgress,
		DisposedTo:    "Bidang Ketenteraman dan Ketertiban Umum",
		ActionNotes:   "Tim Patroli Regu 2 Satpol PP telah meninjau lokasi dan memberikan surat peringatan pertama kepada pemilik lapak.",
		CreatedAt:     time.Now().Add(-12 * time.Hour),
	}

	// 7. Seed Dynamic Organization Nodes (SOTK)
	store.Organizations[1] = &domain.OrganizationNode{
		ID:            1,
		OrgCode:       "ORG-KST-01",
		ParentCode:    nil,
		PositionTitle: "Kepala Satuan Polisi Pamong Praja",
		OfficialName:  "Andi Arwin Azis S.STP., M.M",
		Rank:          "Eselon IIa",
		Description:   "Memimpin perumusan kebijakan, pengawasan operasional trantibum di 24 kabupaten/kota, serta koordinasi ketenteraman wilayah bersama Forkopimda Sulawesi Selatan.",
		SubUnits:      []string{},
		PhotoURL:      "/uploads/foto_pimpinan/pimpinan_cutout_hd.png",
		Level:         1,
		SortOrder:     1,
	}
	store.Organizations[2] = &domain.OrganizationNode{
		ID:            2,
		OrgCode:       "ORG-SEK-01",
		ParentCode:    strPtr("ORG-KST-01"),
		PositionTitle: "Sekretariat",
		OfficialName:  "Sekretaris Satpol PP Sulsel",
		Rank:          "Eselon IIIa",
		Description:   "Mengoordinasikan urusan perencanaan program, pengelolaan keuangan anggaran, dan tata usaha kepegawaian.",
		SubUnits:      []string{"Sub Bagian Umum, Kepegawaian dan Hukum", "Sub Bagian Program dan Pelaporan", "Sub Bagian Keuangan dan Pengelolaan Aset"},
		Level:         2,
		SortOrder:     1,
	}
	store.Organizations[3] = &domain.OrganizationNode{
		ID:            3,
		OrgCode:       "ORG-BID-01",
		ParentCode:    strPtr("ORG-KST-01"),
		PositionTitle: "Bidang Penegakan Peraturan Daerah",
		OfficialName:  "Kabid Penegakan Perda",
		Rank:          "Eselon IIIb",
		Description:   "Memimpin penindakan yustisial dan non-yustisial serta penyidikan pelanggaran peraturan daerah oleh PPNS.",
		SubUnits:      []string{"Seksi Pengawasan dan Pembinaan Penegakan Perda", "Seksi Penindakan, Penyidikan dan PPNS Daerah"},
		Level:         2,
		SortOrder:     2,
	}
	store.Organizations[4] = &domain.OrganizationNode{
		ID:            4,
		OrgCode:       "ORG-BID-02",
		ParentCode:    strPtr("ORG-KST-01"),
		PositionTitle: "Bidang Ketenteraman dan Ketertiban Umum",
		OfficialName:  "Kabid Trantibum",
		Rank:          "Eselon IIIb",
		Description:   "Mengendalikan operasi patroli ketertiban umum, penertiban fasilitas publik, dan pengawalan pejabat negara.",
		SubUnits:      []string{"Seksi Operasi dan Pengendalian Ketertiban", "Seksi Pengamanan, Pengawalan dan Patroli Wilayah"},
		Level:         2,
		SortOrder:     3,
	}
	store.Organizations[5] = &domain.OrganizationNode{
		ID:            5,
		OrgCode:       "ORG-BID-03",
		ParentCode:    strPtr("ORG-KST-01"),
		PositionTitle: "Bidang Pelindungan Masyarakat dan Damkar",
		OfficialName:  "Kabid Linmas & Damkar",
		Rank:          "Eselon IIIb",
		Description:   "Menyelenggarakan pembinaan teknis anggota Satlinmas dan koordinasi penanganan kedaruratan kebakaran daerah.",
		SubUnits:      []string{"Seksi Satuan Pelindungan Masyarakat (Satlinmas)", "Seksi Kesiapsiagaan, Kebakaran dan Bencana"},
		Level:         2,
		SortOrder:     4,
	}
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func sortedOrganizations() []*domain.OrganizationNode {
	list := make([]*domain.OrganizationNode, 0, len(store.Organizations))
	for _, node := range store.Organizations {
		list = append(list, node)
	}
	sort.Slice(list, func(i, j int) bool {
		if list[i].Level != list[j].Level {
			return list[i].Level < list[j].Level
		}
		return list[i].SortOrder < list[j].SortOrder
	})
	return list
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
			"version":   "1.1.0",
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
		user, exists := store.Users[strings.ToLower(strings.TrimSpace(req.Username))]
		store.mu.RUnlock()

		if !exists || user.Password != req.Password {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "Username atau kata sandi tidak valid"})
			return
		}
		if !user.IsActive {
			writeJSON(w, http.StatusForbidden, map[string]string{"error": "Akun ini sedang dinonaktifkan oleh Superadmin"})
			return
		}

		token := fmt.Sprintf("jwt-mock-%s-%d", user.Role, user.ID)
		writeJSON(w, http.StatusOK, map[string]any{
			"message": "Login berhasil",
			"token":   token,
			"user": map[string]any{
				"id":        user.ID,
				"name":      user.Name,
				"username":  user.Username,
				"email":     user.Email,
				"role":      user.Role,
				"bidang":    user.Bidang,
				"is_active": user.IsActive,
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

	// 7. Dynamic Organization Structure (Public Read)
	mux.HandleFunc("GET /api/v1/portal/organization", func(w http.ResponseWriter, r *http.Request) {
		store.mu.RLock()
		defer store.mu.RUnlock()
		writeJSON(w, http.StatusOK, map[string]any{
			"success": true,
			"data":    sortedOrganizations(),
		})
	})

	// 8. Admin Organization Structure CRUD
	mux.Handle("POST /api/v1/admin/organization", middleware.RequireAuth(domain.RoleStaff, domain.RoleAdmin, domain.RoleSuperadmin)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req domain.OrganizationNode
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Format payload struktur tidak valid"})
			return
		}
		req.OrgCode = strings.ToUpper(strings.TrimSpace(req.OrgCode))
		if req.OrgCode == "" || strings.TrimSpace(req.PositionTitle) == "" {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "ID Khusus dan nama jabatan wajib diisi"})
			return
		}

		store.mu.Lock()
		for _, existing := range store.Organizations {
			if strings.EqualFold(existing.OrgCode, req.OrgCode) {
				store.mu.Unlock()
				writeJSON(w, http.StatusConflict, map[string]string{"error": "ID Khusus unit sudah terdaftar"})
				return
			}
		}
		req.ID = store.NextOrgID
		store.NextOrgID++
		if req.SubUnits == nil {
			req.SubUnits = []string{}
		}
		store.Organizations[req.ID] = &req
		store.mu.Unlock()

		writeJSON(w, http.StatusCreated, map[string]any{
			"success": true,
			"data":    req,
		})
	})))

	mux.Handle("PUT /api/v1/admin/organization/{id}", middleware.RequireAuth(domain.RoleStaff, domain.RoleAdmin, domain.RoleSuperadmin)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idVal, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "ID unit tidak valid"})
			return
		}
		var req domain.OrganizationNode
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Format payload tidak valid"})
			return
		}
		req.OrgCode = strings.ToUpper(strings.TrimSpace(req.OrgCode))
		req.ID = uint(idVal)
		if req.SubUnits == nil {
			req.SubUnits = []string{}
		}

		store.mu.Lock()
		if _, exists := store.Organizations[req.ID]; !exists {
			store.mu.Unlock()
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "Simpul struktur organisasi tidak ditemukan"})
			return
		}
		store.Organizations[req.ID] = &req
		store.mu.Unlock()

		writeJSON(w, http.StatusOK, map[string]any{
			"success": true,
			"data":    req,
		})
	})))

	mux.Handle("DELETE /api/v1/admin/organization/{id}", middleware.RequireAuth(domain.RoleStaff, domain.RoleAdmin, domain.RoleSuperadmin)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idVal, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "ID unit tidak valid"})
			return
		}
		store.mu.Lock()
		if _, exists := store.Organizations[uint(idVal)]; !exists {
			store.mu.Unlock()
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "Simpul struktur tidak ditemukan"})
			return
		}
		delete(store.Organizations, uint(idVal))
		store.mu.Unlock()

		writeJSON(w, http.StatusOK, map[string]any{"success": true})
	})))

	// 9. SIAP TERTIB: Kirim Laporan Publik + Notifikasi Email Pelapor
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

		if req.EvidenceImageURL == "" && req.EvidencePhotos != "" {
			req.EvidenceImageURL = req.EvidencePhotos
		}
		if req.EvidencePhotos == "" && req.EvidenceImageURL != "" {
			req.EvidencePhotos = req.EvidenceImageURL
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

		emailNotice := utils.SendComplaintEmail(
			cfg,
			req.ReporterEmail,
			ticket,
			string(req.Status),
			req.RegencyName,
			"",
			"",
		)

		slog.Info("Laporan aduan baru diterima", "ticket", ticket, "anonim", req.IsAnonymous, "jenis", req.ViolationType, "email_mode", emailNotice.Mode)

		writeJSON(w, http.StatusCreated, map[string]any{
			"success":     true,
			"message":     "Laporan aduan Anda berhasil diterima dan sedang diverifikasi petugas.",
			"ticket_code": ticket,
			"email":       emailNotice,
			"data":        req,
		})
	})

	// 10. SIAP TERTIB: Lacak Tiket Aduan Publik (Selalu Masked)
	mux.HandleFunc("GET /api/v1/siaptertib/lacak/{ticket}", func(w http.ResponseWriter, r *http.Request) {
		ticket := strings.ToUpper(strings.TrimSpace(r.PathValue("ticket")))
		store.mu.RLock()
		comp, exists := store.Complaints[ticket]
		store.mu.RUnlock()

		if !exists {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "Nomor tiket aduan tidak ditemukan"})
			return
		}

		// Public privacy rule: ALWAYS mask reporter identity regardless of is_anonymous flag
		maskedName, maskedNIK, maskedPhone, maskedEmail := utils.MaskReporterIdentity(
			comp.ReporterName,
			comp.ReporterNIK,
			comp.ReporterPhone,
			comp.ReporterEmail,
		)
		evidence := comp.EvidenceImageURL
		if evidence == "" {
			evidence = comp.EvidencePhotos
		}

		writeJSON(w, http.StatusOK, map[string]any{
			"success": true,
			"data": map[string]any{
				"ticket_code":        comp.TicketCode,
				"status":             comp.Status,
				"is_anonymous":       comp.IsAnonymous,
				"reporter_name":      maskedName,
				"reporter_nik":       maskedNIK,
				"reporter_phone":     maskedPhone,
				"reporter_email":     maskedEmail,
				"masked_name":        maskedName,
				"masked_nik":         maskedNIK,
				"masked_phone":       maskedPhone,
				"masked_email":       maskedEmail,
				"violation_type":     comp.ViolationType,
				"regency_name":       comp.RegencyName,
				"district_name":      comp.DistrictName,
				"village_name":       comp.VillageName,
				"address_detail":     comp.AddressDetail,
				"latitude":           comp.Latitude,
				"longitude":          comp.Longitude,
				"description":        comp.Description,
				"evidence_image_url": evidence,
				"evidence_photos":    evidence,
				"disposed_to":        comp.DisposedTo,
				"action_notes":       comp.ActionNotes,
				"created_at":         comp.CreatedAt,
			},
		})
	})

	// 11. Admin Satpol PP: List All Complaints
	mux.Handle("GET /api/v1/admin/complaints", middleware.RequireAuth(domain.RoleStaff, domain.RoleAdmin, domain.RoleSuperadmin)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	})))

	// 12. Admin Satpol PP: Update Complaint Status, Disposisi & Email Notification
	mux.Handle("PUT /api/v1/admin/complaints/{ticket}/status", middleware.RequireAuth(domain.RoleStaff, domain.RoleAdmin, domain.RoleSuperadmin)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ticket := strings.ToUpper(strings.TrimSpace(r.PathValue("ticket")))
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
		snapshot := *comp
		store.mu.Unlock()

		emailNotice := utils.SendComplaintEmail(
			cfg,
			snapshot.ReporterEmail,
			snapshot.TicketCode,
			string(snapshot.Status),
			snapshot.RegencyName,
			snapshot.DisposedTo,
			snapshot.ActionNotes,
		)

		writeJSON(w, http.StatusOK, map[string]any{
			"success": true,
			"message": "Status penanganan aduan berhasil diperbarui",
			"email":   emailNotice,
			"data":    snapshot,
		})
	})))

	// 13. Superadmin Diskominfo: User RBAC Management
	mux.Handle("GET /api/v1/superadmin/users", middleware.RequireAuth(domain.RoleSuperadmin)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		store.mu.RLock()
		defer store.mu.RUnlock()
		list := make([]*domain.User, 0, len(store.Users))
		for _, u := range store.Users {
			list = append(list, u)
		}
		writeJSON(w, http.StatusOK, map[string]any{
			"success": true,
			"data":    list,
		})
	})))

	mux.Handle("POST /api/v1/superadmin/users", middleware.RequireAuth(domain.RoleSuperadmin)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Name     string      `json:"name"`
			Username string      `json:"username"`
			Email    string      `json:"email"`
			Password string      `json:"password"`
			Role     domain.Role `json:"role"`
			Bidang   string      `json:"bidang"`
			IsActive *bool       `json:"is_active"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Payload pengguna tidak valid"})
			return
		}
		uname := strings.ToLower(strings.TrimSpace(req.Username))
		if uname == "" || strings.TrimSpace(req.Name) == "" {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Nama dan username wajib diisi"})
			return
		}
		store.mu.Lock()
		if _, dup := store.Users[uname]; dup {
			store.mu.Unlock()
			writeJSON(w, http.StatusConflict, map[string]string{"error": "Username sudah terdaftar"})
			return
		}
		pass := strings.TrimSpace(req.Password)
		if pass == "" {
			pass = "satpolpp2026"
		}
		active := true
		if req.IsActive != nil {
			active = *req.IsActive
		}
		user := &domain.User{
			ID:        store.NextUserID,
			Name:      strings.TrimSpace(req.Name),
			Username:  uname,
			Email:     strings.TrimSpace(req.Email),
			Password:  pass,
			Role:      req.Role,
			Bidang:    strings.TrimSpace(req.Bidang),
			IsActive:  active,
			CreatedAt: time.Now(),
		}
		store.NextUserID++
		store.Users[uname] = user
		store.mu.Unlock()

		writeJSON(w, http.StatusCreated, map[string]any{
			"success": true,
			"data":    user,
		})
	})))

	mux.Handle("PUT /api/v1/superadmin/users/{id}", middleware.RequireAuth(domain.RoleSuperadmin)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idVal, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "ID user tidak valid"})
			return
		}
		var req struct {
			Name     string      `json:"name"`
			Email    string      `json:"email"`
			Password string      `json:"password"`
			Role     domain.Role `json:"role"`
			Bidang   string      `json:"bidang"`
			IsActive *bool       `json:"is_active"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Payload tidak valid"})
			return
		}

		store.mu.Lock()
		var target *domain.User
		for _, u := range store.Users {
			if u.ID == uint(idVal) {
				target = u
				break
			}
		}
		if target == nil {
			store.mu.Unlock()
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "Pengguna tidak ditemukan"})
			return
		}
		if req.Name != "" {
			target.Name = strings.TrimSpace(req.Name)
		}
		if req.Email != "" {
			target.Email = strings.TrimSpace(req.Email)
		}
		if req.Role != "" {
			target.Role = req.Role
		}
		if req.Bidang != "" {
			target.Bidang = strings.TrimSpace(req.Bidang)
		}
		if req.Password != "" {
			target.Password = req.Password
		}
		if req.IsActive != nil {
			target.IsActive = *req.IsActive
		}
		target.UpdatedAt = time.Now()
		snapshot := *target
		store.mu.Unlock()

		writeJSON(w, http.StatusOK, map[string]any{
			"success": true,
			"data":    snapshot,
		})
	})))

	mux.Handle("PATCH /api/v1/superadmin/users/{id}/status", middleware.RequireAuth(domain.RoleSuperadmin)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idVal, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "ID user tidak valid"})
			return
		}
		var req struct {
			IsActive bool `json:"is_active"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Payload tidak valid"})
			return
		}
		store.mu.Lock()
		var target *domain.User
		for _, u := range store.Users {
			if u.ID == uint(idVal) {
				target = u
				break
			}
		}
		if target == nil {
			store.mu.Unlock()
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "Pengguna tidak ditemukan"})
			return
		}
		target.IsActive = req.IsActive
		target.UpdatedAt = time.Now()
		store.mu.Unlock()

		writeJSON(w, http.StatusOK, map[string]any{"success": true})
	})))

	mux.Handle("DELETE /api/v1/superadmin/users/{id}", middleware.RequireAuth(domain.RoleSuperadmin)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idVal, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "ID user tidak valid"})
			return
		}
		store.mu.Lock()
		var keyToDelete string
		for k, u := range store.Users {
			if u.ID == uint(idVal) {
				keyToDelete = k
				break
			}
		}
		if keyToDelete == "" {
			store.mu.Unlock()
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "Pengguna tidak ditemukan"})
			return
		}
		delete(store.Users, keyToDelete)
		store.mu.Unlock()

		writeJSON(w, http.StatusOK, map[string]any{"success": true})
	})))

	// 14. Superadmin Diskominfo: Dynamic CMS Settings Update
	mux.Handle("PUT /api/v1/superadmin/settings", middleware.RequireAuth(domain.RoleSuperadmin)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	})))

	// 15. Strict Image Evidence Upload (MIME Magic-Bytes Validated & 5MB Max)
	imageStorage, err := storage.NewImageStorage("./uploads/evidence")
	if err != nil {
		slog.Error("Gagal inisialisasi storage gambar", "error", err)
	}

	mux.HandleFunc("POST /api/v1/upload/evidence", func(w http.ResponseWriter, r *http.Request) {
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

	rateLimiter := middleware.NewIPRateLimiter(60, time.Second)
	corsMiddleware := middleware.CORS([]string{
		"*",
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
