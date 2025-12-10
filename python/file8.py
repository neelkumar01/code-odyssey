# simple inheritance

# parent class
class Monster:

    def __init__(self, health, energy):
        self.health = health
        self.energy = energy

    def attack(self):
        print('Attack done')
        self.energy -= 10
        print(f'Remaining energy: {self.energy}')

    def move(self, speed):
        print(f'moving with speed of {speed}')

# child class
class Shark(Monster):

    def __init__(self, speed, health, energy):
        # Monster.__init__(self, health, energy)
        super().__init__(health, energy)
        self.speed = speed

    def bite(self):
        print('bitten!')

    # overwriting inherited method
    def move(self):
        print(f'moving with speed of {self.speed}')

shark = Shark(speed=240, health=50, energy=100)

print(shark.speed)
print(shark.health)

shark.attack()
print(shark.energy)

shark.bite()

shark.move()