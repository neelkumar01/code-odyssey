const user = {
    Name: 'neel',
    address: {houseNo:62}
}

console.log(user.address.houseNo)

// if any key don't exist and then we print so it will give error!

// solution :-
// using ?.

// console.log(user.address.pincode)   --> error

console.log(user?.address?.pincode)   // undefined
