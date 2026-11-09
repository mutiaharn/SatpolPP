# Dokumen Implementasi Form Pengaduan SIAP TERTIB
**Instansi:** Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan  
**Kode Task Magang:** WIS-010  
**Penyusun:** Mutiah Arinil Fayza Nusar  
**Repositori Proyek:** https://github.com/mutiaharn/SatpolPP  
**Tautan Source Code:** https://github.com/mutiaharn/SatpolPP/blob/main/frontend/src/lib/components/ComplaintForm.svelte  

---

## 1. Fitur Utama Formulir
- **Validasi Ketat Identitas:** Memastikan NIK terdiri dari 16 digit angka sesuai standar KTP elektronik Indonesia.
- **Mode Masking Anonim:** Menyaring tampilan nama dan NIK pada publikasi publik untuk keamanan warga.
- **Peniadaan GPS Otomatis:** Mengganti deteksi GPS otomatis dengan dropdown terstruktur manual (Kabupaten/Kota se-Sulsel, Kecamatan, Kelurahan, Alamat Detail) sesuai arahan pimpinan.
- **Validasi Berkas Gambar:** Hanya menerima file gambar (JPG, PNG, WEBP) maksimal 5MB dengan pratinjau thumbnail instan.
