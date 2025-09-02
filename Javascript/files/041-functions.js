function greet() {
    console.log('Hello neel :)')
}

greet()

function drive() {
    console.log('Car is driving...')
}

function applyBrakes() {
    console.log('Car stopped!')
}

drive()
applyBrakes()


// return keyword

function tellMyName() {
    return 'neel'
}
let user_name = tellMyName()
console.log(user_name)


// function parameters

function sum(x,y) {
    return x+y
}
let a = 12
let b = 18
console.log('a + b =', sum(a,b))

function isEven(num) {
    return (num % 2 === 0)
}

console.log(isEven(12))

