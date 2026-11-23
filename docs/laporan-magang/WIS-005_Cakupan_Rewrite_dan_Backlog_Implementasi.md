# Dokumen Backlog Implementasi & Dashboard Penugasan
## Sprint Revisi Desain & Antarmuka Terpadu (13 – 23 November 2026)

* **Instansi:** Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan  
* **Unit Kerja Magang:** Dinas Komunikasi, Informatika, Statistik dan Persandian (Diskominfo-SP) Prov. Sulsel  
* **Penyusun:** Mutiah Arinil Fayza Nusar  
* **Kode Dokumen:** WIS-005  
* **Periode Pengerjaan:** 13 November 2026 – 23 November 2026  
* **Repositori Proyek:** https://github.com/mutiaharn/SatpolPP (Branch: `main`)  
* **Status Keseluruhan:** 100% Selesai (Release Candidate - Build Verified)

---

## 1. Latar Belakang & Fokus Sprint (13–23 November 2026)

Sprint revisi intensif periode **13 hingga 23 November 2026** berfokus pada modernisasi estetika visual, optimasi ergonomi pengguna, pembersihan tata letak (*de-cluttering*), dan peningkatan aksesibilitas seluruh antarmuka portal publik Satpol PP Prov. Sulsel serta modul SIAP TERTIB.

Pekerjaan pada rentang waktu ini menindaklanjuti masukan evaluasi teknis, meliputi:
1. Transformasi bahasa desain menjadi material kaca translusen (*Liquid Glass*) dengan tata cahaya *living ambient lighting*.
2. Implementasi modul navigasi kilat *Global Command Palette* (`Ctrl+K` / `⌘K`).
3. Modernisasi antarmuka autentikasi petugas (*split-screen admin login*).
4. Perampingan tajuk halaman dalam (*Slim Hero Header* ~60–80px) dan integrasi *breadcrumb trail*.
5. Restrukturisasi panggung sambutan Kasatpol PP menjadi kutipan inspiratif terfokus (*inspiring vocal point quote*).
6. Penyelarasan ergonomi antarmuka dengan standar *Apple Human Interface Guidelines (HIG)* (area sentuh ≥40–44 pt, tipografi native SF Pro, dan transisi taktil pegas).
7. Pengujian kompilasi produksi (*production build*) serta penyusunan dokumen laporan naratif komprehensif berformat DOCX.

---

## 2. Rincian Task Pokok & Sub-Task (Jobdesk 13–23 November 2026)

### TASK 1: Perancangan & Implementasi Liquid Glass Design System
* **Kode Task:** WIS-015
* **Rentang Tanggal:** 13 – 18 November 2026
* **Status:** Selesai (100%)
* **Deskripsi:** Merancang dan menerapkan arsitektur token desain visual modern berbasis material kaca optik (*Liquid Glass*) dengan saturasi warna tinggi dan tata cahaya dinamis.
* **Sub-Tasks:**
  * [x] **Sub-task 1.1:** Menyusun token CSS material kaca translusen (`.ds-glass`, `.ds-glass-strong`, `.ds-glass-dark`, `.ds-glass-gold`, `.ds-glass-azure`) dengan efek pembiasan `backdrop-filter: blur(16px)` dan saturasi 180%.
  * [x] **Sub-task 1.2:** Menambahkan aksen pantulan kristal atas (*specular inner rim highlight* `inset 0 1px 1px 0 rgba(255, 255, 255, 0.25)`) pada seluruh kontainer kartu.
  * [x] **Sub-task 1.3:** Mengintegrasikan lapisan gradasi pencahayaan hidup (*living ambient glow*) pada tata letak global untuk menciptakan kedalaman spasial saat halaman digulir.
  * [x] **Sub-task 1.4:** Menyelaraskan palet identitas resmi: *Deep Obsidian Navy* (`#070e1c`), *Sovereign Azure* (`#246bfd`), dan *Praja Amber* (`#F08300`).
* **Deliverables:** `frontend/src/app.css`, `frontend/src/routes/+layout.svelte` (Commit: `56110f3`).

---

