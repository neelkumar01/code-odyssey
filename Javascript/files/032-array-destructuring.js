const folder = ['doc1', 'doc2', 'doc3']

// storing items in variables

let [var1, var2, var3] = folder

console.log(var1)
console.log(var2)
console.log(var3)

const fruits = ['mango']

// when some variables go undefined

let [fruit1, fruit2, fruit3] = fruits

console.log(fruit1)
console.log(fruit2)
console.log(fruit3)

// specifying variables

const myArray = [1, 2, 3, 4]

let [item1, , item2, item3] = myArray

console.log(item2)

// below, ...newArray stores rest of values

const animals = ['tiger', 'lion', 'elephant', 'jaguar', 'dolphin']

let [a1, a2, a3, ...newArray] = animals

console.log(newArray)