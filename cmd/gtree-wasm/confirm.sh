#!/bin/sh

echo -n "service_worker.js の CACHE_VERSION 変数の変更/ 新規作成したjsファイルを urlsToCache に追加しましたか? [y/N]: "
read ANS

case $ANS in
  [Yy]* )
    return 0
    ;;
  * )
    return 1
    ;;
esac