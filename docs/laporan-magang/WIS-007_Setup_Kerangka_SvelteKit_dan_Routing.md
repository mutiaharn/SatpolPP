# Dokumen Arsitektur SvelteKit & Sistem Routing
**Instansi:** Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan  
**Kode Task Magang:** WIS-007  
**Penyusun:** Mutiah Arinil Fayza Nusar  
**Repositori Proyek:** https://github.com/mutiaharn/SatpolPP  
**Tautan Source Code:** https://github.com/mutiaharn/SatpolPP/tree/main/frontend/src/routes  

---

## 1. Pilihan Teknologi Frontend
- **Framework:** SvelteKit 2.5 + Svelte 4
- **CSS Engine:** Tailwind CSS 3.4
- **Build Tool:** Vite 5.2
- **State Management:** Svelte Reactive Stores (`appState.ts`)

## 2. Struktur Direktori Routing
```
src/routes/
├── +layout.svelte         # Root layout dengan isolasi Navbar/Footer publik
├── +page.svelte           # Halaman Beranda Utama (Hero 100vh 3D, Warta, Galeri)
├── berita/                # Indeks Warta Berita dengan on-chain search & pagination
│   └── [slug]/            # Detail warta berita komprehensif
├── dokumen/               # Transparansi dokumen publik & regulasi Perda (DataTable + Modal)
├── galeri/                # Galeri foto kegiatan multimedia (Lightbox full navigation)
├── lapor/                 # Formulir pelaporan warga SIAP TERTIB
├── lacak/                 # Pelacakan status tiket pengaduan
├── profil/[slug]/         # Profil instansi, visi misi, dan struktur organisasi
├── admin/                 # Panel operasional petugas piket & verifikator aduan
└── superadmin/            # Konsol administrator sistem Diskominfo
```
