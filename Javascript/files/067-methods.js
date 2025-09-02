// function inside object

// example 1

const user1 = {
    Name: 'neel',
    age:18,
    about: function() {
        console.log(`user name : ${this.Name} and age: ${this.age}`)
    }
}
user1.about()

// here 'this' is whole 'user' object

const user2 = {
    Name: 'krishna',
    age: 18,
    info: function() {
        console.log(this)
    }
}
user2.info()



// example 2

function bio() {
    console.log(`Person${this.id} is ${this.Name} and likes ${this.hobby}`)
}

const person1 = {id:1, Name:'neel', hobby:'coding', about:bio}

const person2 = {id:2, Name:'mohan', hobby:'dancing', about:bio}

const person3 = {id:3, Name:'soham', hobby:'singing', about:bio}

person1.about()
person2.about()
person3.about()