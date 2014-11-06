#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
mkdir -p $DIR/data
docker run -d -h `hostname` -p 8080:8080 -p 28015:28015 -p 29015:29015 -v $DIR/data:/data dockerfile/rethinkdb rethinkdb -d /data --bind all --canonical-address `curl icanhazip.com`