### TASK 2: Implementasi Global Command Palette & Navigasi Cepat (Ctrl+K)
* **Kode Task:** WIS-016
* **Rentang Tanggal:** 18 November 2026
* **Status:** Selesai (100%)
* **Deskripsi:** Membangun antarmuka pencarian universal instan ala macOS Spotlight / Command Palette untuk mempermudah akses navigasi publik dan verifikasi data.
* **Sub-Tasks:**
  * [x] **Sub-task 2.1:** Mengembangkan komponen modal `CommandPalette.svelte` dengan *backdrop blur* gelap sinematik.
  * [x] **Sub-task 2.2:** Menyusun indeks pencarian terintegrasi untuk menu rute navigasi, katalog warta berita, direktori dokumen perda, dan pelacakan tiket aduan.
  * [x] **Sub-task 2.3:** Menghubungkan *listener* pintasan papan ketik universal `Ctrl+K` dan `⌘K` serta tombol pemicu pencarian pada bilah navigasi utama (`Navbar.svelte`).
  * [x] **Sub-task 2.4:** Mengimplementasikan penyaringan teks reaktif seketika (*zero-latency instant filter*) dengan indikator *badge* kategori.
* **Deliverables:** `frontend/src/lib/components/ui/CommandPalette.svelte`, `frontend/src/lib/components/layout/Navbar.svelte` (Commit: `56110f3`).

---

### TASK 3: Modernisasi Antarmuka Masuk Petugas & CMS Split-Screen
* **Kode Task:** WIS-017
* **Rentang Tanggal:** 18 November 2026
* **Status:** Selesai (100%)
* **Deskripsi:** Memperbarui halaman masuk administrator internal dengan tata letak *split-screen* modern dan pengamanan visual kredensial.
* **Sub-Tasks:**
  * [x] **Sub-task 3.1:** Merombak tata letak login satu kolom lama menjadi komposisi *split-screen* futuristik.
  * [x] **Sub-task 3.2:** Merancang panel samping kiri dengan latar belakang grafis identitas korps praja, lencana resmi beranimasi halus, dan ringkasan fitur proteksi sistem.
  * [x] **Sub-task 3.3:** Menerapkan panel kanan dengan kartu kaca melayang (*floating glass card*), kolom input berbingkai interaktif, dan tombol aksi tegas.
  * [x] **Sub-task 3.4:** Mengintegrasikan pemisahan hak akses berjenjang antara peran Admin Petugas dan Superadmin Komando.
* **Deliverables:** `frontend/src/routes/admin/login/+page.svelte` (Commit: `7ab582f`).

---

### TASK 4: Perampingan Tata Letak & Standardisasi Slim Hero Header Subpage
* **Kode Task:** WIS-018
* **Rentang Tanggal:** 18 – 20 November 2026
* **Status:** Selesai (100%)
* **Deskripsi:** Menghilangkan ruang kosong berlebih (*vertical clutter*) pada seluruh halaman dalam dengan menerapkan tajuk ramping (*Slim Hero Header*) dan navigasi remah roti.
* **Sub-Tasks:**
  * [x] **Sub-task 4.1:** Memangkas tinggi banner halaman dalam dari ukuran besar menjadi tajuk ramping ringkas berketinggian ~60–80px.
  * [x] **Sub-task 4.2:** Mengintegrasikan hierarki navigasi remah roti (*breadcrumb trail*) interaktif pada tajuk subpage untuk orientasi lokasi pengguna yang presisi.
  * [x] **Sub-task 4.3:** Menerapkan standardisasi tajuk ramping secara seragam di 8 halaman subpage: Profil, Bidang, Berita, Dokumen Publik, Galeri Kegiatan, Kontak, Lapor Aduan, dan Lacak Tiket.
  * [x] **Sub-task 4.4:** Mengharmonisasikan konsistensi padding, margin kontainer, dan hierarki judul pada 27 berkas antarmuka frontend.
* **Deliverables:** Seluruh direktori rute subpage `frontend/src/routes/(app)/*` (Commit: `eef4e9d`).

---

