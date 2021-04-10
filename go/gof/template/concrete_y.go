package template

import (
	"fmt"
)

type siteXCrawler struct {
	name string
	path string
}

func NewSiteXCrawler(name, path string) *siteXCrawler {
	return &siteXCrawler{
		name: name,
		path: path,
	}
}

func (x *siteXCrawler) Name() string {
	return x.name
}

func (x *siteXCrawler) Get() error {
	fmt.Printf("Get request: %s\n", x.path)
	return nil
}

func (x *siteXCrawler) Scrape() error {
	fmt.Println("Scraping now...")
	return nil
}

func (x *siteXCrawler) Store() error {
	fmt.Println("Stored!")
	return nil
}
