#!/bin/sh
app=$(pwd)
pkgFile="server.go"
outputPath="bin"
output="$outputPath/server"
src="$app/$pkgFile"

printf "\nBuilding: $app\n"
time go build -o $output $src
printf "\nBuilt: $app size:"
ls -lah $output | awk '{print $5}'
printf "\nDone building: $app\n\n"