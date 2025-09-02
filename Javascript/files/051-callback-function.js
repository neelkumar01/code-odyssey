// accessing function as parameter in another function

let connectPort = (port) => {
    console.log(`Device connected to port ${port}`)
}

let startPort = (callback) => {
    console.log('Enabling soon...')
    callback(8)
}

startPort(connectPort)

