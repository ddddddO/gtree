## 概要
- オブジェクト指向/スクリプト言語
- すべての値はオブジェクト

## 変数
- 英小文字/_ 始まり

## 定数
- 英大文字始まり。慣習的にすべて大文字にする。
- 再代入して実行したとき、警告されるが処理は止まらない。

## 文字列
### ""
- 特殊文字が使用可能
- 式展開が可能
- '' でも文字列を表現できるが、上記はできない
```ruby
irb(main):002:0> puts "hello\nworld\n#{111 * 5}"
hello
world
555
=> nil
irb(main):003:0> puts 'hello\nworld\n#{111 * 5}'
hello\nworld\n#{111 * 5}
=> nil
```

### %, %Q, %q
```ruby
irb(main):034:0> %(hell"o)
=> "hell\"o"
irb(main):035:0> %Q(hell"o)
=> "hell\"o"
irb(main):036:0> %q(hell"o)
=> "hell\"o"
```


## オブジェクト
- 命令(メソッド)を持ったデータ型。呼び出せるメソッドはオブジェクトの種類により変化
- オブジェクトの種類 = クラス
- 具体的な値 = インスタンス
---
- <オブジェクト>.class で、クラスの種類を返す
```ruby
irb(main):005:0> p "ss".class
String
=> String
```

- <オブジェクト>.methods で、使えるメソッドを返す
```ruby
irb(main):007:0> p 555.methods
[:-@, :**, :<=>, :upto, :<<, :<=, :>=, :==, :chr, :===, :>>, :[], :%, :&, :inspect, :+, :ord, :-, :/, :*, :size, :succ, :<, :>, :to_int, :coerce, :divmod, :to_s, :to_i, :fdiv, :modulo, :remainder, :abs, :magnitude, :integer?, :numerator, :denominator, :to_r, :floor, :ceil, :round, :truncate, :gcd, :to_f, :^, :odd?, :even?, :allbits?, :anybits?, :nobits?, :downto, :times, :pred, :pow, :bit_length, :digits, :rationalize, :lcm, :gcdlcm, :next, :div, :|, :~, :+@, :eql?, :singleton_method_added, :i, :real?, :zero?, :nonzero?, :finite?, :infinite?, :step, :positive?, :negative?, :arg, :rectangular, :rect, :real, :imag, :abs2, :imaginary, :angle, :phase, :conjugate, :to_c, :polar, :conj, :clone, :dup, :quo, :between?, :clamp, :instance_variable_set, :instance_variable_defined?, :remove_instance_variable, :instance_of?, :kind_of?, :is_a?, :tap, :instance_variable_get, :public_methods, :instance_variables, :method, :public_method, :singleton_method, :define_singleton_method, :public_send, :extend, :to_enum, :enum_for, :pp, :=~, :!~, :respond_to?, :freeze, :object_id, :send, :display, :nil?, :hash, :class, :singleton_class, :itself, :yield_self, :taint, :untaint, :tainted?, :untrusted?, :untrust, :frozen?, :trust, :singleton_methods, :methods, :private_methods, :protected_methods, :!, :equal?, :instance_eval, :instance_exec, :!=, :__id__, :__send__]
```

## !
- 元のオブジェクトを書き換える
- 破壊的メソッド

非破壊的メソッド
```ruby
irb(main):005:0> msg = "message!"
=> "message!"
irb(main):006:0> msg.upcase
=> "MESSAGE!"
irb(main):007:0> puts msg
message!
=> nil
```

破壊的メソッド
```ruby
irb(main):008:0> msgg = "messagge!"
=> "messagge!"
irb(main):009:0> msgg.upcase!
=> "MESSAGGE!"
irb(main):010:0> puts msgg
MESSAGGE!
=> nil
```

## ?
- 真偽値を返すメソッド

## 配列
- 配列の添字で範囲指定ができる
```ruby
irb(main):011:0> r = [0, "one", 2]
=> [0, "one", 2]
irb(main):012:0> p r[0..2]
[0, "one", 2]
=> [0, "one", 2]
irb(main):013:0> p r[0...2]  # 0番目~2番目の直前まで
[0, "one"]
=> [0, "one"]
irb(main):014:0> p r[9]  # 範囲外はnilが返却
nil
=> nil
```

- 要素末尾に追加は、push(または、**<<**)
```ruby
irb(main):015:0> r.push("three")
=> [0, "one", 2, "three"]
```
or
```ruby
irb(main):016:0> r << 4
=> [0, "one", 2, "three", 4]
```

- 他
```ruby
irb(main):017:0> r.size  # 要素の数
5
```

- %W で""文字列の配列
```ruby
irb(main):039:0> ar = ["red", "green"]
=> ["red", "green"]
irb(main):040:0> arr = %W(red green)
=> ["red", "green"]
```

- %w で、、？
```ruby
irb(main):001:0> ar = ['red', 'green']
=> ["red", "green"]
irb(main):002:0> arr = %w(red green)
```

