editcss:
	npx tailwindcss -i ./css/tailwind_base.css -o ./css/main.css --watch

serve:
	python3 -m http.server 8080

# build:
# 	GOOS=js GOARCH=wasm go build -o main.wasm main.go

# NOTE:
# go-tomlとyaml.v3を除けばtinygoでビルド可能。
# その時のサイズが692KBだったので、goでのビルド時の3MBよりかなり削減できる。
# goでビルドした.wasmがhttpレスポンス時にgzip圧縮されて約1MBだったので、tinygoビルド版では230KBくらいになると推測。かなりサイズ小さく配信できそう。
# -> gzip圧縮で268KBだった。
# ただ、goのwasm_exec.jsではなく、tinygo/targets/wasm_exec.js を使用する。ref: https://tinygo.org/docs/guides/webassembly/#how-it-works
tinybuild:
	tinygo build -o main.wasm -target wasm -tags wasm main.go

deploy: tinybuild
	cp css/main.css css/toast.css ../../docs/css
	cp -r js/ ../../docs/
	cp index.html service_worker.js robots.txt sitemap.xml ../../docs
	mv main.wasm ../../docs
