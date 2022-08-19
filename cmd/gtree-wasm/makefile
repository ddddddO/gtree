serve:
	go run server.go

build:
	GOOS=js GOARCH=wasm go build -o main.wasm main.go

# NOTE:
# go-tomlとyaml.v3を除けばtinygoでビルド可能。
# その時のサイズが700KB弱だったので、goでのビルド時の3MBよりかなり削減できる。
# goでビルドした.wasmがhttpレスポンス時にgzip圧縮されて約1MBだったので、tinygoビルド版では230KBくらいになると推測。かなりサイズを小さく配信できそう。
# ただ、goのwasm_exec.jsではなく、tinygo/targets/wasm_exec.js を使用するとのこと。ref: https://tinygo.org/docs/guides/webassembly/#how-it-works
# tinybuild:
# 	tinygo build -o main.wasm -target wasm main.go

deploy: build
	cp index.html main.wasm wasm_exec.js sitemap.xml ../../docs