### TASK 5: Restrukturisasi Panggung Sambutan Kasatpol PP sebagai Vocal Point
* **Kode Task:** WIS-019
* **Rentang Tanggal:** 18 – 20 November 2026
* **Status:** Selesai (100%)
* **Deskripsi:** Mengubah penyajian sambutan pimpinan dari boks artikel kaku menjadi panggung kutipan inspiratif yang menjadi titik fokus (*vocal point*) beranda.
* **Sub-Tasks:**
  * [x] **Sub-task 5.1:** Meniadakan format kartu sambutan berbentuk kotak teks biasa yang memadati beranda.
  * [x] **Sub-task 5.2:** Menjadikan kutipan visi pimpinan sebagai *vocal point* utama: "Menegakkan Perda Secara Humanis, Persuasif, dan Berkeadilan".
  * [x] **Sub-task 5.3:** Merancang ornamen visual panggung *frosted arch halo backdrop* bertingkat di belakang foto resmi Kasatpol PP Sulsel.
  * [x] **Sub-task 5.4:** Menambahkan lencana identitas pejabat melayang (*floating officer identity badge*) dengan tipografi aksen emas.
  * [x] **Sub-task 5.5:** Menghapus kartu promosi statis ganda di beranda agar alur antarmuka mengalir dinamis langsung ke warta kegiatan dan layanan pengaduan.
* **Deliverables:** `frontend/src/routes/+page.svelte` (Commit: `eef4e9d`, `7ba97ad`, `a3659a9`).

---

### TASK 6: Penyelarasan Standar Ergonomi Antarmuka Apple HIG
* **Kode Task:** WIS-020
* **Rentang Tanggal:** 21 – 23 November 2026
* **Status:** Selesai (100%)
* **Deskripsi:** Mengoptimalkan seluruh komponen interaktif agar memenuhi standar keramahan sentuhan, tipografi retina, dan mikro-interaksi taktil sesuai Apple Human Interface Guidelines.
* **Sub-Tasks:**
  * [x] **Sub-task 6.1:** Mengalibrasi area target sentuh minimal 40–44 pt (40–44 px) pada tautan navigasi desktop, menu hamburger mobile, tombol pencarian, tombol filter kategori, kontrol paginasi, dan kapsul indikator carousel beranda.
  * [x] **Sub-task 6.2:** Menerapkan tatanan tipografi native Apple SF Pro (`-apple-system, BlinkMacSystemFont, "SF Pro Display", "SF Pro Text"`) dengan *subpixel anti-aliasing* dan *optimizeLegibility*.
  * [x] **Sub-task 6.3:** Menambahkan mikro-interaksi fisik taktil saat elemen ditekan (`active:scale-[0.975]`) dengan kurva transisi pegas alami (`cubic-bezier(0.16, 1, 0.3, 1)`).
  * [x] **Sub-task 6.4:** Mengintegrasikan cincin fokus visual ganda (*specular focus halo ring*) untuk kenyamanan navigasi aksesibilitas papan ketik sesuai standar WCAG AA.
* **Deliverables:** Seluruh komponen tombol, slider, dan navigasi di `frontend/src/` (Commit: `7409ab0`).

---

### TASK 7: Penyusunan Laporan Progres Naratif DOCX & Verifikasi Kompilasi
* **Kode Task:** WIS-021
* **Rentang Tanggal:** 23 November 2026
* **Status:** Selesai (100%)
* **Deskripsi:** Menyusun dokumen laporan kemajuan teknis format DOCX lengkap dengan narasi mendalam dan 10 tangkapan layar antarmuka, serta menjalankan uji build rilis.
* **Sub-Tasks:**
  * [x] **Sub-task 7.1:** Mengambil dan mengurasi 10 tangkapan layar antarmuka sistem beresolusi tinggi yang merepresentasikan seluruh hasil pembaruan visual.
  * [x] **Sub-task 7.2:** Mengembangkan skrip generator dokumen otomatis berbasis `python-docx` (`generate_docx.py`) dengan tipografi elegan dan tabel terstruktur.
  * [x] **Sub-task 7.3:** Menyusun dokumen laporan kemajuan teknis setebal 6 bab dengan penjelasan naratif komprehensif (`Laporan_Progress_UIUX_SatpolPP_Sulsel.docx`, 2.46 MB).
  * [x] **Sub-task 7.4:** Menjalankan verifikasi kompilasi produksi SvelteKit (`npm run build`) dengan status *zero-error*.
  * [x] **Sub-task 7.5:** Menyinkronkan seluruh perubahan kode ke repositori GitHub cabang `main`.
