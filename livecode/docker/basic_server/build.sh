#!/bin/bash

docker build --progress=plain -t basic_server:0.0.1 -f ./docker/basic_server/Dockerfile.server .