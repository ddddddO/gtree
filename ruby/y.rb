require "oj"

json = '{"ContentType":"application/json", "aaa": "AAA"}'

res = Oj.load(json)
p res
