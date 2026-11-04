@echo off
title Setup Git Alias & Remote Repository
color 0B
cls
echo =========================================================================
echo    KONFIGURASI GIT DENGAN IDENTITAS / ALIAS KHUSUS (LOCAL REPO)
echo =========================================================================
echo.
echo Script ini akan:
echo 1. Menginisialisasi Git Repository lokal pada folder monorepo ini (jika belum).
echo 2. Mengatur user.name dan user.email HANYA untuk repository ini (LOCAL).
echo    (Konfigurasi Git Global komputer Anda TIDAK AKAN BERUBAH).
echo 3. Mengaitkan remote origin ke repositori GitHub private Anda.
echo.
echo =========================================================================
echo.

cd /d "%~dp0.."

if not exist ".git" (
    echo [*] Menginisialisasi Git Repository...
    git init
    git branch -M main
) else (
    echo [*] Git Repository sudah terinisialisasi.
)

echo.
set /p GIT_NAME="Masukkan Nama Alias yang diinginkan (Contoh: Tim Pengembang Sulsel): "
set /p GIT_EMAIL="Masukkan Email Alias yang diinginkan (Contoh: dev@sulselprov.go.id): "
set /p REPO_URL="Masukkan URL Git Repo Private Anda (HTTPS atau SSH): "

echo.
echo [*] Menyimpan konfigurasi lokal...
git config user.name "%GIT_NAME%"
git config user.email "%GIT_EMAIL%"

if not "%REPO_URL%"=="" (
    git remote remove origin 2>nul
    git remote add origin %REPO_URL%
    echo [*] Remote origin berhasil diatur ke: %REPO_URL%
)

echo.
echo =========================================================================
echo    KONFIGURASI SELESAI
echo =========================================================================
echo  Identitas Author:
git config user.name
git config user.email
echo.
echo  Remote URL:
git remote -v
echo.
echo Anda dapat melakukan commit pertama dengan perintah:
echo   git add .
echo   git commit -m "feat: initial monorepo release portal satpol pp & siaptertib"
echo   git push -u origin main
echo =========================================================================
pause
