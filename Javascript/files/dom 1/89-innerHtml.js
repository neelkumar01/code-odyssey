const headline = document.querySelector('#my-heading')

console.log(headline.innerHTML)

headline.innerHTML = 'Neel\'s Homepage'

const myInfo = document.querySelector('.my-info')

console.log(myInfo.innerHTML)

// adding extra p tag in inner html...

myInfo.innerHTML += '<p class="my-para" >This para added using innerHTML</p>'
