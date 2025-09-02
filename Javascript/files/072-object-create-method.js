const obj1 = {
    key1: 'value1',
    key2: 'value2'
}

const obj2 = {
    key1: 'value1'
}
console.log(obj2.key2)  // undefined

const obj3 = Object.create(obj1)
console.log(obj3)

obj3.key1 = 'abc'
console.log('obj3 :', obj3)

console.log(obj3.key2)  
// key2 not present in obj3 so it takes key2 from obj1