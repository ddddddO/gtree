up: down
	docker-compose up --build

down:
	docker-compose down

build-api:
	docker build -t short-url-api -f dockerfiles/api/Dockerfile .

test:
	go test -v -coverprofile=cover.out ./...

cover: test
	go tool cover -html=cover.out -o cover.html