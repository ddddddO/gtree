# ShellScript基礎
#### 文字列
- シングルコーテーションで囲うと、変数が展開されない

#### 特殊変数
```shell
#! /bin/bash
# コマンド引数
echo $1
echo $2

echo $0 # ファイル名
echo $# # 付与された引数の数
echo $@ # or $* すべての引数
```

#### 入力取得
```shell
#! /bin/bash
read s
echo "hay! $s"
```

```shell
#! /bin/bash
read -p "Name: " s # コンソールに"Name: "を出力
echo "hay! $s"
```

```shell
#! /bin/bash
read -p "Colors: " a b c # 複数入力
echo "color: $a"
echo "color: $b"
echo "color: $c"
```

#### 配列
```shell
#! /bin/bash
list=(a b)
echo ${list[0]}
echo ${list[1]}
echo ${list[@]} # 全要素
echo ${#list[@]} # 個数

list[0]=x
echo ${list[0]}

list+=(zz yy)
echo ${list[@]}
```

#### 計算
```shell
rslt=`expr 9 + 9`
echo $rslt

rslt=$((99 + 99))
echo $rslt
```

#### if
```shell
name="gg"
if [ "$name" = "ff" ]; then
    echo "Name: $name"
elif [ "$name" = "gg" ]; then
    echo "gg"
else
    echo "not found"
fi        
```
二重大かっこ[[]]

#### for
```shell
for i in 1 2 3 4; do
    echo $i
done

for i in {1..5}; do
    echo $i
done

for ((i=0; i<5; i++)); do
    echo $i
done

list=(A B C)
for l in ${list[@]}; do
    echo $l
done

`expr date`
(date)
for s in $(date); do
    echo $s
done
```

#### while
```shell
i=0
while ((i < 6)); do
    ((i++))
    if ((i == 3)); then
        continue
    elif ((i == 4)); then
        break
    fi

   echo $i
done      
```

```shell
while :   # 無限ループ
do
    read -p "please input: " msg
    if [[ $msg == "end" ]]; then
        break
    fi

    echo $msg
done
```

```shell
# 1行毎にlineに格納しながらsample.txtを読み込む
while read line; do
    echo "$line"
done < sample.txt
```

sample
```shell
while read line; do
    echo "$line"
done
```
`cat <ファイル> | ./sample` で、<ファイル>内容を１行毎に出力


#### case
```shell
read -p "prease input color: " color
case "$color" in
    red)
        echo "tomato"
        ;;
    blue|black) # ワイルドカード/部分一致等も可能
        echo "see"
        ;;
    yellow)
        echo "kimi"
        ;;
    *)  # default
        echo "not found.."
esac
```

```shell
select color in red blue yellow black; do # ループ・番号選択
    case "$color" in
        red)
            echo "tomato"
            ;;
        blue|black)
            echo "see"
            ;;
        yellow)
            echo "kimi"
            ;;
        *)
            echo "not found.."
            break
    esac
done  
```


#### 関数
```shell
#h() {        # でも可
function h() {
    if [[ "$1" == "yes" ]]; then  # $1 関数の第一引数
        echo "input $1!"
        return 0  # 終了ステータス(0~255。基本0/1のみでよさそう)
    else
        return 1
    fi
}
#h "no"
h "yes"
```

```shell
function h() {
    local name="sss"  # local を変数に付けなければ、どこからでも参照可
    echo "name is $name"
}
h
echo $name  # 関数内でlocalと付与されているため、nameは空
```