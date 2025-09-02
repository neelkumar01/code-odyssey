/*
- prototype stores methods that can be accessed by different objects
*/


function createUser(Name, age) {
    this.userName = Name
    this.userAge = age
}

createUser.prototype.about = function() {
    console.log(this.userName, this.userAge)
}
createUser.prototype.introduce = function() {
    console.log(`Hi I am ${this.userName}`)
}

const user1 = new createUser('neel', 18)
const user2 = new createUser('krishna', 12)

user1.about()
user2.about()
user1.introduce()


console.log('\nuser1 object keys :-')
for (let key in user1) {
    console.log(key)
}
// it also includes keys from prototype


// printing keys excluding ones from prototype
console.log('\nthese keys are not from prototype :-')
for (let key in user1) {
    if (user1.hasOwnProperty(key)) {
        console.log(key)
    }
}

