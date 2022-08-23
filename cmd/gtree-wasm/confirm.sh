#!/bin/sh

echo -n "service_worker.js の CACHE_VERSION 変数を更新しましたか? [y/N]: "
read ANS

case $ANS in
  [Yy]* )
    return 0
    ;;
  * )
    return 1
    ;;
esac