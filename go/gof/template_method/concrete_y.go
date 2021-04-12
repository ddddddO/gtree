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

func (x *siteYCrawler) Name() string {
	return x.name
}

func (x *siteYCrawler) Get() error {
	fmt.Printf("Get request: %s\n", x.path)
	return nil
}

func (x *siteYCrawler) Scrape() error {
	fmt.Println("Scraping now.")
	fmt.Println("Scraping now..")
	fmt.Println("Scraping now...")
	return nil
}

func (x *siteYCrawler) Store() error {
	fmt.Println("Stored!")
	return nil
}
