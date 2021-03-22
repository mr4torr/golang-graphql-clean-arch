#!/bin/sh
app=$(pwd)
pkgFile="server.go"
src="$app/$pkgFile"

printf "\nGenerate running\n"
sh "$app/ops/generate.sh"

printf "\nMigrate running\n"
sh "$app/ops/migrate.sh"

printf "\nStart Server runningapp\n"

# Set all ENV vars for the server to run
export $(grep -v '^#' .env | xargs) && time go run $src
# This should unset all the ENV vars, just in case.
unset $(grep -v '^#' .env | sed -E 's/(.*)=.*/\1/' | xargs)
printf "\nStopped running: $app\n\n"