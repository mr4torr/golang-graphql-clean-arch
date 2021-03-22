#!/bin/bash
echo "* Migrate"
app="$(pwd)/ops/cmd"
pkgFile="migrate.go"
src="$app/$pkgFile"

export $(grep -v '^#' .env | xargs) && time go run $src
