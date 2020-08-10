gen:
	go build -o syscaller

runtcp:
	./check.sh tcp

runfile:
	./check.sh file

# tcpdump:https://kazmax.zpp.jp/cmd/t/tcpdump.1.html
td:
	sudo tcpdump -i lo -x -X -vvv -nnn