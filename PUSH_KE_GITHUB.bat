@echo off
title Push Satpol PP & SIAP TERTIB ke GitHub (mutiaharn)
color 0B
cls
echo =========================================================================
echo       PUSH KE REPOSITORI GITHUB: mutiaharn/SatpolPP
echo =========================================================================
echo.
echo Identitas Author Commit:
git config user.name
git config user.email
echo.
echo Remote Origin:
git remote get-url origin
echo.
echo =========================================================================
echo Memulai proses autentikasi via browser...
echo Browser akan otomatis terbuka untuk mengesahkan akun 'mutiaharn'.
echo Anda cukup klik 'Authorize' pada browser Anda.
echo (Asisten AI tidak akan melihat, membaca, atau mengambil kredensial Anda).
echo =========================================================================
echo.
echo Tekan sembarang tombol untuk membuka autentikasi browser dan push...
pause >nul
echo.
echo [*] Melakukan git push...
git push -u origin main
echo.
if %ERRORLEVEL% EQU 0 (
    color 0A
    echo =========================================================================
    echo BERHASIL: Seluruh commit atomic telah terunggah ke repositori GitHub!
    echo URL: https://github.com/mutiaharn/SatpolPP
    echo =========================================================================
) else (
    color 0C
    echo =========================================================================
    echo Gagal atau dibatalkan. Pastikan Anda telah mengotorisasi akun 'mutiaharn' di browser.
    echo =========================================================================
)
echo.
pause
