// call method

const user1 = {
    Name: 'neel',
    age: 18,
    about: function() {
        console.log(`name:${this.Name} & age:${this.age}`)
    }
}

const user2 = {
    Name: 'krishna',
    age: 12
}

user1.about()  
// 'this' here is 'user1' 

user1.about.call(user2)  
// here 'about' now connected with user2 or 'this' here becomes 'user2'


// example 2

function myAccount() {
    console.log(`Your email is ${this.email} and password is ${this.password}`)
}

const myUser = {
    email: 'abc@gmail.com',
    password: '123#'
}

myAccount.call(myUser)
// Here, we set 'this' = 'myUser'