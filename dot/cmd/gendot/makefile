CNT=11

preopen: gendot svg

gendot:
	go run main.go --digraph --cnt $(CNT) > dot.dot

# gendot後、.dotのnodeの最終行を, -> ;に修正する
svg:
	dot -Tsvg dot.dot -o dot.svg
