// function to create object like above
const userMethods = {
    about: function() {
        console.log(`I likes ${this.hobby}`)
    },
    account: function() {
        console.log(`email is ${this.email} and password is ${this.password}`)
    }
}


function createUser(Name, age, email, password, location, hobby) {
    const myUser = {}
    myUser.Name = Name
    myUser.age = age
    myUser.email = email
    myUser.password = password
    myUser.location = location
    myUser.hobby = hobby

    myUser.account = userMethods.account
    myUser.about = userMethods.about

    return myUser
}

const user1 = createUser('neel', 18, 'abc@gmail.com', '123#', 'India', 'programming')

console.log(user1)

user1.about()
user1.account()

