// let, const are block scope  -  can be accessed within block {} not outside

// var is func scope

{
    let port1 = 'connected'
    const port3 = 'not connected'
    var port2 = 'connected'

    console.log(port3)
}

// console.log(port)  --> error

console.log(port2)

