const correct_url = 'https://jsonplaceholder.typicode.com/posts'

const wrong_url = 'https://jsonplaceholder.typicode.com/posts123'

const xhr = new XMLHttpRequest()

xhr.open('GET', correct_url)   // opens url

xhr.onload = () => {   // when ready state : 4

    if (xhr.status >= 200 && xhr.status < 300) {
        const data = JSON.parse(xhr.response)
        console.log(data)
    } else {
        console.log('something went wrong')
    }
}

xhr.onerror = () => {
    console.log('network / internet connection related error')
}

xhr.send()   // send http req and get response