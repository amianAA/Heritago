#!/bin/bash
printf "\nRegenerating gqlgen files\n"
go run -v github.com/99designs/gqlgen
printf "\nDone.\n\n"