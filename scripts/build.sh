#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# Root directory
SUBNET_EVM_PATH=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    cd .. && pwd
)

# Load the constants
source "$SUBNET_EVM_PATH"/scripts/constants.sh

if [[ $# -eq 1 ]]; then
    BINARY_PATH=$1
elif [[ $# -eq 0 ]]; then
    BINARY_PATH="$BUILD_DIR"
else
    echo "Invalid arguments to build subnet-evm. Requires zero (default binary path) or one argument to specify the binary path."
    exit 1
fi

# Get current datetime as version
VERSION=$(date "+%Y%m%d_%H%M%S")

# Build Subnet EVM, which is run as a subprocess
echo "Building Subnet EVM @ Version: $VERSION at $BINARY_PATH"
go build -ldflags "-X github.com/ava-labs/subnet-evm/plugin/evm.Version=$VERSION $STATIC_LD_FLAGS" -o "$BINARY_PATH/$FILE_BUILD_NAME" "plugin/"*.go
echo "Built successfully: $BINARY_PATH/$FILE_BUILD_NAME"
