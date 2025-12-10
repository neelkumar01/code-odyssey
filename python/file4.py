# functions

def testFunc():
    print('hello from function')

testFunc()

def add(n1, n2):
    res = n1 + n2
    print(res)

add(2,3)

def basket(arg1= 10, arg2= 'none', arg3= (0,0,0)):
    print(arg1)
    print(arg2)
    print(arg3)

basket(arg2= 'neel', arg1= 24, arg3= (1,2,3))

# list unpacking

def printAll(*args):
    for arg in args:
        print(arg)
printAll(1,2,3, 'python')

# keyword unpacking

def printArgs(**args):
    print(args)

printArgs(arg1= 1, arg2= 24, arg3= 100)

def myFunc(*args, **kwargs):
    print(args)
    print(kwargs)

myFunc(1,2,3,4,5, arg1='apple', arg2='banana')


# lambda functions
res = lambda q: q+1
print(res(2))

add = lambda a,b: a+b
print(add(10,30))


