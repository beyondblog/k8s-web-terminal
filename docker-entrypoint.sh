#!/bin/bash
set -e

k8s_api=http://127.0.0.17:8080
listen_port=8088
if [ "$K8S_API" ]; then
    k8s_api=$K8S_API
fi

if [ "$LISTEN_PORT" ]; then
    listen_port=$LISTEN_PORT
fi

set -- "$@" --k8s_api=$k8s_api --port=$listen_port

exec "$@"
