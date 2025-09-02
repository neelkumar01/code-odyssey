// XHR : xml http request

// creating an object for xhr
const xhr = new XMLHttpRequest()

// console.log(xhr)


console.log(xhr.readyState)  // 0 : client created, open() not called

const endPoint_url = 'https://jsonplaceholder.typicode.com/posts'
xhr.open('GET', endPoint_url)

console.log(xhr.readyState)  // 1 : open() called

xhr.onreadystatechange = function() {

    console.log(xhr.readyState)
    // 2 : headers received
    // 3 : loading / downloading
    // 4 : operation complete

    if (xhr.readyState == 4) {
        const response = xhr.response   // in json, type : string
        const data = JSON.parse(response)  // type : object
        console.log(data)
    } 
    /*
    - we get status code for every request
    - status : 2xx --> in this range, it's OK!
    - status : 404 --> error!
    */
}

xhr.onload = function() {
    console.log(xhr.readyState)
} // it runs only when ready state = 4

xhr.send()