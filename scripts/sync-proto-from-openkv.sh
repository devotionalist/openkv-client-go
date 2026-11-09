#!/usr/bin/env bash
# Optional: sync IDL from a local openkv checkout (sibling ../openkv).
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
SRC="${OPENKV_ROOT:-$ROOT/../openkv}/proto/kv/v1/kv.proto"
DST="$ROOT/proto/kv/v1/kv.proto"
cp "$SRC" "$DST"
perl -i -pe 's|option go_package = ".*?";|option go_package = "github.com/devotionalist/openkv-client-go/gen/kv/v1;kvv1";|' "$DST"
echo "synced $DST from $SRC"
