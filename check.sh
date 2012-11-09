#!/bin/bash

RED='\033[1;31m'
GREEN='\033[1;33m'
stopColor='\033[0m'

echo 'Running go unit tests + benchmarks'
go test -test.v -test.bench '.*' github.com/Amodio/amazing-gomaze/gomaze
if [ $? -eq 0 ]; then
    echo -e "${GREEN}Success$stopColor"
else
    echo -e "${RED}Failure$stopColor"
fi
