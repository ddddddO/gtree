# 使い方
## LAN内
### 短縮URL生成サービスを提供する側のセットアップ
1. docker-compose.yaml内の`environment`に、`- SHORTENED_URLS_SVC_HOST=http://<サービス提供するホストの任意な名前>`を記載する。
2. `make up`を実行する。

### 短縮URL生成サービスを利用する側のセットアップ
1. /etc/hostsファイルに、以下を追記する。  
`<サービス提供するホストのローカルIP>   <サービス提供するホストの任意な名前>`  
Windowsは、[こちら](https://www.fonepaw.jp/solution/edit-windows-hosts.html)を参考にして追加する。

2. curl or Web page で短縮URLを生成する。
    - curl
        1. `curl http://<サービス提供するホストの任意な名前>/surls/ -d 'url=<短縮したいURL>'` を実行する。
    - Web page
        1. `http://<サービス提供するホストの任意な名前>` をブラウザのURLバーに入力する。
        2. `URL`フォームに短縮したいURLを入力してENTER
3. 短縮URLが返却される。
