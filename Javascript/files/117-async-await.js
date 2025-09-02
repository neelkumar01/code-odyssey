const URL = 'https://jsonplaceholder.typicode.com/posts'

// asyn function returns promise

async function getUrl() {
    const response = await fetch(URL)   // gives resolved promise
    console.log(response)

    const data = await response.json()
    console.log(data)
}

getUrl()