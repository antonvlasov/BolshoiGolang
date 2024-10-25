#!/bin/bash

docker build --progress=plain -t basic_client:0.0.1 -f ./docker/server_client/Dockerfile.client .