#! /usr/bin/env ruby
module Debug
	def self.info
    puts "#{self.class} debug info"
  end
end

class Musician
  include Debug
end

class Audience
  include Debug
end

Debug.info

Musician.new.info
Audience.new.info
