#! /bin/bash


set -x

# https://qiita.com/kite_999/items/e77fb521fc39454244e7
makefile() {
	cat <<EOS > x.txt
red
	blue
		yellow
EOS
}

# ファイル・ディレクトリ存在チェック
# https://www.server-memo.net/shellscript/file_check.html
if [ ! -f x.txt ]; then
	makefile
fi

