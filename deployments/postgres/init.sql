-- PostgreSQL Initial Bootstrap for Satpol PP Sulsel & SIAP TERTIB
-- Executed automatically on first container volume initialization

ALTER DATABASE satpolpp_sulsel_db SET timezone TO 'Asia/Makassar';

CREATE EXTENSION IF NOT EXISTS pgcrypto;
