from abstract import AbstractCrawler
from concrete_x import SiteXCrawler
from concrete_y import SiteYCrawler

if __name__ == '__main__':
    crawler_x: AbstractCrawler = SiteXCrawler('site-xxx', 'https://xxxx.com')
    crawler_y: AbstractCrawler = SiteYCrawler('site-yyy', 'https://yyyy.com')

    crawler_x.execute()
    crawler_y.execute()

    # Output:
    # start site-xxx crawl.
    # Get request: https://xxxx.com
    # Scraping now...
    # Stored!
    # end site-xxx crawl.
    #
    # start site-yyy crawl.
    # Get request: https://yyyy.com
    # Scraping now.
    # Scraping now..
    # Scraping now...
    # Stored!
    # end site-yyy crawl.
    #