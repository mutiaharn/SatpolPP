# Dokumen Inventarisasi Halaman & Fitur Eksisting Website Lama
**Instansi:** Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan  
**Kode Task Magang:** WIS-001  
**Penyusun:** Mutiah Arinil Fayza Nusar  
**Repositori Proyek:** https://github.com/mutiaharn/SatpolPP  

---

## 1. Latar Belakang & Tujuan
Website lama Satpol PP Provinsi Sulawesi Selatan memiliki beberapa keterbatasan struktural, navigasi, dan teknologi. Inventarisasi ini bertujuan untuk memetakan seluruh tautan aktif, arsitektur informasi, konten regulasi, serta kelemahan navigasi situs sebelumnya sebagai acuan perancangan portal baru berbasis SvelteKit & Tailwind CSS.

## 2. Struktur Menu & Daftar URL Halaman Eksisting
| No | Nama Menu / Fitur | URL / Jalur Lama | Status Konten | Catatan Evaluasi |
|---|---|---|---|---|
| 1 | Beranda Utama | `/index.php` | Usang | Tampilan belum responsif, hero slider pecah pada layar mobile |
| 2 | Profil Pimpinan & Visi Misi | `/profil.php` | Arsip | Teks statis, foto pimpinan resolusi rendah dengan watermark lama |
| 3 | Struktur Organisasi | `/struktur.php` | Statis | Diagram bagan berbentuk gambar raster buram, tidak dapat dibaca di layar HP |
| 4 | Regulasi & Dokumen Perda | `/download.php` | Sebagian Broken | Banyak link unduhan PDF 404, tidak memiliki filter kategori atau pencarian |
| 5 | Berita & Kegiatan | `/berita.php` | Parsial | Format tanggal tidak baku, thumbnail kartu tidak proporsional |
| 6 | Pengaduan Warga | `/kontak.php` | Form Email Sederhana | Tidak ada resi tiket, tidak ada fitur lacak, tidak ada perlindungan anonim |

## 3. Identifikasi Masalah & Broken Links
1. **Broken Links & Asset 404:** Sejumlah link dokumen produk hukum daerah dan lampiran SOP mengalami kegagalan akses (error 404).
2. **Navigasi Membingungkan:** Menu bertingkat terlalu dalam dan sulit diakses pengguna perangkat telepon genggam.
3. **Ketiadaan Sistem Tracking Aduan:** Warga yang mengirim aspirasi tidak memperoleh nomor tiket resi untuk memantau proses tindak lanjut petugas di lapangan.

## 4. Rekomendasi untuk Sistem Baru
- Modernisasi antarmuka menggunakan palet resmi Cokelat Praja (`#3d261a`) dan Gold (`#F08300`).
- Implementasi aplikasi mandiri SIAP TERTIB dengan resi tiket instan (`TTB-YYYYMMDD-XXXX`).
- Tabel dokumen publik dinamis dengan fitur *on-chain search* dan paginasi interaktif.
