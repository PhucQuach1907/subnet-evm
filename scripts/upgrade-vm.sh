#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# Root directory
SUBNET_EVM_PATH=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    cd .. && pwd
)

# Stop avalanche node
echo "Stopping avalanche node..."
avalanche node local stop

# Build new version
echo "Building new version..."
./scripts/build.sh

# Get plugin directory
PLUGIN_DIR="$HOME/.avalanchego/plugins"
LOCAL_PLUGIN_DIR="$HOME/.avalanche-cli/local/atichain-local-node-fuji/plugins"

# Run expect script
echo "Updating VM..."
avalanche blockchain upgrade vm atichain

# Get the new plugin file
NEW_PLUGIN=$(ls -t "$PLUGIN_DIR" | head -n1)

# Copy to local plugin directory
echo "Copying plugin to local directory..."
cp "$PLUGIN_DIR/$NEW_PLUGIN" "$LOCAL_PLUGIN_DIR/"

echo "Upgrade completed successfully!" 

echo "Starting avalanche node..."
avalanche node local start atichain-local-node-fuji --avalanchego-path "$HOME/blockchain/avalanchego/build/avalanchego" -f