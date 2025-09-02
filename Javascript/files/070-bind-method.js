function myAccount() {
    console.log(`User name is ${this.Name}`)
}

const user = {
    Name: 'neel'
}

const func = myAccount.bind(user)

func()