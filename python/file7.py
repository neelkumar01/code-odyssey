# classes and scope

health = 10

def update_health(value):
    global health
    health += value

update_health(20)

print(health)


class Monster:
    def __init__(self, health, energy):
        self.health = health
        self.energy = energy

    def update_energy(self, value):
        self.energy += value

monster = Monster(health= 100, energy= 50)

print(f'current energy: {monster.energy}')
monster.update_energy(30)
print(f'updated energy: {monster.energy}')

