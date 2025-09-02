// cloning objects

const obj1 = {
    Name: 'neel',
    id: 24
}

const obj2 = obj1

// console.log(obj1)
// console.log(obj2)

/*
- both obj1, obj2 points to same address
- changes in one get reflected in both
*/

// Solving above problem :-

// method 1 : spread operator
const obj3 = {...obj1}
obj1.age = 18

console.log(obj1)
console.log(obj2)
console.log(obj3)   // remain unchanges

// method 2 : Object.assign
const obj4 = Object.assign({}, obj1)
obj1.location = 'India'

console.log('obj1 :', obj1)
console.log('obj4 :', obj4)

