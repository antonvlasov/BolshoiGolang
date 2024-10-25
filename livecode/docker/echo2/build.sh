#!/bin/bash

docker build --progress=plain -t echo:0.0.2 -f ./docker/echo2/Dockerfile.echo .