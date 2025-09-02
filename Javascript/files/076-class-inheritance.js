class animal {
    constructor(name, age) {
        this.name = name
        this.age = age
    }

    eat() {
        console.log(`${this.name} is eating`)
    }

    isCute() {
        console.log(true)
    }
}

const myPuppy = new animal('Tukku', 0.6)

myPuppy.eat()

myPuppy.isCute()



// Class Inheritance

class dog {
    constructor(name, age) {
        this.name = name
        this.age = age
    }

    eat() {
        console.log(`${this.name} is eating`)
    }

    isCute() {
        console.log(true)
    }
}
const tommy = new dog('tommy', 1)

/*
- for creating dog class, we have same constructor and methods as in animal class

- it is tedious to write same code again for dog class

- now we make cat class with inheritance
*/

class cat extends animal {

}

const myCat = new cat('kitty', 2)

myCat.eat()