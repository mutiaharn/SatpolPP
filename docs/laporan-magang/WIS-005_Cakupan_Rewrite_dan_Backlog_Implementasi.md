# Dokumen Cakupan Rewrite & Backlog Implementasi Lengkap
**Instansi:** Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan  
**Kode Task Magang:** WIS-005  
**Penyusun:** Mutiah Arinil Fayza Nusar  
**Repositori Proyek:** https://github.com/mutiaharn/SatpolPP  
**Pembaruan Terakhir:** 23 November 2026  
**Status Keseluruhan:** Selesai (Release Candidate - Build Verified)

---

## 1. Cakupan Minimum Viable Product (MVP) & Ruang Lingkup Sistem
Proyek penulisan ulang (*rewrite*) portal resmi Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan dan integrasi aplikasi pengaduan masyarakat **SIAP TERTIB** (Sistem Informasi Penegakan Perda & Trantibum) dirancang untuk mentransformasi sistem monolitik lama menjadi aplikasi web modern berbasis **SvelteKit, TypeScript, dan Tailwind CSS**.

Sistem baru ini menggabungkan layanan aspirasi publik, pelacakan tiket aduan secara *real-time*, keterbukaan dokumen regulasi daerah, publikasi warta lapangan, dan sistem manajemen konten (CMS) internal yang aman, cepat, dan responsif di berbagai perangkat (*mobile*, tablet, dan *desktop*).

---

## 2. Arsitektur Komponen UI Reusable
Seluruh elemen antarmuka diabstraksikan menjadi komponen pakai-ulang (*reusable components*) untuk memastikan kemudahan pemeliharaan kode dan konsistensi visual:
1. **`DataTable.svelte`:** Komponen tabel universal dengan fitur penyaringan teks reaktif (*on-chain search*), pemilih jumlah baris dinamis (5, 10, 25, 50), paginasi cerdas dengan *ellipsis*, dan slot baris kustom.
2. **`FileUploader.svelte`:** Komponen input pengunggahan berkas bukti foto kejadian dan thumbnail berita bergaya modern dengan kompresi otomatis dan pratinjau instan.
3. **`Modal.svelte`:** Dialog modal interaktif berbasis standar aksesibilitas WAI-ARIA untuk inspeksi dokumen regulasi PDF dan rincian berkas aduan tanpa berpindah halaman.
4. **`Lightbox.svelte`:** Penampil foto dokumentasi galeri beresolusi tinggi dengan kendali navigasi papan ketik (panah kiri/kanan dan tombol Escape) serta latar belakang gelap sinematik.
5. **`CommandPalette.svelte`:** Modul pencarian global cepat (*Spotlight/Command Palette*) yang dipicu melalui pintasan papan ketik universal `Ctrl+K` atau `⌘K` untuk menavigasi menu, mencari berita, mengunduh perda, atau melacak tiket secara instan.
6. **`StatusBadge.svelte`:** Indikator status aduan dinamis (Diterima, Terverifikasi, Disposisi, Tindakan Lapangan, Selesai/BAP) dengan skema warna semantik dan kontras tinggi.
7. **`Navbar.svelte` & `Footer.svelte`:** Komponen navigasi kaca melayang (*floating glass bar*) dan kaki halaman terpadu yang otomatis terisolasi saat admin berada pada rute internal `/admin/*`.

---

## 3. Product & Sprint Backlog Implementasi Berdasarkan Jobdesk Riil

Berikut adalah pembagian rincian pekerjaan (*jobdesk*), tugas pokok (*tasks*), dan sub-tugas (*sub-tasks*) yang telah dikerjakan secara tuntas:

