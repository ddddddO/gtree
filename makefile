fmt:
	go fmt ./...

test:
	rm -rf ./root/ ./root1/ ./root2/
	go clean -testcache
	go test ./... -race -v

cyclo:
	gocyclo .

all: fmt test cyclo