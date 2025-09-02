// map also takes callback func

// map always returns an array, crucial that callback func returns something

const numbers = [1,2,3,4]

const myArray = numbers.map((number) => {
    return (number * number)
})

console.log(myArray)


const users = [
    {Name: 'neel', id: 1},
    {Name: 'ram', id: 2}
]

let userNames = users.map((object) => {
    return (object.Name)
})

console.log(userNames)