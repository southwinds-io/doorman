#!/usr/bin/env bash

# launch ultralight registry
docker run \
    --name uar \
    -p 8082:8080 \
    -e "OX_ADMIN_USER=admin" \
    -e "OX_ADMIN_PWD=admin" \
    -d \
    quay.io/artisan/artr-basic

# launch container registry
docker run \
  --name registry \
  -p 5000:5000 \
  --restart always \
  -d \
  registry:2