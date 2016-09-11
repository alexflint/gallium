#!/bin/bash

ROOT=$(dirname $(dirname $(dirname $0) ) )
SRC=$ROOT/vendor/brightray/vendor/download/libchromiumcontent/Release/chrome_sandbox
DST=$ROOT/out/Debug/chrome-sandbox

if [ $EUID != 0 ]; then
  echo You must be root to run this script.
  exit 1
fi

rm -f "$DST"               &&
  cp "$SRC" "$DST"         &&
  chown root.root "$DST"   &&
  chmod 4755 "$DST"        || (
  echo Could not copy "$SRC" to "$DST" setuid.
  rm -f "$DST"
  exit 1
)
