---
type: source
title: "Deploy: baca semua config sebelum action"
slug: deployment-read-all-before-action
status: insight
created: 2026-07-15
updated: 2026-07-15
category: devops
---
# Deploy: baca semua config sebelum action
## Lesson

Saat diminta deploy, jangan jalankan script deploy langsung tanpa baca konfigurasi dan semua script terkait.

### Checklist wajib sebelum deploy:

1. **Baca `.deploy`** — berisi `SERVER_USER`, `SERVER_HOST`, `SERVER_PATH`, `SERVICE_NAME`
2. **Baca semua script di `scripts/`** — `deploy.sh` (on-server), `update-deploy.sh` (remote restart via SSH), `first-deploy.sh` (initial setup)
3. **Cek `.gitignore` untuk file build-critical** — `go.sum` di-ignore tapi crucial buat Go build
4. **Cek `AGENTS.md` bagian Gotchas** — ada catatan soal `go.sum`, `.vite-port`, dll
5. **Cek status service sebelum restart** — biar tau kondisi awal

### Alur deploy yang benar:

```bash
# 1. Source config
source .deploy

# 2. SSH + build + restart dari lokal (update deploy)
ssh "$SERVER_USER@$SERVER_HOST" "cd $SERVER_PATH && git pull && npm run build:all && sudo systemctl restart $SERVICE_NAME"
```

Atau jalanin `scripts/deploy.sh` langsung di server via SSH interaktif.

*Category: devops*
---
*Captured: 2026-07-15*
## Related
_Add links to related pages._