# refs
## 「依存関係の検査 : make とシェルスクリプト」
## https://docs.oracle.com/cd/E19620-01/805-5827/6j5gfrand/index.html

## make オプション(-B 無条件に全てのターゲットを make する。...?)
## https://linuxjm.osdn.jp/html/GNU_make/man1/make.1.html#lbAF

############## test 1 ##############
DIR=./a/b

test1: $(DIR)/c.txt test1-alt
	@for i in `seq 2`; \
	do \
		echo 'test1 exec'; \
	done

$(DIR)/c.txt:
	@echo 'Hello!' # @付与でmake実行時、実行コマンド非表示

test1-alt:
	@echo 'test1-alt exec'


############## test 2 ##############
test2:
	@echo "$(dir $@)" # output: ./


############## test 3 ##############
# http://www.jsk.t.u-tokyo.ac.jp/~k-okada/makefile/
## $@ ターゲットファイル名
## 「関数の利用」
### dir ディレクトリ部分の抽出
### wildcard ワイルドカードによりファイル名リストを返す
./a/b/test3:
	@echo "$(dir $@)" # output: a/b/
	@echo $@          # output: a/b/test3


############## test 4 ##############
test4: $(wildcard ./test4/*.txt)
	@echo 'test4 exec'
	@echo $(wildcard ./test4/*.txt) # output: ./test4/3.txt ./test4/1.txt ./test4/2.txt

test4-alt: $(wildcard ./test4/non_exist_dir/*.json)
	@echo 'test4-alt exec'
	@echo $(wildcard ./test4/non_exist_dir/*.json) # output: 


############## test 5 ##############
# http://masahir0y.blogspot.com/2012/06/make-phony-force.html
## .PHONYターゲット .PHONYにターゲットを記載することで、カレントディレクトリに記載したターゲット名と同名のファイル/ディレクトリが存在しても、ターゲットを実行する 

.PHONY: test5 test5-dir test5-file # <- コメントアウトして挙動を確認
test5:
	@echo 'test5'

test5-dir:
	@echo 'test5-dir'


############## test 6 ##############
# || true について
test6:
	@echo 'test6 exec'
	@bash ./scripts/test6.sh || true # test6.sh の実行時必ずエラーになる。「|| true」の付与で、後続の「@echo 'test6 end'」が実行される
	@echo 'test6 end'


############## test 7 ##############
# rsyncについて(以下サイトに、cp と比較したメリット/-aオプション/-vオプション/--deleteオプションについてあり)
# http://x68000.q-e-d.net/~68user/unix/pickup?rsync
.PHONY: test7
test7:
	# 前処理
	@mkdir -p ./test7/a/b
	@touch ./test7/a/b/c.txt
	@mkdir -p ./test7-dest/xxxxx  # 消えるディレクトリ
	@touch ./test7-dest/xxxxx.txt # 消えるファイル

	@echo '--before rsync exec--'
	tree ./test7
	tree ./test7-dest

	@echo '--rsync exec--'
	rsync -av --delete ./test7/ ./test7-dest/

	@echo '--after rsync exec--'
	tree ./test7-dest


############## test 8 ##############
# https://qiita.com/chibi929/items/b8c5f36434d5d3fbfa4a
## error関数 について
## makeコマンド後ろにINPUT＿VERを指定する。ex) make test8 INPUT_VER_1=v1 INPUT_VER_2=v2
INPUT_VER_1=$(error please input ver 1)
INPUT_VER_2=$(error please input ver 2) # 一つでも指定しなければstop
test8:
	@echo 'test8 start'

	@echo $(INPUT_VER_1) $(INPUT_VER_2)

	@echo 'test8 end'


############## test 9 ##############
# -B について
# 以下は、-Bオプションの説明では無いけど、ファイルのタイムスタンプを見てmakeの実行有無が決まるそう。
# -Bを付けると、無条件にすべてのターゲットを実行
## https://linuxcommand.net/make/
test9.txt:
	@echo 'test9 exec'
	echo 'xxxxxx' > test9.txt # ここをコメントアウトしても、すでに実行していれば、「make: 'test9.txt' is up to date.」


############## test 10 ##############
mk-test10:
	touch test10.txt

# test10.txtが存在していれば実行できる
test10: test10.txt
	@echo 'test10 exec'

rm-test10:
	rm test10.txt