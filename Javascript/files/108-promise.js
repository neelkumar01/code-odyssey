// promise in js

// context : my promise is to create sandwich

console.log('start')
items = ['bread', 'veggies', 'mayo', 'ketchup']

// creating promise

/*
- when promise made true : resolve executes
- when promise made false : reject executes
*/
const sandwich_promise = new Promise((resolve, reject) => {

    if (items.includes('bread') && items.includes('veggies')) {
        resolve({ value: 'veg loaded sandwich' })

    } else {
        reject('not enough items to made one!')
    }
})


// consuming promise
/*
- .then method is used
- takes two call back functions
- 1st one for resolve
- 2nd one for reject or use .catch 
*/
sandwich_promise.then((mySandwich) => {
    console.log('let\'s eat', mySandwich.value)
}
    // , (error) => {console.log(error)}
).catch((error) => { console.log(error) })

setTimeout(() => {
    console.log('hello from setTimeout')
}, 0)

console.log('end')




// function returning promise

console.log()  // just to recognize new code output

// context : promise to make maggi

function maggiPromise() {
    const ingredients = ['water', 'maggi', 'masala']
    return new Promise((resolve, reject) => {
        if (ingredients.includes('water') && ingredients.includes('maggi')) {
            resolve({ value: 'maggi' })
        } else {
            reject('not enough items are there')
        }
    })
}

maggiPromise().then((myMaggi) => {
    console.log('Let\'s eat', myMaggi.value)
}).catch((error) => { console.log(error) })

