const users = [
    {id:1, Name:'neel', section:2},
    {id:2, Name:'ram', section:1}
]

console.log(users)

for (let user of users) {
    console.log(user.id, user.Name)
}
