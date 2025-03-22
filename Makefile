# After the command below: ./gtree template
hotreload:
	air

# MEMO: ver/rev を入れて、手動でGitHub releases にアップロードする
# sudo GOOS=wasip1 GOARCH=wasm go build -ldflags '-X main.Version=1.9.6 -X main.Revision=xxx' -o gtree.wasm ./cmd/gtree
wasi:
	GOOS=wasip1 GOARCH=wasm go build -o gtree.wasm ./cmd/gtree

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
	go test . ./markdown/... -race -coverprofile=coverage.out -covermode=atomic -v

view_cover: sweep cover
	go tool cover -html=coverage.out -o coverage.html

bench: sweep
	go test -benchmem -bench Benchmark -benchtime 100x benchmark_simple_test.go
	go test -benchmem -bench Benchmark -benchtime 100x benchmark_pipeline_test.go

cyclo: sweep
	gocyclo .

credit:
	gocredits . > CREDITS

tapeogp:
	LS_COLORS='di=32:fi=01;34' vhs assets/demo_ogp.tape

tape:
	LS_COLORS='di=32:fi=01;34' vhs assets/demo.tape

tapeall: tapeogp tape

treemap: cover
	go-cover-treemap -statements -coverprofile coverage.out > assets/test_treemap.svg

all: fmt test bench cyclo