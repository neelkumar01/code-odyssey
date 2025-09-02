class user {

    constructor(firstName, lastName, age) {
        this.firstName = firstName
        this.lastName = lastName
        this.age = age
    }

    get fullName() {
        console.log(this.firstName, this.lastName)
    }

    // changeName(firstName, lastName) {
    //     this.firstName = firstName
    //     this.lastName = lastName
    // }

    set changeName(fullName) {
        let [firstName, lastName] = fullName.split(" ")
        this.firstName = firstName
        this.lastName = lastName
    }
}

const user1 = new user('neel', 'kumar', 18)

// calling fullName() method as property  :  using 'get'

user1.fullName

/*
to use changeName(), i will do this,
-->   user1.chnageName('mohan', 'sharma')

but i want to chnage like this,
-->   user1.changeName = 'mohan sharma'

# so we use 'set' method
*/

user1.changeName = 'mohan sharma'

user1.fullName