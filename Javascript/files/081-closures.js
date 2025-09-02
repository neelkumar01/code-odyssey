// example 1

function printFullName(fname, lname) {

    function print() {
        console.log(fname, lname)
    }
    return print
}
const myName = printFullName('neel', 'kumar')
myName()
// closures : fname, lname


// example 2

function func1(c) {
    const a = 'a'
    const b = 'b'
    
    function myFunc() {
        console.log(a,b,c)
    }

    return myFunc
} 

const x = func1('c')
x()
// Closures : a,b,c