from factory import Factory
from character import Character
from character_hero import Hero

# Concrete Creator役。具体的な製品を作るクラス。
class HeroFactory(Factory):
    def __init__(self):
        self.__hero_guild = []

    def create_character(self, name: str) -> Character:
        hero = Hero(name)
        return hero

    def belong_to_group(self, hero: Hero):
        self.__hero_guild.append(hero)

    def call_heros(self) -> []:
        return [hero.name() for hero in self.__hero_guild]
