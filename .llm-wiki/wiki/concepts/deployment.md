# Deployment

- Git-based: `git pull → make build → sudo systemctl restart laju-go`
- Docker: multi-stage build di `Dockerfile`
- Systemd: service setup di `systemd/laju-go.service`
- Cross-compile: `GOOS=linux GOARCH=amd64 go build` (CGO-free via modernc.org/sqlite)
- `dist/` is gitignored kecuali `.gitkeep`. Build artifacts tidak di-commit.
