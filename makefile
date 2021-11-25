fmt:
	go fmt

test:
	go test ./... -race -v

cyclo:
	gocyclo .

all: fmt test cyclo