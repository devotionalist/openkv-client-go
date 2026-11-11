#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
OUT="$ROOT/gen"
mkdir -p "$OUT"
export PATH="${PATH}:$(go env GOPATH)/bin"
protoc -I "$ROOT/proto" \
  --go_out="$OUT" --go_opt=paths=source_relative \
  --go-grpc_out="$OUT" --go-grpc_opt=paths=source_relative \
  "$ROOT/proto/kv/v1/kv.proto" \
  "$ROOT/proto/kv/v1/ds.proto"
echo "generated $OUT/kv/v1 (kv + ds)"
