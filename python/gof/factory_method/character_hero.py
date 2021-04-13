from character import Character

# Concrete Product役。具体的な製品を定めるクラス。
class Hero(Character):
    def __init__(self, name:str):
        self.__name = name
        print('HERO {} が 誕生した！'.format(self.__name))

    def name(self) -> str:
        return self.__name

    def attack(self):
        print('HERO {} は 攻撃した！'.format(self.__name))
        print('ビッグバン')
        print('メラゾーマ')

    def escape(self):
        print('HERO {} は 逃げ出した！'.format(self.__name))