* **Deliverables:** `Laporan_Progress_UIUX_SatpolPP_Sulsel.docx`, repositori GitHub `mutiaharn/SatpolPP` (Commit: `3f8593a`).

---

## 3. Matriks Dashboard Penugasan (Siap Salin ke Magangki')

Tabel di bawah ini dirancang khusus untuk mempermudah pengisian formulir penugasan (*Task & Sub-task Modal*) pada portal **Magangki'**:

| Kode Task | Judul Penugasan (*Task Name*) | Rentang Tanggal | Status | Deliverable / Output Utama |
|:---:|:---|:---:|:---:|:---|
| **WIS-015** | Perancangan & Implementasi Liquid Glass Design System | 13–18 Nov 2026 | **Selesai (100%)** | Token CSS material kaca optik, refraksi *living ambient*, & palet Deep Obsidian Navy / Sovereign Azure |
| **WIS-016** | Implementasi Global Command Palette & Navigasi Cepat (Ctrl+K) | 18 Nov 2026 | **Selesai (100%)** | Modal pencarian instan universal, indeks rute menu/berita/dokumen, & pintasan keyboard `Ctrl+K` |
| **WIS-017** | Modernisasi Halaman Masuk Administrator & CMS Split-Screen | 18 Nov 2026 | **Selesai (100%)** | Tata letak login *split-screen*, kartu kaca melayang, ilustrasi identitas, & segregasi hak akses |
| **WIS-018** | Perampingan Tata Letak & Standardisasi Slim Hero Header Subpage | 18–20 Nov 2026 | **Selesai (100%)** | Tajuk ramping 60–80px di 8 subpage utama, navigasi *breadcrumb trail*, & eliminasi *vertical clutter* |
| **WIS-019** | Restrukturisasi Panggung Sambutan Kasatpol PP sebagai Vocal Point | 18–20 Nov 2026 | **Selesai (100%)** | Panggung kutipan inspiratif pimpinan, ornamen *frosted arch halo backdrop*, & lencana dinas melayang |
| **WIS-020** | Penyelarasan Standar Ergonomi Antarmuka Apple HIG | 21–23 Nov 2026 | **Selesai (100%)** | Target sentuh ≥40–44 pt, tipografi native SF Pro, respons taktil pegas (`active:scale-[0.975]`), & cincin fokus WCAG AA |
| **WIS-021** | Penyusunan Laporan Progres Naratif DOCX & Verifikasi Kompilasi | 23 Nov 2026 | **Selesai (100%)** | Dokumen naratif DOCX (6 bab, 10 screenshot HD), kompilasi `npm run build` sukses, & sinkronisasi GitHub `main` |

---

## 4. Definisi Selesai (*Definition of Done* / DoD)

Seluruh tugas dalam sprint revisi 13–23 November dinyatakan selesai berdasarkan kriteria keberhasilan berikut:
1. **Verifikasi Kompilasi:** Lolos kompilasi penuh TypeScript dan SvelteKit (`npm run build`) tanpa peringatan galat (*zero error*).
2. **Kepatuhan Ergonomi & Aksesibilitas:** Seluruh komponen interaktif mematuhi batas sentuh ergonomis minimal 40–44 pt dan lulus rasio kontras teks WCAG AA (≥ 4.5:1).
3. **Integritas Versi Repositori:** Seluruh kode sumber telah di-*commit* secara berkala dengan pesan jelas dan diunggah (*push*) ke repositori resmi cabang `main`.
4. **Dokumentasi & Bukti Visual:** Laporan progres naratif format DOCX telah disusun rapi menyertakan bukti tangkapan layar antarmuka sistem.
