from character import Character

class Hero(Character):
    def __init__(self, name:str):
        self.__name = name
        print('HERO: {} が 誕生した！'.format(self.__name))

    def attack(self):
        print('ビッグバン')
        print('メラゾーマ')

    def escape(self):
        print('HERO は 逃げ出した！')