# Dokumen Cakupan Rewrite & Backlog Implementasi
**Instansi:** Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan  
**Kode Task Magang:** WIS-005  
**Penyusun:** Mutiah Arinil Fayza Nusar  
**Repositori Proyek:** https://github.com/mutiaharn/SatpolPP  

---

## 1. Cakupan Minimum Viable Product (MVP)
Rewrite portal Satpol PP Sulsel dirancang untuk menggantikan sistem monolitik lama dengan aplikasi web modern SvelteKit yang cepat, ringan, dan ramah pengguna mobile.

## 2. Komponen UI Reusable yang Diabstraksikan
1. **DataTable.svelte:** Komponen tabel universal dengan filter reaktif *on-chain*, paginasi cerdas, selector baris, dan custom slot.
2. **FileUploader.svelte:** Input unggah gambar berkas foto kejadian & thumbnail berita bergaya Shadcn UI dengan preview thumbnail instan.
3. **Modal.svelte:** Dialog modal berbasis standar aksesibilitas WAI-ARIA untuk inspeksi berkas PDF dan detail pengaduan.
4. **Lightbox.svelte:** Penampil foto galeri multimedia resolusi tinggi dengan navigasi panah kiri/kanan keyboard.
5. **Navbar & Footer:** Komponen tata letak publik yang otomatis terisolasi saat admin berada pada rute internal `/admin/*`.
