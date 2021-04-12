from abc import ABCMeta, abstractmethod

class Character(metaclass=ABCMeta):
    @abstractmethod
    def attack(self):
        pass

    @abstractmethod
    def escape(self):
        pass