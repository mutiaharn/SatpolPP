# Dokumen Penugasan & Laporan Teknis: WIS-022
**Nama Pekerjaan:** Refaktorisasi UI/UX Apple HIG, Reduksi Palet 4 Warna, Restrukturisasi Navigasi, dan Autoscroll Kontekstual  
**Instansi:** Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan  
**Penyusun:** Mutiah Arinil Fayza Nusar  
**Periode Pelaksanaan:** 25 November 2026  
**Status:** Selesai (100%)  
**Repositori:** https://github.com/mutiaharn/SatpolPP (Commit: `cb419a9`)

---

## 1. Latar Belakang & Tujuan
Menyempurnakan antarmuka publik dan modul SIAP TERTIB agar memenuhi standar estetika minimalis modern kelas dunia sesuai *Apple Human Interface Guidelines (HIG)*, menyederhanakan palet warna visual maksimal 4 warna utama, menata ulang hierarki navigasi header & footer, memperbaiki bug interaksi tombol (*dodge button*), serta mengintegrasikan *smooth contextual autoscroll* untuk kelancaran navigasi warga.

## 2. Rincian Implementasi & Pekerjaan
1. **Pustaka Komponen Modular Apple HIG & Refaktorisasi Antarmuka:**
   - Membangun komponen UI atomik baru: `Button.svelte`, `IconButton.svelte`, `GlassCard.svelte`, `Input.svelte`, `StatusBadge.svelte`, dan `CopyButton.svelte`.
   - Mengadopsi prinsip Apple HIG: radius melengkung halus (*continuous corners*), tipografi SF Pro Display/Text, mikro-umpan balik pegas (*tactile spring feedback*), dan target sentuh minimal 44 pt.
   - Merefaktor halaman autentikasi (`/admin/login`), dasbor petugas (`/admin/dashboard`), dasbor komando (`/superadmin/dashboard`), serta formulir pengaduan (`/lapor`).

2. **Pembersihan Visual & Penegakan Palet Maksimal 4 Warna:**
   - Mengonsolidasikan seluruh palet warna website menjadi maksimal 4 warna resmi:
     1. `#F08300` (Praja Amber - Aksen Utama & Status Pending)
     2. `#0B132A` (Deep Obsidian Navy - Latar Belakang & Identitas Korps)
     3. `#FFFFFF` / `#F8FAFC` (Pure & Off-White - Tipografi & Kontras)
     4. `#64748B` / `#E2E8F0` (Slate Neutral - Border & Teks Sekunder)
   - Menghapus efek specular halo artefak kotak, border kaku, dan ambient mesh multi-warna (eliminasi sisa warna biru/sky/indigo).

3. **Penyederhanaan Arsitektur Navigasi Portal:**
   - **Minimalisasi Header:** Menghapus menu "BERITA" dan "DOKUMEN" dari navbar utama karena sudah tersemat langsung di dasbor beranda; menghapus tautan eksternal "Layanan PPID" dari bilah kontak atas.
   - **Restrukturisasi Footer:** Memindahkan direktori menu "PROFIL" dan "BIDANG" ke footer menggantikan kolom tautan dan regulasi lama untuk keseimbangan informasi.

4. **Perbaikan Bug Interaksi & Responsivitas Viewport Laptop:**
   - **Fix Dodge Button:** Mengatasi bug tombol "Jelajahi Portal" yang bergeser ke samping saat diklik dengan memisahkan positioning margin-centering (`left-0 right-0 mx-auto w-max`) dari transform scaling (`active:scale-[0.97]`).
   - **Vertical Centering Seksi Berita:** Menata seksi warta berita di beranda dengan pemusatan vertikal seimbang (`md:min-h-[calc(100vh-96px)] flex flex-col justify-center`) dan memadatkan tinggi kartu agar tidak terpotong taskbar pada layar laptop 1366x768.

5. **Peningkatan UX Pelacakan Tiket & Autoscroll Kontekstual:**
   - **Pelebaran Kartu Lacak Tiket (`/lacak`):** Mengubah tata letak kartu pelacakan dari `max-w-2xl` satu kolom menjadi `max-w-5xl` dua kolom dasbor kompak, memangkas tinggi kartu lebih dari 50%.
   - **Smooth Autoscroll:**
     - Pada `/lacak`: Hasil pelacakan tiket otomatis bergulir pas ke layar saat tombol "Cari Laporan" ditekan, sehingga rincian status, catatan tindak lanjut, dan footer langsung terlihat tanpa perlu scroll manual.
     - Pada `/lapor`: Otomatis scroll mulus ke kartu resi bukti laporan setelah formulir diterbitkan.
     - Pada `/kontak`: Otomatis scroll ke banner konfirmasi pesan terkirim.
   - **Floating Back to Top Button:** Menambahkan tombol melayang untuk kembali ke bagian paling atas halaman dengan transisi halus saat halaman digulir melebihi 300px.

6. **Verifikasi Kompilasi & Repositori:**
   - Pengujian `npm run build` sukses 100% tanpa galat (*zero build error*).
   - Seluruh perubahan (10 commit) telah berhasil di-*push* ke cabang `main` repositori GitHub resmi.

## 3. Berkas Kode & Deliverables
* `frontend/src/lib/components/ui/*` (Button, IconButton, GlassCard, Input, StatusBadge, CopyButton, ScrollToTop)
* `frontend/src/lib/components/layout/Navbar.svelte` & `Footer.svelte`
* `frontend/src/lib/components/complaint/ComplaintForm.svelte` & `TicketTracker.svelte`
* `frontend/src/routes/+page.svelte`, `+layout.svelte`, `/lapor`, `/lacak`, `/kontak`, `/admin/*`, `/superadmin/*`
* Repositori GitHub: `https://github.com/mutiaharn/SatpolPP` (Commit: `cb419a9`)