### SPRINT 1: Analisis Kebutuhan, Inventarisasi & Penyiapan Fondasi Sistem
* **Task 1: Audit dan Inventarisasi Situs Eksisting (WIS-001)**
  * **Sub-task 1.1:** Melakukan *crawling* dan pemetaan seluruh URL menu serta dokumen pada website lama Satpol PP Sulsel.
  * **Sub-task 1.2:** Mengidentifikasi *broken links*, aset gambar beresolusi rendah, dan galat 404 pada direktori regulasi daerah.
  * **Sub-task 1.3:** Menyusun dokumen rekomendasi modernisasi tata kelola informasi portal publik.
* **Task 2: Perancangan Alur Pengguna & SOP SIAP TERTIB (WIS-002)**
  * **Sub-task 2.1:** Merumuskan alur penanganan laporan 4 tahapan: Verifikasi &rarr; Disposisi &rarr; Penindakan &rarr; Berita Acara (BAP).
  * **Sub-task 2.2:** Menetapkan format standar nomor resi tiket pengaduan publik (`TRB-YYYYMM-XXXXX`).
  * **Sub-task 2.3:** Merancang kebijakan privasi warga (*data masking*) pada NIK KTP dan nomor kontak pelapor.
* **Task 3: Penyiapan Lingkungan Pengembangan & Kerangka Kerja (WIS-006 & WIS-007)**
  * **Sub-task 3.1:** Inisialisasi proyek SvelteKit dengan adapter statis/node, TypeScript, dan Tailwind CSS.
  * **Sub-task 3.2:** Optimalisasi aset logo resmi, banner dinas, dan foto pimpinan ke format WebP terkompresi.
  * **Sub-task 3.3:** Konfigurasi struktur routing modular (`/`, `/profil/*`, `/bidang/*`, `/berita/*`, `/dokumen`, `/lapor`, `/lacak`, `/kontak`, `/admin/*`).

---

### SPRINT 2: Pengembangan Fitur Inti Pelayanan Publik (SIAP TERTIB & Portal Berita)
* **Task 4: Implementasi Formulir Pengaduan Warga SIAP TERTIB (WIS-010)**
  * **Sub-task 4.1:** Membangun formulir pengaduan interaktif dengan validasi wajib 16 digit NIK, nomor WhatsApp aktif, dan alamat kejadian.
  * **Sub-task 4.2:** Mengintegrasikan opsi pelaporan anonim terlindungi dan penyamaran otomatis (*masking*) data pribadi.
  * **Sub-task 4.3:** Mengintegrasikan komponen pengunggahan bukti foto lapangan dengan pembatasan ukuran berkas.
  * **Sub-task 4.4:** Menetapkan koordinat default Sulawesi Selatan (Makassar) tanpa memicu *permission popup* GPS yang mengganggu kenyamanan warga.
  * **Sub-task 4.5:** Menerbitkan resi tiket digital instan dan simulasi notifikasi surel konfirmasi pelapor.
* **Task 5: Implementasi Modul Pelacakan Tiket Real-Time (WIS-012)**
  * **Sub-task 5.1:** Membangun konsol pelacakan nomor resi dengan pencarian cepat berbasis URL parameter (`/lacak?ticket=...`).
  * **Sub-task 5.2:** Menampilkan visualisasi garis waktu (*step timeline*) progres penanganan aduan dari verifikasi awal hingga penyusunan BAP.
  * **Sub-task 5.3:** Menampilkan foto bukti penindakan lapangan oleh personel pleton operasional.
* **Task 6: Implementasi Direktori Dokumen Publik & Warta Berita (WIS-011)**
  * **Sub-task 6.1:** Mengompilasi dan mengkategorisasi berkas transparansi daerah (Renja, Renstra, DPA, LAKIP, Keuangan, Perda, SOP).
  * **Sub-task 6.2:** Menghubungkan data dokumen dengan komponen `DataTable` interaktif berkemampuan *live search*.
  * **Sub-task 6.3:** Membangun halaman katalog berita dan halaman detail dinamis (`/berita/[slug]`) dengan estimasi waktu baca dan artikel rekomendasi.

---

