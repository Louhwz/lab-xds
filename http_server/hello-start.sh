#!/bin/sh

./code/hello-envoy &
./code/lightest &
envoy -c /etc/hello-envoy-config.yaml --service-cluster service_bbc
