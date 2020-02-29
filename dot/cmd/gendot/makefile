CNT=11

preopen: gendot svg

gendot:
	CGO_LDFLAGS="`mecab-config --libs`" \
	CGO_CFLAGS="-I`mecab-config --inc-dir`" \
	go run main.go --digraph --cnt $(CNT) > dot.dot

svg:
	dot -Tsvg dot.dot -o dot.svg
	@echo "Next, 'open dot.svg'"
