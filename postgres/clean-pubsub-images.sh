#! /bin/bash

set -ex

function down() {
	docker-compose down
}

function cleanup() {
	docker rmi postgres_pub -f
	docker rmi postgres_sub1 -f
	docker rmi postgres_sub2 -f
}

function up() {
	docker-compose up
}

down
cleanup
up
