#!/bin/sh

set -e

cd "${BUILT_PRODUCTS_DIR}/${1}.framework"
shift

while [ ! -z "${1}" ]; do
  echo "Adding symlink for ${1}"
  ln -sf Versions/Current/"${1}" "${1}"
  shift
done
