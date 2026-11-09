# Dokumen Spesifikasi Alur Pengguna (User Flow) & SOP SIAP TERTIB
**Instansi:** Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan  
**Kode Task Magang:** WIS-002  
**Penyusun:** Mutiah Arinil Fayza Nusar  
**Repositori Proyek:** https://github.com/mutiaharn/SatpolPP  

---

## 1. Deskripsi Alur Layanan SIAP TERTIB
Aplikasi SIAP TERTIB dirancang untuk memfasilitasi pelaporan pelanggaran Peraturan Daerah (Perda) dan gangguan ketenteraman dan ketertiban umum (Trantibum) di 24 Kabupaten/Kota se-Sulawesi Selatan secara transparan dan akuntabel.

## 2. Tahapan Alur Pengguna (Citizen Journey)
1. **Pengisian Identitas & Opsi Anonim:**
   - Warga memasukkan NIK (16 digit), Nama Lengkap, dan No. WhatsApp aktif.
   - Tersedia opsi centang **"Lindungi Identitas Saya (Mode Anonim)"** untuk menyamarkan nama dan NIK pada publikasi publik.
2. **Pemilihan Klasifikasi & Titik Kejadian:**
   - Memilih 1 dari 4 jenis pelanggaran: *Gangguan Ketenteraman & Ketertiban Umum, Pelanggaran Perda, Penyerobotan Aset Pemprov, Layanan Linmas & Bencana Sosial*.
   - Memilih Kabupaten/Kota, Kecamatan, Kelurahan, dan menulis alamat detail tempat kejadian.
3. **Unggah Bukti Foto & Uraian Kronologi:**
   - Mengunggah berkas foto bukti kejadian (JPG/PNG/WEBP maks 5MB).
4. **Penerbitan Resi Tiket Digital:**
   - Sistem secara instan menghasilkan kode resi unik berformat: `TTB-YYYYMMDD-XXXX`.

## 3. Penyelarasan dengan Standar Operasional Prosedur (SOP)
| Status SOP | Petugas Berwenang | Durasi Standar | Tindakan |
|---|---|---|---|
| **MENUNGGU_VERIFIKASI** | Petugas Piket / Verifikator | 1 x 24 Jam | Memvalidasi keabsahan data KTP dan bukti gambar pelanggaran |
| **DISPOSISI_BIDANG** | Kepala Seksi / Kabid | 1 x 24 Jam | Meneruskan tiket ke seksi operasional / unit patroli terkait |
| **PENINDAKAN_LAPANGAN** | Regu Patroli Pamong Praja | 1 - 3 Hari | Turun ke lokasi melakukan penghentian, mediasi, atau penyegelan |
| **SELESAI** | Kasatpol / Penyidik PPNS | - | Berita Acara Pemeriksaan (BAP) diterbitkan dan tiket ditutup |
