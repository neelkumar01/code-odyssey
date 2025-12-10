# OOP - object oriented programming

#-----------------------------------------------------------

''''
* object is container for variables and functions
* each obj has attribute and methods

* class - blueprint for object
'''

class Monster:
    # attributes
    health = 90
    energy = 40

    # methods
    def attack(self, damage):
        print('Attack by monster!')
        print(f'{damage} damage was dealt')

        self.energy -= 10
        print(f'{self.energy} energy left')

monster = Monster()

print(monster.health)
print(monster.energy)
monster.attack(30)


class Car:
    fuel = 100

    def run(self, speed):
        print('\ncar is running')
        print(f'Current speed: {speed}')

        self.fuel -= 5
        print(f'Fuel left : {self.fuel}')

car = Car()
car.run(80)