### SPRINT 3: Sistem Administrasi, Investigasi Forensik & Manajemen Konten (CMS)
* **Task 7: Otentikasi Petugas & Antarmuka Dashboard Admin**
  * **Sub-task 7.1:** Membangun halaman login petugas modern dengan tata letak *split-screen* dan validasi kredensial.
  * **Sub-task 7.2:** Menerapkan pembagian peran otorisasi berjenjang antara Admin Petugas Pengolah Aduan dan Superadmin Komando.
  * **Sub-task 7.3:** Mengembangkan tata letak *DashboardShell* dengan bilah sisi (*sidebar*) responsif dan navigasi ringkas.
* **Task 8: Panel Penanganan Aduan & Investigasi Forensik**
  * **Sub-task 8.1:** Membangun tabel manajemen tiket masuk dengan pemfilteran status dan disposisi bidang operasional.
  * **Sub-task 8.2:** Mengimplementasikan fitur *Forensic Unmasking* yang dilindungi kata sandi master untuk membuka identitas asli pelapor khusus keperluan penyelidikan yustisi/PPNS, lengkap dengan pencatatan audit log.
* **Task 9: Manajemen Konten Kelembagaan & Pengolah Kata Kaya Format**
  * **Sub-task 9.1:** Membangun editor struktur organisasi dinamis (*OrganizationManager*) untuk pemutakhiran bagan hierarki pejabat eselon secara visual.
  * **Sub-task 9.2:** Membangun modul publikasi warta berita resmi dengan pengolah kata kaya format (*RichTextEditor*).

---

### SPRINT 4: Modernisasi Desain Sistem & Arsitektur Visual Kaca (Liquid Glass)
* **Task 10: Rekayasa Material Kaca Optik & Palet Warna Resmi**
  * **Sub-task 10.1:** Membangun sistem token CSS material kaca translusen (`.ds-glass`, `.ds-glass-strong`, `.ds-glass-dark`) dengan efek saturasi 180% dan pembiasan *backdrop-filter*.
  * **Sub-task 10.2:** Menyusun lapisan gradasi *living ambient* pada tata letak global untuk menciptakan refraksi cahaya yang hidup saat digulir.
  * **Sub-task 10.3:** Menyelaraskan palet identitas daerah resmi: *Deep Obsidian Navy* (`#070e1c`), *Sovereign Azure* (`#246bfd`), dan *Praja Amber* (`#F08300`).
  * **Sub-task 10.4:** Menambahkan detail pantulan cahaya kristal atas (*specular inner rim highlight*) pada seluruh kontainer kartu.
* **Task 11: Integrasi Pencarian Instan Global (Command Palette)**
  * **Sub-task 11.1:** Membangun komponen modal `CommandPalette.svelte` yang mengindeks navigasi rute, artikel berita, dokumen perda, dan pelacakan tiket aduan.
  * **Sub-task 11.2:** Menghubungkan pintasan papan ketik universal `Ctrl+K` / `⌘K` dan tombol pemicu pencarian pada bilah navigasi utama.
* **Task 12: Penataan Ulang Tata Letak & Pembersihan Visual (De-cluttering)**
  * **Sub-task 12.1:** Menyeragamkan seluruh banner halaman dalam menjadi *Slim Hero Header* ramping (~60–80px) dengan remah roti (*breadcrumb trail*) terintegrasi.
  * **Sub-task 12.2:** Merestrukturisasi panggung sambutan Kasatpol PP: meniadakan boks statis kaku, menjadikan kutipan pimpinan sebagai *vocal point* utama (&ldquo;Menegakkan Perda Secara Humanis, Persuasif, dan Berkeadilan&rdquo;) dengan foto resmi ber-halo bertingkat.
  * **Sub-task 12.3:** Menghilangkan kartu statis ganda di beranda agar alur interaksi mengalir alami dari slider utama ke warta kegiatan dan dokumen publik.

