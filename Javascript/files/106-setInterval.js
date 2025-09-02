// setInterval runs call back after given interval

console.log('start')

let interval = 1000

const id = setInterval(() => {
    console.log('inside setInterval')
}, interval)

console.log('end')

clearInterval(id)  // setInterval dont execute!




const box = document.querySelector('.box')

const colors = ['#FF7075', '#C170FF', '#FF70E0', '#70BAFF', '#FFD270', '#A2FF9C']
setInterval(()=>{
    box.style.backgroundColor = colors[Math.floor(Math.random() * colors.length)]
}, 2000)