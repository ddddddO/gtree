## クラスメソッド、クラス変数、定数アクセス
```ruby
#! /usr/bin/env ruby
class Human
  CLASSNAME = "Human"  # 定数
  @@count = 0          # クラス変数: インスタン共通に使用可能
  attr_accessor :name

  def initialize(name)
    @@count += 1
    @name = name
  end

  def greet
    puts "Hello! #{name} desu."
  end

  # クラスメメソッド: self.メソッド名
  def self.info
    puts "How many human? -> #{@@count}"
  end
end

p Human::CLASSNAME
yamada = Human.new("YAMADA")
Human.info
suganami = Human.new("SUGANAMI")
Human.info

# "Human"
# How many human? -> 1
# How many human? -> 2
```

## 継承
```ruby
#! /usr/bin/env ruby
class Human
  attr_accessor :name
  def initialize(name)
    @name = name
  end

  def greet
    puts "Hello! #{name} desu."
  end
end

class Musician < Human
  def sing
    puts "#{@name} sings."
  end
end

yamada = Musician.new("YAMADA")
yamada.greet
yamada.sing

# Hello! YAMADA desu.
# YAMADA sings.
```

オーバーライド
```ruby
#! /usr/bin/env ruby
class Human
  attr_accessor :name
  def initialize(name)
    @name = name
  end

  def greet
    puts "Hello! #{name} desu."
  end
end

class Musician < Human
  def sing
    puts "#{@name} sings."
  end

  def greet  # ここでオーバーライド
    puts "Hello! #{name} desu. I am Musician."
  end
end

yamada = Musician.new("YAMADA")
yamada.greet
yamada.sing

# Hello! YAMADA desu. I am Musician.
# YAMADA sings.
```

## メソッドのアクセス権
- public
- protected
- private: レシーバーを指定できない

private
```ruby
#! /usr/bin/env ruby
class Human
  def greet  # public
    puts "Hello!"
    greetPrivate
  end

  private  # この後に定義されたメソッドはprivate

  def greetPrivate
    puts "Hello!(private)"
  end
end
matuda = Human.new
matuda.greet
# matuda.greetPrivate  # エラー。private method `greetPrivate' called for #<Human:0x000056259d23c918> (NoMethodError)

# Hello!
# Hello!(private)
```

継承した時の挙動(サブクラスから呼び出し)
```ruby
#! /usr/bin/env ruby
class Human
  def initialize(name)
    @name = name
  end

  def greet  # public
    puts "Hello!"
    greetPrivate
  end

  private  # この後に定義されたメソッドはprivate

  def greetPrivate
    puts "Hello!(private)"
  end
end

class Musician < Human
  def greetByMusician
    puts "Iam Musician."
    greetPrivate  # サブクラスから呼び出せる
  end
end

matuda = Musician.new("MATUDA")
matuda.greetByMusician

# Iam Musician.
# Hello!(private)
```

継承した時の挙動(オーバーライド)
```ruby
#! /usr/bin/env ruby
class Human
  def initialize(name)
    @name = name
  end

  def greet  # public
    puts "Hello!"
    greetPrivate
  end

  private  # この後に定義されたメソッドはprivate
  
  def greetPrivate
    puts "Hello!(private)"
  end
end

class Musician < Human
  def greetPrivate  # オーバーライド可能
    puts "Iam Musician.Hello!(private)"
  end
end

matuda = Musician.new("MATUDA")
matuda.greetPrivate
puts
matuda.greet  # 出力から、スーパークラスのメソッド内で呼ばれるgreetPrivateは、オーバーライドされたものとわかる。

# Iam Musician.Hello!(private)
#
# Hello!
# Iam Musician.Hello!(private)
```

## モジュール
- 名前空間
- モジュール名は大文字始まり
- メソッドや定数をまとめるもの
- インスタンスの作成、継承が出来ない
- ミックスイン: 継承関係に無いクラスへ、共通の機能として組み込みたいときとか

```ruby
#! /usr/bin/env ruby
module Music
  VERSION = 1.0
  
  def self.encode
    puts "encoding..."
  end
  
  def self.decode
    puts "decoding..."
  end
end

p Music::VERSION
Music.encode
Music.decode

# 1.0
# encoding...
# decoding...
```

ミックスイン
```ruby
#! /usr/bin/env ruby
module Debug
  def info
    puts "#{self.class} debug info"
  end
end

class Musician
  include Debug
end

class Audience
  include Debug
end

Musician.new.info
Audience.new.info 

# Musician debug info
# Audience debug info
```


## 例外
```ruby
#! /usr/bin/env ruby
piece = 3
p 100 / piece

piece = 0
begin
  p 100 / piece
rescue => e
  puts "exception!"
  p e
  p e.class
  p e.message
end

puts "---"

def calc(piece)
  return 100 / piece
end

begin
  calc(0)
rescue => ex
  puts "exception!"
  p ex
  p ex.class
  p ex.message
end

# 33
# exception!
# #<ZeroDivisionError: divided by 0>
# ZeroDivisionError
# "divided by 0"
# ---
# exception!
# #<ZeroDivisionError: divided by 0>
# ZeroDivisionError
# "divided by 0"
```

ensure(例外の発生有無にかかわらず実行するブロック)
```ruby
#! /usr/bin/env ruby
begin
  p 100 / 0
rescue => ex
  puts "exception!"
ensure
  puts "end.."
end

# exception!
# end..
```

自前の例外クラス
```ruby
#! /usr/bin/env ruby
class MyError < StandardError; end # StandardError: rubyの標準的なエラー

(0..4).each do |i|
  begin
    if i == 3
      raise MyError  # エラーを発生させる
    end
    puts "proc cnt: #{i}"
  rescue MyError
    p "bad 3!"
  end
end

# proc cnt: 0
# proc cnt: 1
# proc cnt: 2
# "bad 3!"
# proc cnt: 4
```