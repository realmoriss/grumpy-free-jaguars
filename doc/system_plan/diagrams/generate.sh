#!/usr/bin/env sh
set -ex

for diagram in *.puml; do
    plantuml -tpng "${diagram}"
done
