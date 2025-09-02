const btn1 = document.querySelector('.btn1')

console.log(btn1)

btn1.onclick = function() {
    console.log('hi')
}

const btn2 = document.querySelector('.btn2')

function reloadMe() {
    console.log('I am reloaded')
}
btn2.addEventListener('click', reloadMe)

