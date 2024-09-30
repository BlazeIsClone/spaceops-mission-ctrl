#!/bin/bash

if ! command -v docker-compose &>/dev/null; then
    echo "Error: docker-compose is not installed. Aborting."
    exit 1
fi

cd deployments/local

if [ "$1" == "up" ]; then
    docker-compose -f docker-compose.yml up -d
elif [ "$1" == "down" ]; then
    docker-compose -f docker-compose.yml down
else
    echo "Usage: $0 [up|down]"
    exit 1
fi