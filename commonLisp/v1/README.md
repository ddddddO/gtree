## 環境  
WSL Debian  

## 導入  
### linuxbrew  
https://docs.brew.sh/Homebrew-on-Linux  

test -d ~/.linuxbrew && eval $(~/.linuxbrew/bin/brew shellenv)  
test -d /home/linuxbrew/.linuxbrew && eval $(/home/linuxbrew/.linuxbrew/bin/brew shellenv)  
test -r ~/.zshrc && echo "eval \$($(brew --prefix)/bin/brew shellenv)" >>~/.zshrc  
echo "eval \$($(brew --prefix)/bin/brew shellenv)" >>~/.profile  

### Common Lisp環境導入  
https://qiita.com/t-sin/items/054c2ff315ec3b9d3bdc  

### インタプリタ  
起動：ros run  
デバッグ抜け方：abort  

### Lispファイル実行  
ros -l hellolisp.lisp  
http://blog.8arrow.org/entry/2015/06/11/101511  

### エディタ  
lem 未(ros install cxxxr/lem)  
http://blog.8arrow.org/entry/2018/08/14/213428  

## 逆引きCommon Lisp  
https://lisphub.jp/common-lisp/cookbook/  
