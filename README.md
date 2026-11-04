# Monorepo Portal Terpadu Satpol PP Prov. Sulsel & SIAP TERTIB

Repositori monorepo resmi untuk rekonstruksi dan modernisasi portal **Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan** yang digabungkan secara penuh dengan sistem layanan pengaduan **SIAP TERTIB** (Sistem Informasi Penegakan Perda dan Trantibum).

---

## 1. Arsitektur Monorepo

```
Web_Remake/
├── backend/                       # Layanan Backend API (Golang v1.24)
│   ├── cmd/api/main.go            # Entrypoint REST API Gateway
│   ├── internal/
│   │   ├── config/                # Konfigurasi runtime & environment
│   │   ├── core/domain/           # Entity domain model & validasi bisnis
│   │   ├── core/ports/            # Interface ports (Hexagonal Architecture)
│   │   └── utils/                 # Generator tiket unik (TRB-YYYYMM-XXXX)
│   ├── go.mod                     # Dependensi Go
│   └── run_backend.bat            # Launcher mandiri backend
│
├── frontend/                      # Layanan Frontend Web (SvelteKit 2 + TailwindCSS)
│   ├── src/
│   │   ├── lib/
│   │   │   ├── api/client.ts      # HTTP client dengan graceful offline fallback
│   │   │   ├── components/        # Navbar, Footer, ComplaintForm, TicketTracker
│   │   │   ├── stores/appState.ts # Reactive persistent store (Settings, Articles, Complaints)
│   │   │   └── types/             # TypeScript definitions
│   │   └── routes/                # SvelteKit SSR Routes
│   │       ├── +page.svelte       # Beranda utama & slider hero
│   │       ├── berita/            # Indeks & detail warta berita
│   │       ├── dokumen/           # Transparansi dokumen publik (Renja, DPA, LAKIP)
│   │       ├── profil/[slug]/     # Sejarah, Visi-Misi, Struktur, Tupoksi, Pejabat
│   │       ├── bidang/[slug]/     # 5 Bidang teknis Satpol PP
│   │       ├── galeri/            # Galeri multimedia kegiatan
│   │       ├── kontak/            # Alamat & formulir korespondensi
│   │       ├── lapor/             # Form aduan SIAP TERTIB (GPS & NIK)
│   │       ├── lacak/             # Pelacakan tiket aduan real-time
│   │       ├── admin/dashboard/   # Konsol Staff: Meja aduan SOP, CMS Berita, CMS Dokumen, CMS Galeri
│   │       └── superadmin/        # Konsol Diskominfo: Anti-hardcode CMS, RBAC, Wilayah, Audit
│   ├── static/uploads/            # Aset otentik (logo, foto pimpinan, galeri, berita)
│   └── package.json               # Dependensi SvelteKit & Vite
│
├── deployments/                   # Konfigurasi Deployment Server Diskominfo
│   ├── docker/
│   │   ├── Dockerfile.backend     # Multi-stage build Go Alpine static binary (~15MB)
│   │   ├── Dockerfile.frontend    # Multi-stage build Node.js SvelteKit
│   │   └── nginx.conf             # Nginx reverse proxy & SSL termination
│   └── docker-compose.yml         # Orkestrasi container satu-perintah
│
├── scripts/                       # Skrip Otomasi Pengembang
│   ├── setup_git_alias.bat        # Konfigurasi identitas commit alias khusus (lokal repo)
│   └── start_dev.bat              # Menjalankan frontend & backend secara bersamaan
│
├── Makefile                       # Antarmuka CLI terstandar (make dev, make build, make docker-up)
├── .gitignore                     # Filter berkas build, node_modules, dan binary
└── start_frontend_preview.bat     # Launcher 1-klik pengujian tampilan lokal
```

---

## 2. Menjalankan di Lingkungan Lokal (Development)

- **Frontend SvelteKit:** Port `3000` (`http://localhost:3000`)
- **Backend Golang API:** Port `8080` (`http://localhost:8080`)

Jalankan langsung melalui terminal:
```bash
scripts\start_dev.bat
```
Atau jalankan masing-masing:
```bash
# Terminal 1: Backend
cd backend && go run cmd/api/main.go

# Terminal 2: Frontend
cd frontend && npm run dev -- --open
```

---

## 3. Akun Default Pengujian Dashboard (Demo)

| Role | Username | Password | Deskripsi |
|---|---|---|---|
| **Superadmin Diskominfo** | `superadmin` | `satpolpp2026` | Tata kelola CMS anti-hardcoded, RBAC, master wilayah, audit trail |
| **Pimpinan Kasatpol PP** | `kasatpol` | `satpolpp2026` | Akses monitoring laporan dan persetujuan penindakan |
| **Staff Trantibum** | `stafftrantibum` | `satpolpp2026` | Tindak lanjut laporan SOP ketenteraman dan ketertiban umum |
| **Staff Penegakan Perda** | `staffperda` | `satpolpp2026` | Penanganan pelanggaran peraturan daerah dan pembuatan BAP |
