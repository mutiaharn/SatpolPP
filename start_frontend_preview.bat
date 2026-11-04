@echo off
title Portal Satpol PP Sulsel + SIAP TERTIB Remake Preview
color 0F
cls
echo =========================================================================
echo    PORTAL TERPADU SATPOL PP PROV. SULSEL & SIAP TERTIB (REMAKE SVELTEKIT)
echo =========================================================================
echo.
echo  [1] Membuka Frontend SvelteKit di http://localhost:3000 ...
echo  [2] Fitur Siap Diuji:
echo      - Beranda Portal Publik & Slider Kegiatan
echo      - Warta Berita & Detail Berita (/berita)
echo      - Dokumen & Regulasi Kinerja Publik (/dokumen)
echo      - Profil Satpol PP & 5 Bidang Kerja (/profil/* & /bidang/*)
echo      - Galeri Dokumentasi Kegiatan (/galeri)
echo      - Formulir Pengaduan Elektronik SIAP TERTIB (/lapor)
echo      - Pelacakan Status Tiket Aduan Realtime (/lacak)
echo      - Dashboard Penanganan Aduan & SOP Staff Satpol (/admin/dashboard)
echo      - Konsol Pengembang & Tata Kelola CMS Superadmin (/superadmin/dashboard)
echo.
echo  Tekan Ctrl+C di terminal ini untuk menghentikan server lokal.
echo =========================================================================
echo.

cd /d "%~dp0frontend"
npm run dev -- --open
pause
