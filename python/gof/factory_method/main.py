from factory import Factory
from character import Character

from factory_hero import HeroFactory

if __name__ == '__main__':
    hero_factory: Factory = HeroFactory()

    hero_a: Character = hero_factory.create('alis')
    hero_b: Character = hero_factory.create('bob')

    hero_a.attack()
    hero_b.escape()

    print('これまでに作られたHEROたちを呼んだ: {}'.format(hero_factory.call_heros()))

    # Output:
    # HERO alis が 誕生した！
    # HERO bob が 誕生した！
    # HERO alis は 攻撃した！
    # ビッグバン
    # メラゾーマ
    # HERO bob は 逃げ出した！
    # これまでに作られたHEROたちを呼んだ: ['alis', 'bob']