@echo off
title Push Monorepo ke GitHub Private (mutiaharn)
color 0A
cls
echo =========================================================================
echo    PUSH KE REPOSITORI GITHUB: mutiaharn/SatpolPP
echo =========================================================================
echo.
echo Identitas Author yang aktif pada repository ini:
git config user.name
git config user.email
echo.
echo Remote origin:
git remote get-url origin
echo.
echo =========================================================================
echo Pilihan Metode Otentikasi:
echo  1. Push Langsung (Menggunakan Git Credential Manager / Browser Login)
echo  2. Push Menggunakan Personal Access Token (PAT GitHub)
echo =========================================================================
echo.
set /p CHOICE="Pilih metode (1 atau 2): "

if "%CHOICE%"=="2" (
    echo.
    set /p GHPAT="Masukkan GitHub Personal Access Token (PAT): "
    if not "%GHPAT%"=="" (
        git remote set-url origin https://mutiaharn:%GHPAT%@github.com/mutiaharn/SatpolPP.git
        echo [*] Melakukan git push...
        git push -u origin main
        echo [*] Mengembalikan URL remote ke standar...
        git remote set-url origin https://github.com/mutiaharn/SatpolPP.git
    )
) else (
    echo.
    echo [*] Melakukan git push (Login browser akan muncul jika belum terotentikasi)...
    git push -u origin main
)

echo.
echo =========================================================================
echo Selesai. Tekan sembarang tombol untuk keluar...
pause >nul
