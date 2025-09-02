// example 1

console.log(this)
console.log(window)
console.log(Name)
var Name = 'neel'
console.log(Name)

// example 2

console.log(this)
console.log(window)
console.log(myFunction)
console.log(fullName)

function myFunction() {
    console.log('this is my function')
}

var fname = 'neel'
var lname = 'kumar'
var fullName = fname + ' ' + lname

console.log(fullName)


// example 3

// console.log(userId)
let userId = 2
