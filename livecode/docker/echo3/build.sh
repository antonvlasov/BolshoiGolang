#!/bin/bash

docker build --progress=plain -t echo:0.0.3 -f ./docker/echo3/Dockerfile.echo .