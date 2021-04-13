from abc import ABCMeta, abstractmethod

# Product役。生成されるインスタンス(Product)が持つべきインタフェースを定める抽象クラス。
# 具体的な内容は、Concrete Product役が定める。
class Character(metaclass=ABCMeta):
    @abstractmethod
    def attack(self):
        pass

    @abstractmethod
    def escape(self):
        pass