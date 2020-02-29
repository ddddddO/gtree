package analyze

import (
	"log"
)

type DeepLearningAnalyzer struct {
	Url  string
	Ress []string
}

func NewDeepLearningAnalyzer(url string, ress []string) DeepLearningAnalyzer {
	return DeepLearningAnalyzer{
		Url:  url,
		Ress: ress,
	}
}

func (dla DeepLearningAnalyzer) analyze() (interface{}, error) {
	log.Println("DeepLearningAnalyzer proc")
	log.Println(dla.Url, dla.Ress)

	return nil, nil
}
