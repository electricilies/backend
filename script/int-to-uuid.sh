#!/usr/bin/env bash

set -euo pipefail

num="$1"

if [ -z "$num" ]; then
  echo "Usage: $0 <integer>"
  exit 1
fi

# Convert integer to hex (lowercase, no leading 0x)
hex=$(echo "obase=16; $num" | bc | tr '[:upper:]' '[:lower:]')

# Pad to 32 hex chars (128 bits)
hex=$(printf "%032s" "$hex" | tr ' ' '0')

# Format as UUID: 8-4-4-4-12
uuid="${hex:0:8}-${hex:8:4}-${hex:12:4}-${hex:16:4}-${hex:20:12}"

echo "$uuid"
