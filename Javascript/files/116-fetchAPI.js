const URL = 'https://jsonplaceholder.typicode.com/posts'

const func = fetch(URL)

console.log(func)   // it returns promise

fetch(URL)
    .then((response) => {
        console.log(response)
    })
// from above code, we get to know that response has method json


fetch(URL)
    .then((response) => {
        // response.json()  -->  gives promise
        return response.json()
    })
    .then((data) => {
        console.log(data)
    })
    .catch((error) => {     // reject in fetch is for network error
        console.log(error)
    })


// using fetch to post
fetch(URL, {
    method: 'POST',
    body: JSON.stringify({
        body: 'I am body',
        id: 11,
        title: 'I am title',
        userId: 101
    }),
    headers: {
        'Content-type': 'application/json; charset=UTF-8'
    }
})
    .then((response)=>{
        return response.json()
    })
    .then((data)=>{
        console.log(data)
    })