const myPromise = Promise.resolve('hello')

myPromise.then((value) => {
    console.log(value)
})


//  chaining

function herPromise() {
    return new Promise((resolve, reject) => {
        resolve('your gift')
    })
}

herPromise()
    .then((value) => {
        value += ', how is it...'
        return value   // promise is returning
    })
    .then((new_value) => {
        console.log(new_value)
    })
