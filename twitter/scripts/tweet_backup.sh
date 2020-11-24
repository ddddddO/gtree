#!/bin/bash

# 認証設定
# twurl authorize --consumer-key <key> --consumer-secret <secret>
# twurl accounts
# twurl set dddddO60664252 <key>

# ユースケース
# ・リツイート/ツイートしたツイートを定期的に保存する
# ・保存したリツイート/ツイートを確認する
#   -> ・保存したリツイート/ツイートを(複数選択して)まとめる
#   -> ・まとめたリツイート/ツイートをtag-mngへ反映する

# 課題
# ・個別のツイートで見切れるものがある
# ・リツイートしたツイートにぶら下がってるツイート(リツイートしてない)も欲しい
# ・現状、取得リツイートが全然足りてない

CURRENT_DATE=`date "+%Y%m%d_%H%M"`
twurl -H "api.twitter.com" "/1.1/statuses/user_timeline.json" | \
    jq . > data/${CURRENT_DATE}_timeline.json

cat data/${CURRENT_DATE}_timeline.json | jq .[].text > data/${CURRENT_DATE}_extract.txt
