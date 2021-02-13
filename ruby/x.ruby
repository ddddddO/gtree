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
matuda.greet