## ハッシュ
- keyにシンボルを使うと文字列より高速。ハッシュではよく使用される。
```ruby
irb(main):023:0> h = {"name" => "dd", "age" => 27}  # keyが文字列
=> {"name"=>"dd", "age"=>27}

irb(main):024:0> h = {:name => "ddd", :age => 27}  # keyを文字列からシンボルへ
=> {:name=>"ddd", :age=>27}

irb(main):025:0> h = {name: "dddd", age: 27}  # keyにシンボルを使う場合は、:で区切れる(上と同じ)
=> {:name=>"dddd", :age=>27}

irb(main):026:0> h[:age]  # アクセス
=> 27
```

- メソッド
```ruby
irb(main):025:0> h = {name: "dddd", age: 27}
=> {:name=>"dddd", :age=>27}

irb(main):027:0> h.size
=> 2
irb(main):028:0> h.keys
=> [:name, :age]
irb(main):029:0> h.values
=> ["dddd", 27]

irb(main):030:0> h.has_key?(:name)
=> true
irb(main):031:0> h.has_key?(:length)
=> false

irb(main):032:0> a = h.to_a   # ハッシュ -> 配列　変換
=> [[:name, "dddd"], [:age, 27]]
irb(main):033:0> hh = a.to_h  # 配列 -> ハッシュ　変換
=> {:name=>"dddd", :age=>27}
```

## 書式
- `"文字列" % 値` で表わす。
```ruby
irb(main):005:0> "name: %s" % "dd"
=> "name: dd"
irb(main):006:0> "name: %5s" % "dd"
=> "name:    dd"
irb(main):007:0> "name: %-5s" % "dd"
=> "name: dd   "

irb(main):008:0> "age: %d, weight: %f, blood: %s" % [27, 49.195, "B"]
=> "age: 27, weight: 49.195000, blood: B"
irb(main):009:0> "age: %04d, weight: %3.2f, blood: %s" % [27, 49.195, "B"]
=> "age: 0027, weight: 49.20, blood: B"


irb(main):016:0> printf("age: %04d, weight: %3.2f, blood: %s\n", 27, 49.195, "B")
age: 0027, weight: 49.20, blood: B

irb(main):017:0> txt = sprintf("age: %04d, weight: %3.2f, blood: %s\n", 27, 49.195, "B")
=> "age: 0027, weight: 49.20, blood: B\n"
irb(main):018:0> puts txt
age: 0027, weight: 49.20, blood: B
```

## 条件分岐
### if
```ruby
score = 70
if score > 80 then
    puts "good"
elsif score > 60 then
    puts "good..?"
else
    puts "bad"
end
```
or
```ruby
if score > 80
    puts "good"
elsif score > 60
    puts "good..?"
else
    puts "bad"
end
```

後置if
```ruby
irb(main):001:0> score = 50
=> 50
irb(main):002:0> puts "good" if score > 80
=> nil
irb(main):003:0> puts "good" if score > 40
good
=> nil
```

### case
```ruby
sig = "red"
case sig
when "red"
  puts "stop"
when "blue"
  puts "go"
else
  puts "wrong.."
end         
```

## 繰り返し
### for/each
```ruby
for i in 3..5 do
  puts i
end

# 3
# 4
# 5
```
or
```ruby
for i in 3..5
  puts i
end  
```
or
```ruby
(3..5).each do |i|
  puts i
end
```
or
```ruby
(3..5).each { |i| puts i }
```
---

配列
```ruby
for obj in ["red", 55, "yellow"] do
  puts obj
end

# red
# 55
# yellow
```
or
```ruby
["red", 55, "yellow"].each do |obj|
  puts obj
end 
```
or
```ruby
["red", 55, "yellow"].each { |obj| puts obj }
```

---

ハッシュ
```ruby
for key, value in {tea: 200, coffee: 500} do
  puts "#{key}...#{value}yen"
end

# tea...200yen
# coffee...500yen
```
or
```ruby
{tea: 200, coffee: 500}.each do |key, value|
  puts "#{key}...#{value}yen"
end 
```
or
```ruby
{tea: 200, coffee: 500}.each { |key, value| puts "#{key}...#{value}yen" }
```


### while
```ruby
i = 0
while i < 5 do
  puts "hello"
  i += 1
end 
```

### times
```ruby
5.times do
    puts "hello"
end
```
or
```ruby
5.times { puts "hello" }
```

---

```ruby
5.times do |i|
    puts "hello #{i}"
end
```
or
```ruby
5.times { |i| puts "hello #{i}" }
```

### loop
無限ループ
```ruby
i = 0
loop do
  puts i
  i += 1
end 
```

next/break
```ruby
i = 0
loop do
  if i%2 == 0
    i += 1
    next
  elsif i == 7
    break
  end
  puts i
  i += 1
end

# 1
# 3
# 5
```


### メソッド・クラス・モジュール・mixin・例外
- mtd-cls-excp/README.md 参照

## Tips

- riコマンド: Rubyのhelpコマンド

- puts/p
```ruby
irb(main):001:0> puts "hello" # 改行
hello
=> nil
irb(main):002:0> p "hello" # オブジェクト出力/デバッグ用
"hello"
=> "hello"
irb(main):003:0> p 999
999
=> 999
```

- gets: 標準入力を受け付ける(文字列)
- gets.chomp: 取得した文字列から改行コードを取り除く
- do~end ブロックの中身が1行であれば{~} で置き換え可能(2行以上でも;で区切ればイケるけど、微妙そう)