### cors 確認準備
#### オリジンサーバー(raspberry pi)

- インターリンクのDNSでドメイン名ddddddo.workのレコードに以下を追加登録
  - ホスト名:cors
  - タイプ:A
  - データー:133.232.127.225
- `/var/www/cors/index.html` に work/cors/index.html 配置
- `/usr/local/etc/h2o/h2o.conf` に、work/cors/h2o.confのcors.ddddddo.work のとこ追加
- 起動中の h2o プロセスを kill -> `sudo /home/pi/h2o/h2o -c /usr/local/etc/h2o/h2o.conf` で再起動して設定反映
- cors.ddddddo.work がオリジンサーバーのホスト
- ブラウザで、http://cors.ddddddo.work を開く

#### オリジンサーバー(local)
- work/cors直下で`open index.html`　実行

#### APIサーバー(CloudFunction)
- go。work/cors/functions.goのコードをCloudFunctionへコピペ。CloudFunctionで登録する名前はCORSEnabledFunction
- (CORS -> "Access-Control-Allow-Origin" = "http://cors.ddddddo.work" を設定している)

### cors 確認結果
- オリジンサーバー(raspberry pi)からDLしたhtml -> APIサーバー へのfetch成功
- オリジンサーバー(local) -> APIサーバー へのfetch失敗

また、
- GCEを起動(nginxでwork/cors/index.htmlを公開)し、API側でGCEのグローバルIPをORIGINに設定(`http://cors.ddddddo.work`はORIGINから外した状態)し、以下を確認した。
  - オリジンサーバー(raspberry pi)のドメイン(http://cors.ddddddo.work)をブラウザで入力し、出力されたindex.html -> APIでのfetchは失敗
  - オリジンサーバー(GCE)のグローバルIPをブラウザで入力し、出力されたindex.html -> APIでのfetchは成功

### ref
[APIサーバを立てるためのCORS設定決定版](https://qiita.com/hirohero/items/886733f50f37404235db)

2020/03/22