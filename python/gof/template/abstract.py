from abc import ABCMeta, abstractmethod

class AbstractCrawler(metaclass=ABCMeta):
    @abstractmethod
    def name(self) -> str:
        pass

    @abstractmethod
    def get(self):
        pass

    @abstractmethod
    def scrape(self):
        pass

    @abstractmethod
    def store(self):
        pass

    def execute(self):
        site_name = self.name()

        print(f'start {site_name} crawl.')
        self.get()
        self.scrape()
        self.store()
        print(f'end {site_name} crawl.\n')
