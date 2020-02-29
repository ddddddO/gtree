package analyze

import (
	"fmt"
	"log"
	"strings"

	"github.com/bluele/mecab-golang"
)

type DictionaryAnalyzer struct {
	Url  string
	Ress []string
}

func NewDictionaryAnalyzer(url string, ress []string) DictionaryAnalyzer {
	return DictionaryAnalyzer{
		Url:  url,
		Ress: ress,
	}
}

func (da DictionaryAnalyzer) analyze() (interface{}, error) {
	log.Println("DictionaryAnalyzer proc")

	m, err := mecab.New("-d /usr/lib/x86_64-linux-gnu/mecab/dic/mecab-ipadic-neologd")
	if err != nil {
		panic(err)
	}
	defer m.Destroy()

	log.Println(da.Url)
	for _, res := range da.Ress {
		parseByMecab(m, res)
	}

	return nil, nil
}

// http://grahamian.hatenablog.com/entry/2017/12/08/124139
// mecab-ipadic-NEologdのドキュメント：https://github.com/neologd/mecab-ipadic-neologd/blob/master/README.ja.md
func parseByMecab(m *mecab.MeCab, res string) {
	log.Printf("\ntarget sentence\n--->>>%s\n\n", res)

	tg, err := m.NewTagger()
	if err != nil {
		panic(err)
	}
	defer tg.Destroy()
	lt, err := m.NewLattice(res)
	if err != nil {
		panic(err)
	}
	defer lt.Destroy()

	node := tg.ParseToNode(lt)
	for {
		features := strings.Split(node.Feature(), ",")
		if features[0] == "名詞" {
			log.Println(fmt.Sprintf("%s %s", node.Surface(), node.Feature()))
		}
		if node.Next() != nil {
			break
		}
	}
}
