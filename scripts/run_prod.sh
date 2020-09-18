#!/bin/bash

set -e

echo "Running Musick!"

cd /home/arkits/software/musick

./musick > service.log 2>&1 & 