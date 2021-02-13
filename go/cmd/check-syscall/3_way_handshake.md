syscaller/tcpsocket/client.go
```go
package tcpsocket

import (
	"log"
	"net"
)

func RunClient() {
	_, err := net.Dial("tcp", serverAddr+serverPort)
	if err != nil {
		log.Fatal(err)
	}
	//defer conn.Close()

	// _, err = conn.Write([]byte("send from client!"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// bufFromServer := make([]byte, 1024)
	// _, err = conn.Read(bufFromServer)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Print(string(bufFromServer))
}
```

net.Dial関数の呼び出しのみで、3 way handshakeをしていることがわかった。

```sudo tcpdump -i lo -x -X -vvv -nnn
18:32:20 > make td
sudo tcpdump -i lo -x -X -vvv -nnn
tcpdump: listening on lo, link-type EN10MB (Ethernet), capture size 262144 bytes
18:33:51.875834 IP (tos 0x0, ttl 64, id 29978, offset 0, flags [DF], proto TCP (6), length 60)
    127.0.0.1.37590 > 127.0.0.1.8888: Flags [S], cksum 0xfe30 (incorrect -> 0xace8), seq 1318572834, win 43690, options [mss 65495,sackOK,TS val 4271645281 ecr 0,nop,wscale 7], length 
0
        0x0000:  4500 003c 751a 4000 4006 c79f 7f00 0001  E..<u.@.@.......
        0x0010:  7f00 0001 92d6 22b8 4e97 d322 0000 0000  ......".N.."....
        0x0020:  a002 aaaa fe30 0000 0204 ffd7 0402 080a  .....0..........
        0x0030:  fe9c 2261 0000 0000 0103 0307            .."a........
18:33:51.875861 IP (tos 0x0, ttl 64, id 0, offset 0, flags [DF], proto TCP (6), length 60)
    127.0.0.1.8888 > 127.0.0.1.37590: Flags [S.], cksum 0xfe30 (incorrect -> 0x3eb7), seq 3252194121, ack 1318572835, win 43690, options [mss 65495,sackOK,TS val 4271645281 ecr 4271645281,nop,wscale 7], length 0
        0x0000:  4500 003c 0000 4000 4006 3cba 7f00 0001  E..<..@.@.<.....
        0x0010:  7f00 0001 22b8 92d6 c1d8 8b49 4e97 d323  ...."......IN..#
        0x0020:  a012 aaaa fe30 0000 0204 ffd7 0402 080a  .....0..........
        0x0030:  fe9c 2261 fe9c 2261 0103 0307            .."a.."a....
18:33:51.875885 IP (tos 0x0, ttl 64, id 29979, offset 0, flags [DF], proto TCP (6), length 52)
    127.0.0.1.37590 > 127.0.0.1.8888: Flags [.], cksum 0xfe28 (incorrect -> 0x10fc), seq 1, ack 1, win 342, options [nop,nop,TS val 4271645281 ecr 4271645281], length 0
        0x0000:  4500 0034 751b 4000 4006 c7a6 7f00 0001  E..4u.@.@.......
        0x0010:  7f00 0001 92d6 22b8 4e97 d323 c1d8 8b4a  ......".N..#...J
        0x0020:  8010 0156 fe28 0000 0101 080a fe9c 2261  ...V.(........"a
        0x0030:  fe9c 2261                                .."a
18:33:51.885555 IP (tos 0x0, ttl 64, id 1671, offset 0, flags [DF], proto TCP (6), length 52)
    127.0.0.1.8888 > 127.0.0.1.37590: Flags [F.], cksum 0xfe28 (incorrect -> 0x10f2), seq 1, ack 1, win 342, options [nop,nop,TS val 4271645290 ecr 4271645281], length 0
        0x0000:  4500 0034 0687 4000 4006 363b 7f00 0001  E..4..@.@.6;....
        0x0010:  7f00 0001 22b8 92d6 c1d8 8b4a 4e97 d323  ...."......JN..#
        0x0020:  8011 0156 fe28 0000 0101 080a fe9c 226a  ...V.(........"j
        0x0030:  fe9c 2261                                .."a
18:33:51.885598 IP (tos 0x0, ttl 64, id 29980, offset 0, flags [DF], proto TCP (6), length 52)
    127.0.0.1.37590 > 127.0.0.1.8888: Flags [F.], cksum 0xfe28 (incorrect -> 0x10e8), seq 1, ack 2, win 342, options [nop,nop,TS val 4271645290 ecr 4271645290], length 0
        0x0000:  4500 0034 751c 4000 4006 c7a5 7f00 0001  E..4u.@.@.......
        0x0010:  7f00 0001 92d6 22b8 4e97 d323 c1d8 8b4b  ......".N..#...K
        0x0020:  8011 0156 fe28 0000 0101 080a fe9c 226a  ...V.(........"j
        0x0030:  fe9c 226a                                .."j
18:33:51.885613 IP (tos 0x0, ttl 64, id 1672, offset 0, flags [DF], proto TCP (6), length 52)
    127.0.0.1.8888 > 127.0.0.1.37590: Flags [.], cksum 0xfe28 (incorrect -> 0x10e8), seq 2, ack 2, win 342, options [nop,nop,TS val 4271645290 ecr 4271645290], length 0
        0x0000:  4500 0034 0688 4000 4006 363a 7f00 0001  E..4..@.@.6:....
        0x0010:  7f00 0001 22b8 92d6 c1d8 8b4b 4e97 d324  ...."......KN..$
        0x0020:  8010 0156 fe28 0000 0101 080a fe9c 226a  ...V.(........"j
        0x0030:  fe9c 226a                                .."j
```