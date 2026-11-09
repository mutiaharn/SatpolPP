# Dokumen Implementasi Pelacakan Tiket (Ticket Tracker)
**Instansi:** Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan  
**Kode Task Magang:** WIS-012  
**Penyusun:** Mutiah Arinil Fayza Nusar  
**Repositori Proyek:** https://github.com/mutiaharn/SatpolPP  
**Tautan Source Code:** https://github.com/mutiaharn/SatpolPP/blob/main/frontend/src/lib/components/TicketTracker.svelte  

---

## 1. Fungsionalitas Modul Pelacak Tiket
- **Pencarian Kode Resi:** Warga dapat memasukkan kode resi tiket (`TTB-YYYYMMDD-XXXX`) untuk mengetahui posisi penanganan laporannya.
- **Visual Stepper SOP:** Menampilkan diagram status bertahap (*Verifikasi &rarr; Disposisi &rarr; Penindakan &rarr; Selesai*) dengan indikator waktu dan nama bidang yang menangani.
- **Proteksi Identitas:** Data pribadi pelapor otomatis disamarkan (*masking*) agar tidak disalahgunakan pihak ketiga.
- **Penanganan Kasus Kosong (Empty State):** Memberikan pesan informatif jika nomor resi tidak terdaftar di sistem.
