// html file : dom4.html

/*
insertAdjacentHTML(where, html)
beforebegin
afterbegin
beforeend
afterend
*/


const container = document.querySelector('.container')

const myHtml = '<p>My github profile : <a href="https://github.com/neel-epicDeveloper">link</a></p>'

container.insertAdjacentHTML('afterbegin', myHtml)

