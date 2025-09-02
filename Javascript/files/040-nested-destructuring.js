const users = [
    {id:1, Name:'neel'},
    {id:2, Name:'ram'},
    {id:3, Name:'krishna'}
]

// normal destructuring

const [user1, user2, user3] = users

console.log(user1)

console.log(user1.id)
console.log(user1.Name)

// nested destructuring

const [{Name : name1}, {Name : name2} , {id : id3}] = users
// here, from 1st obj, accessed Name and stored in name1 variable

console.log(name1)
console.log(name2)
console.log(id3)

