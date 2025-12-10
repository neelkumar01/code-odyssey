# Data types

'''
set - {1,2, 'fruit'}
in set, each item is unique

int + float = float
'''

# -------------------------------------------------
# numbers

print(type(10))
print(1 + 1.4)

print(float(2))
print(int(2.4))

print(1.1 * 3)   # 3.3000000000000003

# -------------------------------------------------
# strings

name = "neel"
flower = "daisy"
print(name, flower)

sentence = "He said, \'I am great\'"
print(sentence)

var = "Line1: some text \nLine2: some more text"
print(var)

multiLines = '''Hi
I am Neel
Bye
'''
print(multiLines)

repeat = "I am great, I love myself\n" * 5
print(repeat)

user = "neel"
age = 18
bio = 'Hi my name is {} and I am {}yrs old'.format(user, age)
print(bio)

person = "neel"
birth = "15/01/2007"
data = "user: {user}   dob: {dob}".format(user = person, dob = birth)
print(data)

id = 2
work = "remote"
employee = f"Employee id: {id}  work type: {work}"
print(employee)


# -------------------------------------------------
# list and tuples

myList = [1,2,3, "mango"]
myTupple = (1,2,3, "mango")

print(myList[0])
print(myTupple[1])

cat = "kitty"
print(cat[4])

print(myList[0:3])

# -------------------------------------------------
# unpacking

a,b = (10, 20)
print(a)
print(b)

x,y = [1, "apple"]
print(x)
print(y)

fruit, veggie = "blueberyy", "capsicum"
print(fruit)
print(veggie)

health, energy, weapon = 100, 90, "sword"
print(health, energy, weapon)

value1 = 10
value2 = 20
value1, value2 = value2, value1
print(value1)
print(value2)

# -------------------------------------------------
# dictionary

myDict = {
    "fruit": "apple",
    "veggie": "tomato",
    "id": 24,
    "alive": True
}

print(myDict["id"])

# -------------------------------------------------
# set

myset = {1, 2, 3, 3, 3}
print(len(myset))

'''
elements of set can't be accessed
use type conversion to access them
'''

s1 = {1, 2, 3, 4}
s2 = {1, 8, 9}

print(s1.union(s2))
print("union", s1 | s2)

print(s1.intersection(s2))
print("intersection", s1 & s2)

print(s1.difference(s2))
print("difference", s1 - s2)

# -------------------------------------------------
# boolean

print(1 == 2)
print(1 != 2)

print(1 in [1,2,3,4])
print('n' in 'neel')
