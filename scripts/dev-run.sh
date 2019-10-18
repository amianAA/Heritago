#!/bin/sh
app="heritago"

printf "\nStart running: $app\n"
export $(grep -v '^#' .env | xargs)
time ../build/${app} start run
printf "\nStopped running: $app\n\n"