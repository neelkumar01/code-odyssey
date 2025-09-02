class dog {

    constructor(name, color, breed) {
        this.name = name
        this.color = color
        this.breed = breed
    }

    // static method  :  accessed by class name not by object

    static classInfo() {
        console.log('This class describe dogs by name, color, breed')
    }

    // static property

    static describe = 'This is dog class'
}

const puppy = new dog('tukku', 'white and brown', 'shih tzu')

dog.classInfo()

console.log(dog.describe)