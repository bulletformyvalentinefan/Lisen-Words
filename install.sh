#!/bin/bash
set -e
BIN="$(go env GOPATH)/bin/lisen"
go build -o "$BIN" .
chmod +x "$BIN"
echo "lisen instalado globalmente en $BIN"
