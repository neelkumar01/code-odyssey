# dunder methods

'''
* dunder method not created by user
* called by python
'''

class Evil:
    
    def __init__(self, power, damage):
        self.power = power
        self.damage = damage
        print('Evil returns')
    
godzilla = Evil(power= 900, damage= 0)

print(godzilla.power)
print(godzilla.damage)

print(godzilla.__dict__) # give attributes


# classes and methods

def add(a,b):
    print(a+b)

class Test:
    def __init__(self, func):
        self.addBoth = func

test = Test(func= add)

test.addBoth(12, 14)