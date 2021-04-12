from abc import ABCMeta, abstractmethod
from character import Character

class Factory(metaclass=ABCMeta):
    @abstractmethod
    def create_character(self, name: str) -> Character:
        pass

    def create(self, name: str) -> Character:
        character = self.create_character(name)
        return character
