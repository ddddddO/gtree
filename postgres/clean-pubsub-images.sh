#! /bin/bash

set -ex

docker rmi postgres_pub -f
docker rmi postgres_sub -f
