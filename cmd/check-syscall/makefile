gen:
	go build -o syscaller

# https://qiita.com/marumaru-n/items/13ef552adad2d3a4ed1e
genpi:
	GOARCH=arm GOARM=7 GOOS=linux go build -o syscaller-pi

runtcp:
	./check.sh tcp

runfile:
	./check.sh file

# tcpdump:https://kazmax.zpp.jp/cmd/t/tcpdump.1.html
td:
	sudo tcpdump -i lo -x -X -vvv -nnn