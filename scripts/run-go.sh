#!/bin/sh
APP="heritago"
BUILD_PATH="build"
ENV="production"

PROGRAM="${BUILD_PATH}/${APP}"

# TASK
printf "\nStart app: ${APP}\n"
export $(grep -v '^#' ../environment/${ENV}/.env | xargs)
time ../${PROGRAM}
unset $(grep -v '^#' ../environment/${ENV}/.env | sed -E 's/(.*)=.*/\1/' | xargs)
printf "\nStopped app: ${APP}\n\n"
