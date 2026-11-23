# Dokumen Penugasan & Laporan Teknis: WIS-017
**Nama Pekerjaan:** Modernisasi Antarmuka Masuk Administrator & CMS Split-Screen  
**Instansi:** Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan  
**Penyusun:** Mutiah Arinil Fayza Nusar  
**Periode Pelaksanaan:** 18 November 2026  
**Status:** Selesai (100%)  
**Repositori:** https://github.com/mutiaharn/SatpolPP (Commit: `7ab582f`)

---

## 1. Latar Belakang & Tujuan
Mengganti tampilan halaman login admin yang sebelumnya sederhana menjadi antarmuka korporat profesional dengan tata letak *split-screen*, menonjolkan identitas kelembagaan Satpol PP dan keamanan sistem informasi.

## 2. Rincian Implementasi & Pekerjaan
1. **Tata Letak Split-Screen:**
   - Panel kiri: Ilustrasi identitas korps praja, lencana resmi beranimasi halus, dan ringkasan fitur proteksi SIAP TERTIB.
   - Panel kanan: Kartu kaca melayang (*floating glass card*) untuk form kredensial petugas.
2. **Ergonomi Form & Validasi:**
   - Kolom isian NIP/Email dan Kata Sandi dengan efek *focus ring* interaktif.
   - Tombol masuk dengan umpan balik visual saat proses autentikasi berlangsung.
3. **Pemisahan Peran:**
   - Dukungan alur login terpisah antara Petugas Pengolah Aduan dan Superadmin Komando.

## 3. Berkas Kode & Deliverable
* `frontend/src/routes/admin/login/+page.svelte`