---

### SPRINT 5: Penyelarasan Standar Ergonomi Antarmuka Kelas Dunia (Apple HIG)
* **Task 13: Standarisasi Ergonomi Target Sentuh (Minimum 44 pt / 44 px)**
  * **Sub-task 13.1:** Menyesuaikan tinggi interaksi seluruh tautan menu desktop dan item dropdown menjadi minimal 40–44px.
  * **Sub-task 13.2:** Memadatkan pemicu menu mobile (hamburger) dan tombol pencarian mobile ke batas sentuh ergonomis `44×44px`.
  * **Sub-task 13.3:** Meredesain indikator navigasi carousel slider beranda dari titik kecil menjadi kapsul tombol ergonomis dengan area sentuh `32×32px` / `40px` dengan tetap mempertahankan siluet visual yang elegan.
  * **Sub-task 13.4:** Menstandarisasi tombol filter kategori (*filter pills*) pada halaman Berita, Dokumen, dan Galeri menjadi `min-h-[40px] px-4 py-2.5` dengan kontras aktif yang tegas.
  * **Sub-task 13.5:** Memperbesar tombol kontrol paginasi tabel dan berita menjadi `36×36px` / `38px` dengan sudut membulat (*squircle*) yang ramah sentuhan.
* **Task 14: Optimalisasi Tipografi Retina & Aksesibilitas Sensorik**
  * **Sub-task 14.1:** Mengonfigurasi font stack native Apple SF Pro (`-apple-system, BlinkMacSystemFont, "SF Pro Display", "SF Pro Text"`) dan Plus Jakarta Sans.
  * **Sub-task 14.2:** Mengaktifkan penghalusan teks tingkat lanjut (`-webkit-font-smoothing: antialiased; text-rendering: optimizeLegibility;`).
  * **Sub-task 14.3:** Memastikan rasio kontras teks memenuhi standar aksesibilitas WCAG AA (rasio kontras &ge; 4.5:1).
  * **Sub-task 14.4:** Menerapkan efek taktil saat elemen ditekan (`active:scale-[0.975]`) dengan kurva transisi pegas alami (`cubic-bezier(0.16, 1, 0.3, 1)`).
  * **Sub-task 14.5:** Menerapkan cincin fokus visual ganda (*focus-visible halo ring*) untuk navigasi aksesibilitas papan ketik.

---

### SPRINT 6: Pengujian Kualitas, Verifikasi Produksi & Pelaporan Kemajuan
* **Task 15: Verifikasi Kompilasi Produksi & Pengujian Responsivitas**
  * **Sub-task 15.1:** Menjalankan uji *production build* (`npm run build`) dengan hasil kompilasi bersih (zero lint/type error).
  * **Sub-task 15.2:** Menguji responsivitas antarmuka pada berbagai resolusi layar (ponsel 360px–414px, tablet 768px, desktop 1024px–1920px).
* **Task 16: Penyusunan Laporan Progres & Sinkronisasi Repositori**
  * **Sub-task 16.1:** Menyusun dokumen laporan kemajuan teknis format DOCX lengkap dengan narasi mendalam dan 10 tangkapan layar antarmuka sistem.
  * **Sub-task 16.2:** Menyelaraskan seluruh perubahan kode sumber dan mengunggah (*push*) ke repositori GitHub cabang utama (`origin/main`).

---

## 4. Matriks Ringkasan Status Penyelesaian Backlog

