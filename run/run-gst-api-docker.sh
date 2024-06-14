#!/bin/bash

run() {
    _sync_type=$1
    _docker_registry=$2

    if [ -z $_docker_registry]
    then
        _image="gst-api"
    else
        _image="$_docker_registry/gst-api"
    fi

    docker run --rm -p 4000:4000 \
        $_image:dev gst-api
}

run ${1:-"Full"} ${2:-""}