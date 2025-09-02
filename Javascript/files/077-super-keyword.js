class animal {

    constructor(name, age) {
        this.name = name
        this.age = age
    }

    about() {
        console.log(`${this.name} is very cute`)
    }
}

// when we want extra properties

class dog extends animal {

    constructor(name, age, speed) {
        super(name, age)
        this.speed = speed
    }

    run() {
        console.log(`${this.name} is running at ${this.speed}kmph`)
    }

    // chnaging base method of parent class
    about() {
        console.log(`${this.name} is super cute`)
    }
}

const puppy = new dog('tukku', 0.6, 20)

puppy.run()

puppy.about()