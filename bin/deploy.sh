#!/usr/bin/env bash

set -e

function create_containers() {
    docker-compose up -d --build
}

function update_containers() {
    docker-compose build hermes
    docker-compose up --no-deps -d hermes
    # docker rmi $(docker images -a --filter=dangling=true -q)
    # docker rm $(docker ps --filter=status=exited --filter=status=created -q)
    docker system prune -a -f
}

if [ "$1" == "-u" ]; then
    echo "`date`: hermes/deploy: Updating containers"
    update_containers
    echo "`date`: hermes/deploy: Containers updated"
    exit 0
else
    echo "`date`: hermes/deploy: Buidling and creating containers"
    create_containers
    echo "`date`: hermes/deploy: Containers created"
    exit 0
fi
