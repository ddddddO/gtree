# postgres作業場

## PL/pgSQL

## PubSub
- LISTEN
https://www.postgresql.jp/document/11/html/sql-listen.html

- NOTIFY
https://www.postgresql.jp/document/11/html/sql-notify.html
    - トリガー内でNOTIFYを実行することが可能。だから、テーブルの変更時にアプリに通知することが可能。
    - トランザクション内でNOTIFYする場合、トランザクションがコミットされない限り、通知が送られない。
    - 引数(ペイロード)を送れる(`NOTIFY <channel name>, '<payload>'` で送れる)。
    - 関数`pg_notify(<channel name>,<payload>)` でも可。
    - ペイロードは文字列。デフォルトで8000バイト未満。

- jackc/pgx
https://godoc.org/github.com/jackc/pgx#hdr-Listen_and_Notify

- やりたいこと
    - [x] LISTEN用のアプリを2台用意
    - [x] NOTIFY用のアプリを1台用意
    - [x] NOTIFY実行->LISTEN済みアプリからNOTIFYの内容出力を確認

```log
pub_1   | 2019/10/19 15:07:00 1
pub_1   | 2019/10/19 15:07:10 NOTIFY testpubsub, 'nnnnnotify'
sub1_1  | 2019/10/19 15:07:10 catched notify!! by 1
sub1_1  | 2019/10/19 15:07:10 -notification-
sub1_1  | &{PID:59 Channel:testpubsub Payload:nnnnnotify}
sub1_1  | 2019/10/19 15:07:10 --Channel--
sub1_1  | 2019/10/19 15:07:10 testpubsub
sub1_1  | 2019/10/19 15:07:10 --Payload--
sub1_1  | 2019/10/19 15:07:10 nnnnnotify
sub2_1  | 2019/10/19 15:07:10 catched notify!! by 2
sub2_1  | 2019/10/19 15:07:10 -notification-
sub2_1  | &{PID:59 Channel:testpubsub Payload:nnnnnotify}
sub2_1  | 2019/10/19 15:07:10 --Channel--
sub2_1  | 2019/10/19 15:07:10 testpubsub
sub2_1  | 2019/10/19 15:07:10 --Payload--
sub2_1  | 2019/10/19 15:07:10 nnnnnotify
postgres_sub1_1 exited with code 0
postgres_pub_1 exited with code 0
postgres_sub2_1 exited with code 0
```