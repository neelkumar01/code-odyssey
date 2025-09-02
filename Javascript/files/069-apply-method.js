// apply method : for stating other parameters

function myAccount(userName, location) {
    console.log(`Email : ${this.email}
Password : ${this.password}
Username : ${userName}
Location : ${location}

Use these credentials to change your account password`)
}

const user = {
    email: 'xyz@gmail.com',
    password: '123#abc'
}

myAccount.apply(user, ['Neel Kumar', 'Delhi, India'])

// passed rest args in array