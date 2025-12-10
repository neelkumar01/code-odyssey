# flow

x = 5
if x > 10:
    print('x > 10')
elif x != 0:
    print('x is not zero')
else:
    print('x < 10')

# -------------------------------------------------
# nested if

var = 'k'
if var in ['a', 'b', 'k']:
    print('letter present!')
    if var.isalpha():
        print('var is a letter')

# -------------------------------------------------
# match case
mood = 'hungry'
match mood:
    case 'hungry':
        print('I am hungry')
    case 'tired':
        print('I am tired')
    case 'bored':
        print('I am bored')

# -------------------------------------------------
# while loop

num = 1
while num <= 5:
    print(num)
    num += 1

    if num == 4:
        break

# -------------------------------------------------
# for loop

for i in 'neel':
    print(i)


# ternary operator

n = 2
color = 'pink' if n<5 else 'purple'

print(color)

