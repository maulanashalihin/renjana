#!/bin/bash
set -e

# Source profile for non-interactive SSH (Go, cargo, etc.)
if [ -f "$HOME/.bashrc" ]; then
    source "$HOME/.bashrc"
elif [ -f "$HOME/.profile" ]; then
    source "$HOME/.profile"
fi

echo "🚀 Renjana Deploy"
echo "━━━━━━━━━━━━━━━"

# 1. Pull latest
echo "📥 Pulling latest..."
git pull

# 2. Install dependencies (if needed)
echo "📦 Installing deps..."
npm install 2>&1 | tail -1

# 3. Build
echo "🔨 Building..."
npm run build:all

# 4. Migrate (if new migrations exist)
echo "🗄️  Migrating DB..."
~/go/bin/goose -dir migrations sqlite ./data/app.db up 2>&1

# 5. Restart service
echo "🔄 Restarting service..."
sudo systemctl restart renjana.service

# 6. Verify
sleep 1
ACTIVE=$(systemctl is-active renjana.service)
echo ""
if [ "$ACTIVE" = "active" ]; then
    echo "✅ Deploy success! renjana.service is $ACTIVE"
    curl -s -o /dev/null -w "   HTTP %{http_code} — https://renjana.maulanabuilds.com\n" https://renjana.maulanabuilds.com
else
    echo "❌ Service is $ACTIVE — check journalctl -u renjana.service"
    exit 1
fi
