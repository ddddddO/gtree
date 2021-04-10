from abstract import AbstractCrawler

class SiteYCrawler(AbstractCrawler):
    def __init__(self, name:str, path:str):
        self.__name = name
        self.__path = path

    def name(self) -> str:
        return self.__name

    def get(self):
        print('Get request: {}'.format(self.__path))

    def scrape(self):
        print('Scraping now.')
        print('Scraping now..')
        print('Scraping now...')

    def store(self):
        print('Stored!')