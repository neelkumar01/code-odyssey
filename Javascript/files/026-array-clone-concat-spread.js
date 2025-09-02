// method 1
let array1 = ['item1', 'item2']
let array2 = array1.slice(0, array1.length)

console.log(array2)


// method 2
let fruits = ['blueberry', 'apple']
let basket = [].concat(fruits)

console.log(basket)


// method 3
let folder1 = ['doc1', 'doc2', 'doc3']

let folder2 = [...folder1]  // spread operator
console.log(folder2)

let folder3 = [...folder1, 'doc4', 'doc5']
let myArray = [...fruits, ...array1, ...folder1]

