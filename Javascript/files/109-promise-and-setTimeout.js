// to resolve / reject after t = 2s

function myPromise() {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            const value = false
            if (value) {
                resolve()
            } else {
                reject()
            }
        }, 2000)
    })
}

myPromise()
    .then(() => { console.log('resolved') })
    .catch(() => { console.log('rejected') })