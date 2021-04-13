from abc import ABCMeta, abstractmethod
from character import Character

# Creator役。Productを生成する抽象クラス。
# 具体的な内容は、Concrete Creator役が定める。
class Factory(metaclass=ABCMeta):
    @abstractmethod
    def create_character(self, name: str) -> Character:
        pass

    @abstractmethod
    def belong_to_group(self, Character):
        pass

    # テンプレートメソッド
    def create(self, name: str) -> Character:
        character: Character = self.create_character(name)
        self.belong_to_group(character)
        return character

    # 感想
    # 各Factoryで作成される時に共通で何か処理する必要がある、といった場合に使えるのかなあ、と。。まだわからない。
    # 例えば、各Factoryで生成するときに、上で言えばbelong_to_groupメソッドでリストに追加する、とか。
