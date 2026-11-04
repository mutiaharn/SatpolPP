@echo off
echo ==========================================================
echo Menjalankan Backend Golang: Satpol PP Sulsel + SIAP TERTIB
echo ==========================================================
echo.
echo Server berjalan di: http://localhost:8080
echo Health Check      : http://localhost:8080/health
echo API Portal        : http://localhost:8080/api/v1/portal/articles
echo API SIAP TERTIB   : http://localhost:8080/api/v1/siaptertib/lacak/TRB-202611-SAMPLE
echo.
echo Tekan Ctrl+C untuk berhenti.
echo.
api.exe
pause
