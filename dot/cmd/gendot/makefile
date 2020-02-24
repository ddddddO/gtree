gendot:
	go run main.go --digraph > dot.dot

# gendot後、.dotのnodeの最終行を, -> ;に修正する
svg:
	dot -Tsvg dot.dot -o dot.svg
