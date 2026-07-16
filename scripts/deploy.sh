#!/bin/bash
set -e

DIR="$(cd "$(dirname "$0")"/.. && pwd)"

# ── Load SSH config ──
if [ -f "$DIR/.deploy" ]; then
    source "$DIR/.deploy"
else
    echo "❌ .deploy not found — copy .deploy.example to .deploy and fill in your server details"
    exit 1
fi

APP_NAME="${renjana:-laju-go}"
SERVICE="${SERVICE_NAME:-renjana}"

# ── Validate ──
if [ -z "$SERVER_USER" ] || [ -z "$SERVER_HOST" ] || [ -z "$SERVER_PATH" ]; then
    echo "❌ .deploy is missing SERVER_USER, SERVER_HOST, or SERVER_PATH"
    exit 1
fi

SSH_DEST="$SERVER_USER@$SERVER_HOST"

# ── Header ──
echo "🚀 Renjana Deploy"
echo "━━━━━━━━━━━━━━━"
echo "   Server:  $SSH_DEST"
echo "   Path:    $SERVER_PATH"
echo "   Service: $SERVICE"
echo ""

# ── Execute all commands on the remote server ──
ssh -t "$SSH_DEST" "
    set -e

    # Ensure Go and tools are in PATH (non-interactive SSH doesn't source .bashrc)
    export PATH=\"/usr/local/go/bin:\$HOME/go/bin:\$HOME/.local/bin:\$PATH\"

    cd $SERVER_PATH

    echo '📥 Pulling latest...'
    git pull

    echo ''
    echo '📦 Installing deps...'
    npm install 2>&1 | tail -1

    echo ''
    echo '🔨 Building...'
    npm run build:all

    echo ''
    echo '🗄️  Migrating DB...'
    go run github.com/pressly/goose/v3/cmd/goose@latest -dir migrations sqlite ./data/app.db up 2>&1

    echo ''
    echo '🔄 Restarting service...'
    sudo systemctl restart $SERVICE.service

    sleep 2
    ACTIVE=\$(systemctl is-active $SERVICE.service)
    echo ''
    if [ \"\$ACTIVE\" = \"active\" ]; then
        echo '✅ Deploy success! $SERVICE.service is active'
        curl -s -o /dev/null -w \"   HTTP %{http_code} — https://renjana.maulanabuilds.com\\n\" https://renjana.maulanabuilds.com
    else
        echo '❌ Service is \$ACTIVE — check journalctl -u $SERVICE.service'
        exit 1
    fi
"

echo ""
echo "✨ Done."
