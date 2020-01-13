require "mecab"

# twitter api で取得したtextデータが格納されたファイルを読みだす
# (grep -E '^\s{4}"text":' data/timeline.json | sed -e "s/    \"text\": \"RT //g" > data/text.txt)
txts = []
File.open("../data/text.txt"){|f|
    f.each_line{|line|
        txts << line
    }
}

# mecab(install/gem)
# https://madogiwa0124.hatenablog.com/entry/2019/01/13/183433
# apt install libmecab2 libmecab-dev mecab mecab-ipadic mecab-ipadic-utf8 mecab-utils
txts.each do |txt|
    puts MeCab::Tagger.new.parse txt
    puts "----\n\n"
end