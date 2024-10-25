#!/bin/bash

docker run -e BASIC_SERVER_PORT="7500" -p 7500:7500 --rm --name docker-server  basic_server:0.0.1