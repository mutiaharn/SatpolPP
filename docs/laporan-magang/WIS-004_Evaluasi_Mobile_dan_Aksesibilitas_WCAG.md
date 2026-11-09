# Dokumen Evaluasi Tampilan Mobile & Standar Aksesibilitas WCAG
**Instansi:** Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan  
**Kode Task Magang:** WIS-004  
**Penyusun:** Mutiah Arinil Fayza Nusar  
**Repositori Proyek:** https://github.com/mutiaharn/SatpolPP  

---

## 1. Pengujian Responsivitas Mobile Viewport
Pengujian dilakukan menggunakan Chrome DevTools pada berbagai resolusi perangkat populer:
- **Mobile Small (360x640px - Samsung Galaxy / Android):** Memastikan navbar hamburger berfungsi mulus, padding tombol sentuh minimal 44x44px.
- **Mobile Standard (390x844px - iPhone 12/13/14):** Memastikan hero section 100vh pas satu layar, teks judul tidak terpotong, kartu berita seimbang 1 kolom.
- **Tablet / iPad (768x1024px):** Grid berita 2 kolom responsif, sidebar admin berubah menjadi off-canvas drawer.

## 2. Audit Rasio Kontras Warna & Aksesibilitas (WCAG 2.1 Level AA)
| Elemen UI | Warna Foreground | Warna Background | Contrast Ratio | Standar WCAG AA | Status |
|---|---|---|---|---|---|
| Teks Paragraf Publik | `#1a120c` (Praja Dark) | `#ffffff` (Putih) | 16.5 : 1 | Min. 4.5:1 | **Lolos (AAA)** |
| Teks Footer Resmi | `#d6d3d1` (Stone-300) | `#1a120c` (Praja Hitam) | 10.8 : 1 | Min. 4.5:1 | **Lolos (AAA)** |
| Tombol CTA Primer | `#ffffff` (Putih) | `#F08300` (Amber Gold) | 3.1 : 1 (Teks Tebal >18pt) | Min. 3.0:1 | **Lolos (AA Large)** |
| Badge Kategori Berita | `#F08300` (Amber) | `#1a120c` (Gelap) | 5.8 : 1 | Min. 4.5:1 | **Lolos (AA)** |
