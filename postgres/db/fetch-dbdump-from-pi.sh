#! /bin/bash

# DBコンテナに格納するためのテストデータ作成sql取得用sh

set -ex

ssh ochi@ddddddo.work "cd /home/pi/tag-mng/db/bk; pg_dump -U postgres tag-mng > init-dump.sql"
scp ochi@ddddddo.work:/home/pi/tag-mng/db/bk/init-dump.sql /mnt/c/DEV/workspace/GO/src/github.com/ddddddO/work/postgres/db
