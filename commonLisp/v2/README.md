## docs
- 実践Common Lisp

## 環境
- WSL Debian
- Roswell, SBCL
- Vim, slimv
    - https://qiita.com/kedama17/items/93ae7ccd1839f4bbb567
    - vimとREPLの画面を開くには、`vim xxx.lisp` の前に`tmux`を実行する。
        - `C-w + 矢印キー`で画面移動
    - [slimvショートカット一覧](https://gist.github.com/otaon/d702866b15a9b47bbe19ff1261799dd9)

で進めていく。

<details><summary>実践Common Lispの環境構築できなかったメモ</summary>

## 環境
- WSL Debian
- Lisp in a Box
    - https://common-lisp.net/project/lispbox/ より
        - `curl -OL https://common-lisp.net/project/lispbox/test_builds/lispbox-0.7-ccl-1.6-linuxx86-64.tar.gz` でダウンロード
        - `tar -zxvf lispbox-0.7-ccl-1.6-linuxx86-64.tar.gz` で解凍
        - `cd lispbox-0.7`
        - `./lispbox.sh` でエラー(`./lispbox.sh: line 31: /mnt/c/DEV/workspace/GO/src/github.com/ddddddO/work/commonLisp/v2/lispbox-0.7/emacs-23.2/bin/emacs: No such file or directory`)
        - `cd emacs-23.2/bin/`
        - `mv emacs-23.2 emacs`
        - `cd ../../`
        - `./lispbox.sh`でエラー(`/mnt/c/DEV/workspace/GO/src/github.com/ddddddO/work/commonLisp/v2/lispbox-0.7/emacs-23.2/bin/emacs: error while loading shared libraries: libpng12.so.0: cannot open shared object file: No such file or directory`)
        - ` wget -q -O /tmp/libpng12.deb http://mirrors.kernel.org/ubuntu/pool/main/libp/libpng/libpng12-0_1.2.54-1ubuntu1_amd64.deb   && sudo dpkg -i /tmp/libpng12.deb   && rm /tmp/libpng12.deb`
        - `./lispbox.sh` でエラー(`/mnt/c/DEV/workspace/GO/src/github.com/ddddddO/work/commonLisp/v2/lispbox-0.7/emacs-23.2/bin/emacs: error while loading shared libraries: libgconf-2.so.4: cannot open shared object file: No such file or directory`)
        - `sudo apt-get install libgconf-2-4`
        - `./lispbox.sh`
        - 。。。

</details>

<details><summary>Land of Lispの環境構築できなかったメモ</summary>

## docs
- Land of Lisp
- http://diary.wshito.com/comp/lisp/lisp-pm/

## 環境
- WSL Debian
- Lispコンパイラ CLISP

### CLISPインストール
`sudo apt-get install clisp`でインストールできなかったので、過去にインストールしていたRoswellでclispを使えるようにしようとしたが、エラー。
`ros install clisp`
`ros use clisp` でデフォルトで使用する処理系を設定する。`処理系`=コンパイラ？
(`ros help` でヘルプ)

Land of Lisp 12章からCLISP独自のコマンドを使うらしい。。
</details>