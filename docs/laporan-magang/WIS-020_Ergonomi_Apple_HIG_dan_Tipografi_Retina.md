# Dokumen Penugasan & Laporan Teknis: WIS-020
**Nama Pekerjaan:** Penyelarasan Standar Ergonomi Antarmuka Apple HIG  
**Instansi:** Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan  
**Penyusun:** Mutiah Arinil Fayza Nusar  
**Periode Pelaksanaan:** 21 – 23 November 2026  
**Status:** Selesai (100%)  
**Repositori:** https://github.com/mutiaharn/SatpolPP (Commit: `7409ab0`)

---

## 1. Latar Belakang & Tujuan
Menyesuaikan seluruh elemen antarmuka pengguna agar mematuhi standar ergonomi sentuhan dan kejelasan visual kelas dunia berdasarkan *Apple Human Interface Guidelines (HIG)*.

## 2. Rincian Implementasi & Pekerjaan
1. **Kalibrasi Target Sentuh (Minimum 40–44 pt):**
   - Menyesuaikan ukuran interaksi tombol navigasi desktop (`min-h-[40px] px-3.5`).
   - Memadatkan tombol menu hamburger mobile dan tombol pencarian menjadi `44×44px`.
   - Mengubah indikator navigasi carousel beranda menjadi kapsul sentuh ergonomis `32×32px` / `40px`.
   - Menstandarisasi tombol filter kategori (*filter pills*) menjadi `min-h-[40px] px-4 py-2.5`.
   - Memperbesar tombol kontrol paginasi tabel/berita menjadi `36×36px` / `38px`.
2. **Tipografi Native Apple SF Pro:**
   - Mengonfigurasi font stack native: `-apple-system, BlinkMacSystemFont, "SF Pro Display", "SF Pro Text"`.
   - Mengaktifkan rendering teks tingkat lanjut: `-webkit-font-smoothing: antialiased; text-rendering: optimizeLegibility;`.
3. **Mikro-Interaksi Taktil Pegas:**
   - Menerapkan efek umpan balik taktil saat ditekan (`active:scale-[0.975]`) dengan kurva transisi pegas alami (`cubic-bezier(0.16, 1, 0.3, 1)`).
4. **Aksesibilitas Cincin Fokus:**
   - Menyematkan cincin fokus visual ganda (*specular focus halo ring*) untuk navigasi papan ketik (WCAG AA).

## 3. Berkas Kode & Deliverable
* `frontend/src/app.css`
* Seluruh berkas navigasi, slider, dan tombol di `frontend/src/`
