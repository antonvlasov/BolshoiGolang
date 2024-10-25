#!/bin/bash

docker run --rm --name docker-echo --entrypoint '/app/echo' echo:0.0.1 arg45