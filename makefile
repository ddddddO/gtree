fmt:
	go fmt ./...

test:
	go clean -testcache
	go test ./... -race -v

cyclo:
	gocyclo .

all: fmt test cyclo