#!/usr/bin/env bash
set -exuo pipefail

cd "$( dirname "${BASH_SOURCE[0]}" )/.."
source .envrc

GOOS=windows go build -ldflags="-s -w" -o bin/supply.exe mysql/supply/cli
GOOS=windows go build -ldflags="-s -w" -o bin/finalize.exe mysql/finalize/cli
