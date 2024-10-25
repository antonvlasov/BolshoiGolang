#!/bin/bash

docker build --progress=plain -t echo:0.0.4 -f ./docker/echo4/Dockerfile.echo .