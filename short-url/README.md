# 使い方
## LAN内
### 短縮URL生成サービスを提供する側のセットアップ
1. docker-compose.yaml内の`environment`に、`- SHORTENED_URLS_SVC_HOST=http://<サービス提供するホストの任意な名前>`を記載する。
2. `make up`を実行する。

### 短縮URL生成サービスを利用する側のセットアップ
1. /etc/hostsファイルに、以下を追記する。  
`<サービス提供するホストのローカルIP>   <サービス提供するホストの任意な名前>`  
Windowsは、[こちら](https://www.fonepaw.jp/solution/edit-windows-hosts.html)を参考にして追加する。

2. `curl http://<サービス提供するホストの任意な名前>/surls/ -d 'url=<短縮したいURL>'` を実行する。
3. curl実行結果として、短縮URLが返却される。
