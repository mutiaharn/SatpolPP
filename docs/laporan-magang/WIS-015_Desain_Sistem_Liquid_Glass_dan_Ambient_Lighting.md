# Dokumen Penugasan & Laporan Teknis: WIS-015
**Nama Pekerjaan:** Perancangan & Implementasi Liquid Glass Design System dan Living Ambient Lighting  
**Instansi:** Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan  
**Penyusun:** Mutiah Arinil Fayza Nusar  
**Periode Pelaksanaan:** 13 – 18 November 2026  
**Status:** Selesai (100%)  
**Repositori:** https://github.com/mutiaharn/SatpolPP (Commit: `56110f3`)

---

## 1. Latar Belakang & Tujuan
Menindaklanjuti kebutuhan pembaruan estetika antarmuka agar tampil lebih modern, prestisius, dan tidak monoton, dirancang sistem desain berbasis material kaca translusen (*Liquid Glass*) dengan saturasi optik tinggi dan tata pencahayaan hidup (*living ambient glow*).

## 2. Rincian Implementasi & Pekerjaan
1. **Token Desain CSS Glassmorphism:**
   - `.ds-glass`: Kaca translusen dengan `backdrop-filter: blur(16px)` dan saturasi 180%.
   - `.ds-glass-strong`: Kaca pekat untuk elemen navigasi melayang dan dialog modal.
   - `.ds-glass-dark`: Material kontras tinggi untuk latar belakang gelap sinematik.
   - `.ds-glass-gold`: Aksen keemasan khas Praja Amber.
   - `.ds-glass-azure`: Aksen biru kehormatan Sovereign Azure.
2. **Specular Inner Rim Highlight:**
   - Menambahkan pantulan cahaya kristal halus (`inset 0 1px 1px 0 rgba(255, 255, 255, 0.25)`) pada tepi atas setiap kartu.
3. **Living Ambient Glow:**
   - Mengintegrasikan lapisan gradasi radial dinamis pada layout global (`+layout.svelte`) untuk menghasilkan kedalaman visual saat menggulir halaman.
4. **Harmonisasi Palet Warna Resmi:**
   - *Deep Obsidian Navy* (`#070e1c`), *Sovereign Azure* (`#246bfd`), dan *Praja Amber* (`#F08300`).

## 3. Berkas Kode & Deliverable
* `frontend/src/app.css` (Definisi token CSS global)
* `frontend/src/routes/+layout.svelte` (Penerapan ambient glow global)
