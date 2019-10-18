#!/bin/sh
buildPath="build"
app="heritago"
program="$buildPath/$app"

printf "\nStart app: $app\n"
export $(grep -v '^#' .env | xargs)
time ./$program
unset $(grep -v '^#' .env | sed -E 's/(.*)=.*/\1/' | xargs)
printf "\nStopped app: $app\n\n"
