fmt: rmtestdir
	go fmt ./...

rmtestdir:
	rm -rf ./root/ ./root1/ ./root2/ ./root3/ ./root4/ ./root5/ ./root6/ ./root7/ Primate/

test: rmtestdir
	go clean -testcache
	go test ./... -race -v

cyclo: rmtestdir
	gocyclo .

all: fmt test cyclo