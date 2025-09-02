const users = [
    {id:1, name:'neel'},
    {id:2, name:'ram'},
    {id:3, name:'krisna'}
]

const myFriend = users.find((friend) => {
    return friend.id === 3
})

console.log(myFriend)