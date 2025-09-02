const numbers = [1,2,3,4,5]

let myFunc = (value, index) => {
    console.log(`index: ${index} & value: ${value}`)
}

numbers.forEach(myFunc)


const users = [
    {Name: 'neel', id: 1},
    {Name: 'ram', id: 2},
    {Name: 'Krishna', id: 3}
]

users.forEach((obj) => {
    console.log(obj.Name)
})

// forEach() takes callback func as parameter

