const person = {
    name: 'neel',
    age:18,
    email:'xyz@gmail.com'
}

// for in loop

for (let key in person) {
    console.log(key, person[key])
}
// keys of object are always string!

console.log(Object.keys(person))

// for of loop

for (let key of Object.keys(person)) {
    console.log(person[key])
}