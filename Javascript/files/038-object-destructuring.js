const obj1 = {
    id:1,
    email:'xyz@gmail.com',
    code:123,
    myName:'neel'
}

const mail = obj1.email
console.log(mail)

// destructuring

const {id, email, code, myName} = obj1

console.log(myName)
