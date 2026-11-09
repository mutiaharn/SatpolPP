# Dokumen Desain Mockup Figma Form Lapor & Resi Tiket
**Instansi:** Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan  
**Kode Task Magang:** WIS-008  
**Penyusun:** Mutiah Arinil Fayza Nusar  
**Repositori Proyek:** https://github.com/mutiaharn/SatpolPP  

---

## 1. Konsep Desain Antarmuka Form Lapor
Desain antarmuka formulir pengaduan SIAP TERTIB disusun dengan pendekatan *mobile-first clean form* yang terbagi menjadi 3 langkah berurutan:
1. **Identitas Pelapor & Opsi Anonim:** Desain kartu NIK, nama, dan no HP dengan toggle proteksi anonimitas.
2. **Kategori Pelanggaran & Lokasi:** Pilihan radio button bertema Cokelat Praja dan dropdown wilayah terstruktur.
3. **Bukti Lapangan & Kronologi:** Kotak drag-and-drop foto bukti kejadian dengan tombol pratinjau thumbnail.

## 2. Desain Kartu Resi Tiket Digital
Setelah pengiriman berhasil, sistem menampilkan kartu resi digital yang menyerupai surat tanda penerimaan laporan resmi:
- Kode Resi berukuran besar dengan font monospace (`TTB-YYYYMMDD-XXXX`).
- Badge status SOP interaktif (*Menunggu Verifikasi* berlatar amber).
- Tombol aksi satu klik untuk mencetak resi atau membuka halaman pelacakan status tiket.
