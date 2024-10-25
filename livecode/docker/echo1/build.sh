#!/bin/bash

go build -o ./docker/echo1/echo ./cmd/echo/ && \
docker build -t echo:0.0.1 -f ./docker/echo1/Dockerfile.echo ./docker/echo1