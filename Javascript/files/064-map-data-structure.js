/*
object literal, key --> string
*/
const user = {
    Name: 'neel',
    age: 18
}
console.log(user.Name)
console.log(user['age'])


// map : story key,value pair
// different from object in the sense that key can be of any data type in map

const person = new Map()
person.set('userName', 'neel')
person.set('age', 18)
person.set(1, 'one')


// accessing key
console.log(person.get(1))


// loop with map
for (let credential of person.keys()) {
    console.log(credential, ':' ,person.get(credential))
}


// key as object and array
person.set({game:'2d'}, 'platformer game')
person.set([1,2,3], 'myArray')
console.log(person)


// desctructuring 

console.log()
for (let pair of person) {
    console.log(pair)
    console.log(typeof pair) 

    // this shows pair is object
    // both {}, [] are objects
    // checking for object literal or array,

    console.log(Array.isArray(pair))    // true

    // key value pair is array
}

// now destructring,

for (let [key,value] of person) {
    console.log(`${key}:${value}`)
}



// example

const myFriend = {
    id:1,
    Name:'krishna'
}

const extraInfo = new Map()
extraInfo.set(myFriend, {gender:'male', age:18})

console.log(extraInfo)

// getting age
console.log(extraInfo.get(myFriend).age)

