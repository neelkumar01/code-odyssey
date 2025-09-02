const btn3 = document.querySelector('.btn3')

// with arrow func, we get window with this

btn3.addEventListener('click', ()=>{
    console.log(this)
})

// with normal, we get button only
const btn1 = document.querySelector('.btn1')

function clickMe() {
    console.log(this)
}
btn1.addEventListener('click', clickMe)