# complex inheritance

class Monster:

    def __init__(self, health, energy, **kwargs):
        self.health = health
        self.energy = energy
        super().__init__(**kwargs)
    
    def attack(self):
        print('Attack done')
        self.energy -= 10
        print(f'After attack, energy left: {self.energy}')

    def move(self, speed):
        print(f'moving with speed {speed}')

class Fish:

    def __init__(self, speed, has_scales):
        self.speed = speed
        self.has_scales = has_scales

    def swim(self):
        print(f'fish is swimming at a speed of {self.speed}')

class Shark(Monster, Fish):
    
    def __init__(self, bite_power, health, energy, speed, has_scales):
        self.bite_power = bite_power
        super().__init__(health=health, energy=energy, speed=speed, has_scales=has_scales)

shark = Shark(
    bite_power= 200,
    health= 100,
    energy= 500,
    speed= 280,
    has_scales= False
)

print(shark.has_scales)