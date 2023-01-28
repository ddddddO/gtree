sweep:
	rm -rf ./root/ ./root1/ ./root2/ ./root3/ ./root4/ ./root5/ ./root6/ ./root7/ Primate/ gtree/
	rm -rf ./root_a/ ./root_b/ ./root_c/ ./root_d/ ./root_e/ ./root_f/ ./root_g/ ./root_h/ ./root_i/

fmt: sweep
	go fmt ./...

lint: sweep
	golangci-lint run

test: sweep
	go clean -testcache
	go test . -race -v -count=1
	go test ./markdown/... -race -v -count=1

bench: sweep
	go test -benchmem -bench Benchmark tree_handler_benchmark_test.go

cyclo: sweep
	gocyclo .

credit:
	gocredits . > CREDITS

all: fmt test bench cyclo