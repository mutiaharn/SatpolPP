@echo off
title Satpol PP Sulsel Monorepo - Local Development Runner
color 0E
cls
echo =========================================================================
echo    MENJALANKAN MONOREPO SATPOL PP SULSEL & SIAP TERTIB
echo =========================================================================
echo.
echo [1] Menjalankan Backend Golang API (Port 8080)...
start "Backend API (Go)" cmd /k "cd /d %~dp0..\backend && go run cmd\api\main.go"

echo [2] Menjalankan Frontend SvelteKit (Port 3000)...
start "Frontend Web (SvelteKit)" cmd /k "cd /d %~dp0..\frontend && npm run dev -- --open"

echo.
echo Server Backend: http://localhost:8080/api/v1/portal/settings
echo Server Frontend: http://localhost:3000
echo.
echo =========================================================================
