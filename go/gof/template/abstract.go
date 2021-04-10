package template

import (
	"log"
)

// 抽象メソッドの宣言
type AbstractCrawler interface {
	Name() string
	Get() error
	Scrape() error
	Store() error
}

// 抽象クラスの役。
// テンプレートメソッドを定義して、テンプレートメソッド内で呼ばれる抽象メソッドを宣言するところ。
type Abstract struct {
	crawler AbstractCrawler
}

func NewAbstract(crawler AbstractCrawler) *Abstract {
	return &Abstract{
		crawler: crawler,
	}
}

// テンプレートメソッド
func (abs *Abstract) Execute() error {
	log.Printf("start %s crawl.\n", abs.crawler.Name())

	if err := abs.crawler.Get(); err != nil {
		return err
	}

	if err := abs.crawler.Scrape(); err != nil {
		return err
	}

	if err := abs.crawler.Store(); err != nil {
		return err
	}

	log.Printf("end %s crawl.\n\n", abs.crawler.Name())

	return nil
}

// 感想
// Abstract構造体内は委譲でいいと思うし、そもそも、Abstract構造体を無くして、Excuteメソッドをレシーバに紐づけずに関数として定義して、
// 引数にAbstractCrawlerインタフェースを取るようにし、インタフェースを満たす構造体を渡す、の方がシンプルな実装になる。
