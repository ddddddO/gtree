editcss:
	npx tailwindcss -i ./tailwind_base.css -o ./main.css --watch

serve:
	cd ../../docs && python3 -m http.server 8080

# tinygo の version up したら
cpwasmexec:
	cp $(shell tinygo env TINYGOROOT)/targets/wasm_exec.js .

confirm:
	./confirm.sh

# build:
# 	GOOS=js GOARCH=wasm go build -o main.wasm main.go

# NOTE(old):
# go-tomlとyaml.v3を除けばtinygoでビルド可能。
# その時のサイズが692KBだったので、goでのビルド時の3MBよりかなり削減できる。
# goでビルドした.wasmがhttpレスポンス時にgzip圧縮されて約1MBだったので、tinygoビルド版では230KBくらいになると推測。かなりサイズ小さく配信できそう。
# -> gzip圧縮で268KBだった。
# ただ、goのwasm_exec.jsではなく、tinygo/targets/wasm_exec.js を使用する。ref: https://tinygo.org/docs/guides/webassembly/#how-it-works
tinybuild:
	tinygo build -o main.wasm -target wasm -tags tinywasm main.go

deploy: confirm tinybuild
	cleancss -o ../../docs/main.css main.css
	cleancss -o ../../docs/toast.css toast.css
	uglifyjs main.js --compress drop_console=true -o ../../docs/main.js
	uglifyjs toast.js --compress drop_console=true -o ../../docs/toast.js
	uglifyjs tab.js --compress drop_console=true -o ../../docs/tab.js
	uglifyjs wasm_exec.js --compress drop_console=true -o ../../docs/wasm_exec.js
	cp index.html service_worker.js robots.txt sitemap.xml ../../docs
	mv main.wasm ../../docs
