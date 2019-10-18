#!/bin/sh
app="heritago"

printf "\nBuilding: ${app}\n"
time go build -o ../build/${app} ../main.go
printf "\nBuilt: ${app} size:"
ls -lah ../build | awk '{print $5}'
printf "\nDone building: ${app}\n\n"