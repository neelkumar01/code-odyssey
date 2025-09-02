function myFunc() {
    console.log('hello world')
}

// adding property in function
myFunc.myProperty1 = 'value1'
console.log(myFunc)

// prototype

myFunc.prototype.myProperty2 = 'value2'

console.log(myFunc.prototype)