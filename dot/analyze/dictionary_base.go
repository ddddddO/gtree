package analyze

import (
	"log"
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
	log.Println(da.Url, da.Ress)

	return nil, nil
}
