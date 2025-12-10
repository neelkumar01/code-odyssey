# extra class parts

'''
* private attributes - can't be chnaged outside of class

'''

class Monster:

    '''This is class for an evil shark monster'''

    def __init__(self, health, energy):
        self.health = health
        self.energy = energy

        # private attribute
        self._specie = 'shark'

    def attack(self):
        print('attack done')
        self.energy -= 10
        print(f'energy left: {self.energy}')

evil = Monster(health= 100, energy= 999)

print(evil._specie)

if hasattr(evil, 'health'):
    print(f'Health of evil shark is {evil.health}')

setattr(evil, 'weapon', 'black magic')
print(evil.weapon)

print(evil.__doc__)

