# ref
- よくわかるgRPC
- tcpdumpでパケットキャプチャ
    `sudo tcpdump -i lo -A`
    https://takuya-1st.hatenablog.jp/entry/2019/03/11/120641
    https://orebibou.com/2015/05/tcpdump%E3%82%B3%E3%83%9E%E3%83%B3%E3%83%89%E3%81%A7%E8%A6%9A%E3%81%88%E3%81%A6%E3%81%8A%E3%81%8D%E3%81%9F%E3%81%84%E4%BD%BF%E3%81%84%E6%96%B94%E5%80%8B/

    - キャプチャのファイル出力
        `sudo tcpdump -i lo -A -s 0 -w ./raw_http_dump.pcap`
    - ファイルの読み出し
        `tcpdump -A -r raw_http_dump.pcap`