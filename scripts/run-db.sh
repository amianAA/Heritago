#!/bin/sh
APP="heritago"
BUILD_PATH="build"
ENV="develop"

PROGRAM="${BUILD_PATH}/${APP}"

# TASK
printf "\nChecking if PostgreSQL image exists...\n"
if (($(docker image list | grep -c "postgres") > 0)); then
    printf "\n[OK] Image exists\n"
else
    printf "\n[KO] Downloading image...\n"
    docker pull postgres:latest
fi

# TASK
printf "\nChecking DB container exists for '${APP}'...\n"

if (($(docker container list --all | grep -c "${APP}") > 0 )); then
    printf "\n[OK] DB exists. Starting container for ${APP}...\n"
    docker container start ${APP}
else
    printf "\n[KO] Creating DB container for ${APP}...\n"
    docker run --name ${APP} -e POSTGRES_PASSWORD=${APP} -d -p 5432:5432 postgres
fi

printf "\n'${APP} Database is running now'...\n"