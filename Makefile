# After the command below: ./gtree template
hotreload:
	air

sweep:
	rm -rf ./root/ ./root1/ ./root2/ ./root3/ ./root4/ ./root5/ ./root6/ ./root7/ ./root8/ Primate/ gtree/
	rm -rf ./root_a/ ./root_b/ ./root_c/ ./root_d/ ./root_e/ ./root_f/ ./root_g/ ./root_h/ ./root_i/ ./root_j/

fmt: sweep
	go fmt ./...

lint: sweep
	golangci-lint run

test: sweep
	go clean -testcache
	go test . -race -v -count=1
	go test ./markdown/... -race -v -count=1

cover: sweep
	go test . -race -coverprofile=coverage.out -covermode=atomic -v
	go tool cover -html=coverage.out -o coverage.html

bench: sweep
	go test -benchmem -bench Benchmark -benchtime 100x benchmark_simple_test.go
	go test -benchmem -bench Benchmark -benchtime 100x benchmark_pipeline_test.go

cyclo: sweep
	gocyclo .

credit:
	gocredits . > CREDITS

tape:
	vhs demo.tape

all: fmt test bench cyclo