const myArray = [1,2,3,4]
const newArray = [...myArray]

console.log(newArray)

// spread operator in objects

const obj1 = {
    id:12,
    name:'Neel',
    id:18   // this overwrites previous same key
}

const myId = {
    ...obj1,
    location:'India'
}

console.log(myId)


// spreading a string

const obj2 = {
    ...'abc'
}
console.log(obj2)

const obj3 = {
    ...['item1', 'item2', 'item3']
}
console.log(obj3)