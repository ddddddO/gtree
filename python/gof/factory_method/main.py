from factory import Factory
from character import Character

from factory_hero import HeroFactory

if __name__ == '__main__':
    hero_factory: Factory = HeroFactory()

    hero_a: Character = hero_factory.create('alis')
    hero_b: Character = hero_factory.create('bob')

    hero_a.attack()
    hero_b.escape()
