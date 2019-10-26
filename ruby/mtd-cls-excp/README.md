## メソッド
基本
```ruby
#! /usr/bin/env ruby
def hello(name, txt)
  puts "Hello #{name}! #{txt}"
end

hello("d", "Bye")
hello "dd", "Bye" 

# Hello d! Bye
# Hello dd! Bye
```

デフォルト値
```ruby
#! /usr/bin/env ruby
def hello(name = "ddd")
  puts "Hello #{name}!"
end

hello "d"
hello 

# Hello d!
# Hello ddd!
```

返り値
```ruby
#! /usr/bin/env ruby
def hello(name)
  "Hello #{name}!"  # 最後に評価された値を返却
  # retrun "Hello #{name}!"  # 明示的にreturnしてもOK
end

msg = hello "d"
puts msg

# Hello d!
```

## クラス
- クラス名は大文字始まり

基本
```ruby
#! /usr/bin/env ruby
class Human
  def hello  # インスタンスメソッド
    puts "Hello!"
  end
end

h = Human.new
h.hello      

# Hello!
```

コンストラクタ
```ruby
#! /usr/bin/env ruby
class Human
  def initialize(name)
    @name = name  # インスタンス変数。インスタンス内であればどこでも使用可
  end

  def greet  # インスタンスメソッド
    puts "Hello! #{@name} desu."
  end
end

yamada = Human.new("YAMADA")
yamada.greet

suganami = Human.new("SUGANAMI")
suganami.greet                   

# Hello! YAMADA desu.
# Hello! SUGANAMI desu.
```

アクセサ
```ruby
#! /usr/bin/env ruby
class Human
  def initialize(name)
    @name = name
  end

  def greet
    puts "Hello! #{@name} desu."
  end
end

yamada = Human.new("YAMADA")
yamada.greet
yamada.name = "YAMADA MASASHI"  # エラー
yamada.greet
```
↓
```ruby
#! /usr/bin/env ruby
class Human
  attr_accessor :name  # インスタンス変数nameのgetter/setter を作成
  # attr_reader :name  # getterのみ生成

  def initialize(name)
    @name = name
  end

  def greet
    puts "Hello! #{@name} desu."
  end
end

yamada = Human.new("YAMADA")
yamada.greet

yamada.name = "YAMADA MASASHI"  # setter
p yamada.name  # getter
yamada.greet

# Hello! YAMADA desu.
# "YAMADA MASASHI"
# Hello! YAMADA MASASHI desu.
```

self(インスタンス自身 == レシーバー)
```ruby
#! /usr/bin/env ruby
class Human
  attr_accessor :name  # name のgetter/setter を作成

  def initialize(name)
    @name = name  # インスタンス変数。インスタンス内であればどこでも使用可
  end

  def greet
    #puts "Hello! #{@name} desu."
    #puts "Hello! #{self.name} desu."  # self.name == @name。nameのgetterが生成されているので使える
    puts "Hello! #{name} desu."        # selfは更に、省略できる。
  end
end

yamada = Human.new("YAMADA")
yamada.greet   

# Hello! YAMADA desu.
```