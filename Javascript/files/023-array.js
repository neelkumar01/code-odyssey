// arrays : collection of items (ordered)

// array is object

let fruits = ['apple', 'mango', 'blueberry']
console.log(fruits)

console.log(fruits[0])
console.log(fruits[1])
console.log(fruits[2])

let _array = [1,2,3, 2.4, 'item', null, undefined, true]

fruits[0] = 'pineapple'
console.log(fruits)    // arrays are mutable

console.log(typeof fruits)
console.log(Array.isArray(fruits))
