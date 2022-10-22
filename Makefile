sweep:
	rm -rf ./root/ ./root1/ ./root2/ ./root3/ ./root4/ ./root5/ ./root6/ ./root7/ Primate/ gtree/

fmt: sweep
	go fmt ./...

lint: sweep
	golangci-lint run

test: sweep
	go clean -testcache
	go test . -race -v

cyclo: sweep
	gocyclo .

credit:
	gocredits . > CREDITS

all: fmt test cyclo