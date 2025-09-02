// keypress event

const body = document.body

body.addEventListener('keypress', (e) => {
    console.log(e)
})

const btn = document.querySelector('#hover-btn')

btn.addEventListener('mouseover', (e)=>{
    e.target.style.backgroundColor = '#FFD4F2'
})

btn.addEventListener('mouseleave', (e)=>{
    e.target.style.backgroundColor = 'white'
})