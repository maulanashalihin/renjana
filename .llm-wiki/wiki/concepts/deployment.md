# Deployment

## Service Info

- Server: `root.laju.dev` (user: `maulana`)
- Path: `/home/maulana/projects/renjana`
- Service: `renjana.service`
- URL: <https://renjana.maulanabuilds.com>
- Config: `.deploy` file di root project

## Files

| File | Fungsi |
|------|--------|
| `.deploy` | Config: server host, user, path, service name |
| `scripts/deploy.sh` | **Jalanin di server langsung** — git pull → build → restart |
| `scripts/update-deploy.sh` | **Jalanin dari lokal** — SSH ke server, restart service aja |
| `scripts/first-deploy.sh` | Setup awal: systemd, env, direktori |

## Deploy Flow (remote dari lokal)

```bash
source .deploy
ssh "$SERVER_USER@$SERVER_HOST" "cd $SERVER_PATH && git pull && npm run build:all && sudo systemctl restart $SERVICE_NAME"
```

## Gotchas

- `go.sum` di-ignore git — setelah `git pull`, jalanin `go mod tidy` atau `go mod download` dulu sebelum build
- Build order: `vite build` dulu, baru `go build` (karena Go binary butuh `dist/.vite/manifest.json`)
- `storage/` dan subdirektori (`avatars/`, `media/`, `documents/`) dibuat otomatis di startup via `os.MkdirAll` di `main.go`
- Jangan langsung jalanin `scripts/deploy.sh` dari lokal — itu script **on-server**. Kalo mau deploy dari lokal, pake SSH manual atau `update-deploy.sh`
- Selalu baca `.deploy` dulu sebelum SSH — jangan tebak IP/host
