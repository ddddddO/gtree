#! /bin/bash

set -ex

function up() {
	docker-compose -f deployments/01_composes/pubsub/docker-compose.yml up
}

up
