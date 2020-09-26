#!/bin/bash

set -e

echo "We out here in $(pwd)"

echo "==> Killing the old Musick.."
./kill.sh

echo "==> Deleting the old binary..."
rm -rf ../../service/musick

mkdir -p ../../service
mv ../musick ../../service

cd ../../service

echo "==> Starting the new Musick!"
./musick & 