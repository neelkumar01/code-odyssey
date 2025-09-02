// using classes

class createUser {

    // constructor : creates object for us

    constructor(Name, age, email, password) {
        this.Name = Name
        this.age = age
        this.email = email
        this.password = password
    }

    // previoussly using prototype to make a copy of methods that can be accesses by all objects

    // with class, we can define here only

    about() {
        console.log(`User is ${this.Name}`)
    }
}

const user1 = new createUser('neel', 18, 'abc@gmail.com', '123#')

// new keyword calls constructor to create a object

console.log(user1)

user1.about()