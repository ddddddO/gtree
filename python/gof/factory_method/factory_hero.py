from factory import Factory
from character import Character
from character_hero import Hero

class HeroFactory(Factory):
    def __init__(self):
        pass

    def create_character(self, name: str) -> Character:
        hero = Hero(name)
        return hero