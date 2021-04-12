package template_method

import (
	"fmt"
)

type siteYCrawler struct {
	name string
	path string
}

func NewSiteYCrawler(name, path string) *siteYCrawler {
	return &siteYCrawler{
		name: name,
		path: path,
	}
}

func (y *siteYCrawler) Name() string {
	return x.name
}

func (y *siteYCrawler) Get() error {
	fmt.Printf("Get request: %s\n", x.path)
	return nil
}

func (y *siteYCrawler) Scrape() error {
	fmt.Println("Scraping now.")
	fmt.Println("Scraping now..")
	fmt.Println("Scraping now...")
	return nil
}

func (y *siteYCrawler) Store() error {
	fmt.Println("Stored!")
	return nil
}