| Kode Task | Nama Pekerjaan / Jobdesk | Target Sprint | Persentase Selesai | Status |
|:---|:---|:---:|:---:|:---:|
| **WIS-001** | Audit & Inventarisasi Website Lama | Sprint 1 | 100% | Selesai |
| **WIS-002** | User Flow & SOP SIAP TERTIB | Sprint 1 | 100% | Selesai |
| **WIS-003** | Audit Konten & Integrasi OPD | Sprint 1 | 100% | Selesai |
| **WIS-004** | Evaluasi Aksesibilitas WCAG & Mobile | Sprint 1 | 100% | Selesai |
| **WIS-005** | Cakupan Rewrite & Backlog Implementasi | Sprint 1 | 100% | Selesai (Diperbarui) |
| **WIS-006** | Optimasi Aset Grafis & Format WebP | Sprint 1 | 100% | Selesai |
| **WIS-007** | Setup Kerangka SvelteKit & Routing | Sprint 1 | 100% | Selesai |
| **WIS-008** | Desain Mockup Form Lapor & Resi | Sprint 2 | 100% | Selesai |
| **WIS-009** | Implementasi Beranda & Profil Instansi | Sprint 2 | 100% | Selesai |
| **WIS-010** | Implementasi Form Pengaduan SIAP TERTIB | Sprint 2 | 100% | Selesai |
| **WIS-011** | Kompilasi Dokumen Perda & Warta Berita | Sprint 2 | 100% | Selesai |
| **WIS-012** | Implementasi Tracking Resi Tiket Aduan | Sprint 2 | 100% | Selesai |
| **WIS-013** | Pengujian Responsivitas & Bug Layout | Sprint 2 | 100% | Selesai |
| **WIS-014** | Bahan Paparan Teknis Progres Pertama | Sprint 2 | 100% | Selesai |
| **WIS-015** | Otentikasi Petugas & Dashboard Shell CMS | Sprint 3 | 100% | Selesai |
| **WIS-016** | Panel Manajemen Aduan & Investigasi Forensik | Sprint 3 | 100% | Selesai |
| **WIS-017** | Manajemen SOTK & Editor Bagan Organisasi | Sprint 3 | 100% | Selesai |
| **WIS-018** | Editor Warta Berita (Rich Text Editor) | Sprint 3 | 100% | Selesai |
| **WIS-019** | Sistem Desain Material Kaca (Liquid Glass) | Sprint 4 | 100% | Selesai |
| **WIS-020** | Global Command Palette (Pencarian Ctrl+K) | Sprint 4 | 100% | Selesai |
| **WIS-021** | Perampingan Subpage & Panggung Sambutan | Sprint 4 | 100% | Selesai |
| **WIS-022** | Kalibrasi Target Sentuh Ergonomis 44 pt | Sprint 5 | 100% | Selesai |
| **WIS-023** | Standarisasi Tipografi Native SF Pro | Sprint 5 | 100% | Selesai |
| **WIS-024** | Mikro-Interaksi Taktil & Focus Ring Aksesibilitas | Sprint 5 | 100% | Selesai |
| **WIS-025** | Pengujian Build Produksi SvelteKit Zero-Error | Sprint 6 | 100% | Selesai |
| **WIS-026** | Pembuatan Dokumen Laporan Naratif DOCX | Sprint 6 | 100% | Selesai |
| **WIS-027** | Sinkronisasi Repositori Git & Rilis Final | Sprint 6 | 100% | Selesai |

---

## 5. Definisi Selesai (*Definition of Done* / DoD)
Setiap tugas (*task*) dalam backlog dinyatakan selesai apabila memenuhi kriteria berikut:
1. Kode sumber telah diintegrasikan, lolos validasi kompilasi TypeScript dan SvelteKit (`npm run build`) tanpa galat.
2. Tata letak responsif telah diverifikasi bekerja dengan sempurna pada resolusi layar ponsel pintar, tablet, dan komputer desktop.
3. Seluruh elemen interaktif mematuhi standar ergonomi sentuh minimal 44 pt dan lulus uji kontras teks WCAG AA.
4. Perubahan telah dicatat dalam riwayat commit Git dan diunggah ke repositori pusat GitHub (`main`).
5. Dokumentasi teknis dan bukti visual (*screenshots*) telah diperbarui dan dilaporkan ke pembimbing/instansi terkait.
