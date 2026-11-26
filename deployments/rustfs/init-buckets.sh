#!/bin/sh
set -eu

echo "[rustfs-init] Waiting for RustFS S3 API at ${S3_ENDPOINT}..."
until curl -sf "${S3_ENDPOINT}/health" >/dev/null 2>&1 || curl -sf "${S3_ENDPOINT}/" >/dev/null 2>&1; do
  sleep 2
done

echo "[rustfs-init] Configuring mc alias..."
mc alias set rustfs "${S3_ENDPOINT}" "${S3_ACCESS_KEY}" "${S3_SECRET_KEY}"

# 1. Private bucket for citizen complaint evidence (NEVER public)
mc mb --ignore-existing "rustfs/${S3_BUCKET_EVIDENCE:-siaptertib-evidence}"
mc anonymous set none "rustfs/${S3_BUCKET_EVIDENCE:-siaptertib-evidence}"

# 2. Public-read bucket for official government documents (Perda, Renja, LAKIP, SOP)
mc mb --ignore-existing "rustfs/${S3_BUCKET_DOCUMENTS:-portal-documents}"
mc anonymous set download "rustfs/${S3_BUCKET_DOCUMENTS:-portal-documents}"

# 3. Public-read bucket for news thumbnails, gallery photos, leadership assets
mc mb --ignore-existing "rustfs/${S3_BUCKET_MEDIA:-portal-media}"
mc anonymous set download "rustfs/${S3_BUCKET_MEDIA:-portal-media}"

echo "[rustfs-init] Buckets initialized cleanly:"
mc ls rustfs/
