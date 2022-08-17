serve:
	go run server.go

build:
	GOOS=js GOARCH=wasm go build -o main.wasm main.go

deploy:
	cp index.html main.wasm wasm_exec.js ../../../ddddddO.github.io/