#!/bin/bash

set -e

echo "Running Musick!"

./kill.sh

rm -rf ../../service

mkdir -p ../../service

mv musick ../service

cd service

./musick > service.log 2>&1 & 