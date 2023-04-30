#!/bin/bash -e

hash docker

CONTAINER_NAME="jester-db"

# Check if the container exists
if docker container ls -a --format '{{.Names}}' | grep -q "^${CONTAINER_NAME}\$"; then
    echo "Container '${CONTAINER_NAME}' exists."

    # Check if the container is stopped
    if docker container ls --filter status=exited --format '{{.Names}}' | grep -q "^${CONTAINER_NAME}\$"; then
        echo "Starting '${CONTAINER_NAME}' container."
        docker start "${CONTAINER_NAME}"
    else
        echo "Container '${CONTAINER_NAME}' is already running."
    fi
else
    echo "Creating and starting new '${CONTAINER_NAME}' container."
    docker run \
        --name "${CONTAINER_NAME}" \
        --detach \
        -p 5432:5432 \
        -e POSTGRES_PASSWORD=password \
        -e POSTGRES_USER=user \
        -e POSTGRES_DB=jester \
        postgres:15.2
fi
