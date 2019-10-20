#! /bin/bash

set -ex

function down() {
	docker-compose -f deployments/01_composes/pubsub/docker-compose.yml down
}

function cleanup() {
	docker rmi postgres_pub -f
	docker rmi postgres_sub1 -f
	docker rmi postgres_sub2 -f
}

down
cleanup
